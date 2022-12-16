package plugin

import (
	"fmt"
	"sort"
	"time"
)

type Plugin interface {
	Load() error
	// if plugin is basic, then order < 0 and serial load
	// if plugin is't basic, then order > 0 and parallel load
	Order() int
	Enabled() bool
}

var plugins = make(map[string]Plugin)

func log(message string) {
	fmt.Printf("[BOOT-plugin] %v |%s\n", time.Now().Format("2006-01-02 15:04:05"), message)
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
	if err := initConfig(); err != nil {
		return err
	}
	basePlugins := make([]Plugin, 0, len(plugins))
	otherPlugins := make([]Plugin, 0, len(plugins))
	for _, v := range plugins {
		if v.Order() < 0 {
			basePlugins = append(basePlugins, v)
		} else {
			otherPlugins = append(otherPlugins, v)
		}
	}

	if err := baseLoad(basePlugins); err != nil {
		return err
	}
	return fastLoad(otherPlugins)
}

func sortPlugin(plugins []Plugin) {
	sort.Slice(plugins, func(i, j int) bool {
		return plugins[i].Order() < plugins[j].Order()
	})
}

func loadPlugins(plugins []Plugin) error {
	size := len(plugins)
	for i := 0; i < size; i++ {
		plugin := plugins[i]
		if !plugin.Enabled() {
			log(fmt.Sprintf("Plugin [%T] enabled config is false", plugin))
			continue
		}
		log(fmt.Sprintf("Load [%T] plugin", plugin))
		if err := plugin.Load(); err != nil {
			return err
		}
	}
	return nil
}

func baseLoad(plugins []Plugin) error {
	sortPlugin(plugins)
	return loadPlugins(plugins)
}

func fastLoad(plugins []Plugin) error {
	// refresh order by config
	sortPlugin(plugins)
	// use goroutine
	return loadPlugins(plugins)
}
