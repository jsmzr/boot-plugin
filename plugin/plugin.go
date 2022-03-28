package plugin

import "fmt"

// 插件只需要基于接口进行实现即可，无需考虑其他逻辑
type Plugin interface {
	Load() error
}

var plugins = make(map[string]Plugin)

func Register(name string, plugin Plugin) {
	_, ok := plugins[name]
	if ok {
		panic(fmt.Errorf("插件 [%s] 已经加载，请勿重复加载", name))
	}
	plugins[name] = plugin
}

// 后置处理，再插件注册完毕后开始批量载入插件
func PostProccess() error {
	// 可以考虑并行
	for _, v := range plugins {
		err := v.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
