# 2.2 Go 实现

## 三步写完

**第 1 步**：产品接口和具体产品。

```go
--8<-- "codes/creational/factorymethod/notifier.go:product"
```

**第 2 步**：工厂函数。`switch` 只在这里写一次，创建细节也只在这里出现；不认识的名字返回错误，而不是 `nil`。

```go
--8<-- "codes/creational/factorymethod/notifier.go:factory"
```

**第 3 步**：调用方只通过工厂拿到接口，看不到 `EmailNotifier`、`SMSNotifier`。

```go
--8<-- "codes/creational/factorymethod/notifier.go:client"
```

!!! tip "返回接口，还是返回具体类型"

    Go 惯例是「接收接口，返回具体类型」，例如 `bytes.NewBuffer` 返回 `*bytes.Buffer`。只有工厂**要在多个类型中选一个**时才返回接口，这正是工厂方法的场景。

## 使用示例

通知方式一般来自配置文件或用户设置，换成什么都不影响业务代码：

```go
--8<-- "codes/creational/factorymethod/example_test.go:example"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/TODO){ .md-button .md-button--primary target=_blank }

Playground 中还包含 [2.3](registry.md) 的注册表示例，可以一起运行。
