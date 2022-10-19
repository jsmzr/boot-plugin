package plugin

import (
	"fmt"
	"sort"
	"strings"
)

type Plugin interface {
	Load() error
	Order() int
}

var plugins = make(map[string]Plugin)

func Register(name string, plugin Plugin) {
	_, ok := plugins[name]
	if ok {
		panic(fmt.Errorf("plugin [%s] already register", name))
	}
	fmt.Printf("[Bootstrap-Plugin]  Register [%s] plugin.\n", name)
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
	fmt.Printf("[Bootstrap-Plugin]  Plugin sort [%s].", strings.Join(valueTypes, ", "))
	for i := 0; i < len(values); i++ {
		fmt.Printf("[Bootstrap-Plugin] Load [%T]. \n", values[i])
		if err := values[i].Load(); err != nil {
			return err
		}
	}
	return nil
}
