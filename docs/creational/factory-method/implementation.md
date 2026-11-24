# 2.2 Go 实现

## 三步写完

**第 1 步**：`*EmailNotifier` 和 `*SMSNotifier` 都实现 `Notifier`。注册流程只需要它的 `Send` 方法，不必知道具体类型。不熟悉指针接收者，可先看[接口与方法](../../appendix/interfaces.md)。

```go
--8<-- "codes/creational/factorymethod/notifier.go:product"
```

**第 2 步**：`SignUpWith` 接收一个创建函数。它负责发欢迎通知，具体怎么造通知器则由 `create(kind)` 决定。创建失败就直接返回，不发送通知。

```go
--8<-- "codes/creational/factorymethod/notifier.go:client"
```

**第 3 步**：调用方选择要传哪个创建函数。`NewNotifier` 是固定用 `switch` 按名称选择的**简单工厂**；[下一节](registry.md)的 `Create` 改用注册表。两种写法都能插入同一个流程，不用增加工厂类。

## 使用示例

传 `NewNotifier` 创建短信通知器；传 `Create` 使用已登记的微信通知器。传的是函数本身，不是调用结果（见[函数值与回调](../../appendix/callbacks.md)）。

```go
--8<-- "codes/creational/factorymethod/example_test.go:example"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/bjtH4USqxas){ .md-button .md-button--primary target=_blank } [:material-github: 查看完整源码](https://github.com/caticat/design-pattern/tree/master/codes/creational/factorymethod){ .md-button target=_blank }
