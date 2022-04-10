# bootstrap-plguin

bootstrap plugin 旨在简化组件、库的使用，搭配 bootstrap 系列库使用可以简单、快速的完成项目的创建、开发、维护。

## 如何开始

### 插件的开发

1. 依赖 `bootstrap-plugin` 库
2. 在 Load 接口实现中编写具体插件的初始化逻辑
3. 通过 `init` 方法将注册插件

### 插件的使用

1. 依赖所需要的插件库
2. 导入插件库，`import _ "github.com/jsmzr/bootstrap-plugin-logrus/logrus"`
3. 显式初始化插件 `import "github.com/jsmzr/bootstrap-plugin/plugin"`, `plugin.PostProccess()`

完成插件初始化后即可使用对应插件功能
