package plugin

import (
	"fmt"
	"sort"
)

// 插件只需要基于接口进行实现即可，无需考虑其他逻辑
type Plugin interface {
	Load() error
	Order() int
}

var plugins = make(map[string]Plugin)

func Register(name string, plugin Plugin) {
	_, ok := plugins[name]
	if ok {
		panic(fmt.Errorf("插件 [%s] 已经加载，请勿重复加载", name))
	}
	plugins[name] = plugin
}

// 后置处理，插件注册完毕后开始批量载入插件
func PostProccess() error {
	values := make([]Plugin, 0, len(plugins))
	for _, v := range plugins {
		values = append(values, v)
	}
	// 排序后加载插件，避免相互依赖的插件冲突
	sort.Slice(values, func(i, j int) bool {
		return values[i].Order() < values[j].Order()
	})

	for i := 0; i < len(values); i++ {
		fmt.Printf("初始化插件: [%T] \n", values[i])
		if err := values[i].Load(); err != nil {
			return err
		}
	}
	return nil
}
