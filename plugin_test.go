package plugin

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

type TestPlugin struct{}
type Test1Plugin struct{}
type Test2Plugin struct{}
type TestErrorPlugin struct{}
type Test1ErrorPlugin struct{}

func (t *TestPlugin) Load() error {
	return nil
}

func (t *TestPlugin) Order() int {
	return -1
}

func (t *TestPlugin) Enabled() bool {
	return true
}

func (t *Test1Plugin) Load() error {
	return nil
}

func (t *Test1Plugin) Order() int {
	return 100
}
func (t *Test1Plugin) Enabled() bool {
	return false
}
func (t *Test2Plugin) Load() error {
	return nil
}

func (t *Test2Plugin) Order() int {
	return 200
}
func (t *Test2Plugin) Enabled() bool {
	return true
}
func (t *TestErrorPlugin) Load() error {
	return fmt.Errorf("mock load error")
}

func (t *TestErrorPlugin) Order() int {
	return 100
}

func (t *TestErrorPlugin) Enabled() bool {
	return true
}

func (t *Test1ErrorPlugin) Load() error {
	return fmt.Errorf("mock load error")
}

func (t *Test1ErrorPlugin) Order() int {
	return -2
}

func (t *Test1ErrorPlugin) Enabled() bool {
	return true
}

func TestRegister(t *testing.T) {
	plugins = make(map[string]Plugin)
	Register("test", &TestPlugin{})
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("should panic: plugin [test] already register")
		}
	}()
	Register("test", &TestPlugin{})
}

func TestPostProccess(t *testing.T) {
	plugins = make(map[string]Plugin)
	Register("test", &TestPlugin{})
	Register("test1", &Test1Plugin{})
	Register("test2", &Test2Plugin{})
	if len(plugins) != 3 {
		t.Fatal("register failed")
	}
	if err := PostProccess(); err != nil {
		t.Fatal(err)
	}
}

func TestPostProccessError(t *testing.T) {
	viper.Set("boot.config.file", "bad_config_test.yaml")
	if err := PostProccess(); err == nil {
		t.Fatal("init config should be error")
	}

	viper.Set("boot.config.file", "application.yaml")
	plugins = make(map[string]Plugin)
	Register("test", &TestPlugin{})
	Register("testError", &TestErrorPlugin{})
	if err := PostProccess(); err == nil {
		t.Fatal("post proccess should be error")
	}
	plugins = make(map[string]Plugin)
	Register("test", &TestPlugin{})
	Register("testError1", &Test1ErrorPlugin{})
	if err := PostProccess(); err == nil {
		t.Fatal("post proccess should be error")
	}
}
