# boot-plguin

[![Build Status](https://github.com/jsmzr/boot-plugin/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-plugin/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-plugin/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-plugin)

boot plugin 旨在简化组件、库的使用，搭配 boot 系列库使用可以简单、快速的完成项目的创建、开发、维护。

## 如何开始

### 开发进度

- [ ] config
    - [x] boot-plugin-apollo
- [ ] db
    - [x] boot-plugin-oracle
    - [x] boot-plugin-gorm-mysql
- [ ] mertrics
    - [x] boot-plugin-prometheus
- [ ] trace
    - [x] boot-plugin-skywalking
- [ ] log
    - [x] boot-plugin-logrus
- [ ] cache
    - [x] boot-plugin-redis
- [ ] api document
    - [x] boot-plugin-gin-swagger


### 插件的开发

1. 依赖 `boot-plugin` 库
2. 实现定义的 Plugin 接口
    1. 插件开关
    2. 插件加载顺序
    3. 插件加载逻辑
3. 调用 `Register` 方法注册插件

未避免插件扩展问题，通常插件的开关和顺序不应写死，请都使用 viper 来获取

按照约定大于配置的规则，通常插件应该有一些必要的默认配置项，以减少配置管理工作量。通过 `InitDefaultConfig` 方法更新默认配置。

在 plugin 中最终是使用 viper 做配置管理，并未将各配置源的配置项合并。这种情况下并不建议使用 `Unmarshal` 的方式来获取配置。

配置获取优先级为 flag, env, config file, key/value store

当前区分了两种插件顺序

1. order < 0  的会按照顺序先加载，适用于一些基础插件
2. order > 0 的会在基础插件加载完成后再重新排序，再按顺序加载

后续计划支持 order > 的插件并行加载

顺序值当前也做了初步的定义

| 类型 | 顺序 | 备注 |
| ---- | ---- | ---- |
| config | -30 ~ -20 | 配置插件，如 aollo |
| log | -20 ~ -10 | 日志插件，如 logrus, zap |
| 预留 | -10 ~ 0 | |
| db | 0 ~ 10 | 数据库 |
| cache | 10 ~ 20 | 缓存 |

### 插件的使用

1. 依赖所需要的插件库
2. 导入插件库，`import _ "github.com/jsmzr/boot-plugin-logrus"`
3. 显式初始化插件 `import plugin "github.com/jsmzr/boot-plugin"`, `plugin.PostProccess()`

完成插件初始化后即可使用对应插件功能

如果是使用 gin 或者 echo 框架，可以直接使用 `github.com/jsmzr/boot-gin` 和 `boot-echo` 库

```go
func main() {
    if err := boot.Run(); err != nil {
        fmt.Println(err)
    }
}
```
