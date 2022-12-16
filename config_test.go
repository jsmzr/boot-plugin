package plugin

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestInitEnv(t *testing.T) {
	key := "foo"
	value := "bar"
	os.Setenv("BOOT_FOO", value)
	foo := viper.GetString(key)
	if foo == value {
		t.Fatal("viper get foo should not be bar")
	}

	initEnv()
	foo = viper.GetString(key)
	if foo != value {
		t.Fatal("viper get foo should be bar")
	}
}

func TestInitDefault(t *testing.T) {
	fileKey := "boot.config.file"
	pathKey := "boot.config.path"
	initDefault()
	file := viper.GetString(fileKey)
	path := viper.GetString(pathKey)
	if file != "application.yaml" {
		t.Fatalf("viper get %s should be application.yaml", fileKey)
	}
	if path != "." {
		t.Fatalf("viper get %s should be .", pathKey)
	}
}

func TestInitConfig(t *testing.T) {
	if err := initConfig(); err != nil {
		t.Fatal("init application.yaml error should be ignore")
	}
	viper.SetDefault("boot.config.file", "application_test.yaml")
	if err := initLocalConfig(); err != nil {
		t.Fatalf("init application_test.yaml should be success, %v", err)
	}
	name := viper.GetString("boot.application.name")
	if name != "demo" {
		t.Fatal("viper get boot.application.name should be demo")
	}
}
