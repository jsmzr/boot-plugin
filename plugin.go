package plugin

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Plugin interface {
	Load() error
	Order() int
}

var plugins = make(map[string]Plugin)

func log(message string) {
	fmt.Printf("[BOOTSTRAP-plugin] %v |%s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}
func Register(name string, plugin Plugin) {
	_, ok := plugins[name]
	if ok {
		panic(fmt.Errorf("plugin [%s] already register", name))
	}
	log(fmt.Sprintf("Register [%s] plugin", name))
	plugins[name] = plugin
}

func PostProccess() error {
	values := make([]Plugin, 0, len(plugins))
	for _, v := range plugins {
		values = append(values, v)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].Order() < values[j].Order()
	})
	valueTypes := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		valueTypes[i] = fmt.Sprintf("%T", values[i])
	}
	log(fmt.Sprintf("Plugin sort [%s]", strings.Join(valueTypes, ", ")))
	for i := 0; i < len(values); i++ {
		log(fmt.Sprintf("Load [%T]", values[i]))
		if err := values[i].Load(); err != nil {
			return err
		}
	}
	return nil
}
