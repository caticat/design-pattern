# 3.2 Go 实现

## 四步写完

**第 1 步**：抽象产品。每种控件一个接口。

```go
--8<-- "codes/creational/abstractfactory/theme.go:product"
```

**第 2 步**：抽象工厂。每个方法创建一种产品，合起来就是「一族」。

```go
--8<-- "codes/creational/abstractfactory/theme.go:factory"
```

**第 3 步**：具体工厂。浅色主题只生产白底控件；深色主题 `Dark` 写法相同，只是换成黑底。

```go
--8<-- "codes/creational/abstractfactory/light.go:light"
```

**第 4 步**：客户端只认 `Theme`，所有控件都向它要。

```go
--8<-- "codes/creational/abstractfactory/theme.go:client"
```

!!! tip "具体产品为什么不导出"

    `lightButton`、`darkCheckbox` 都是小写。包外拿不到具体类型，只能通过工厂创建。`LoginForm` 只向传入的同一个 `Theme` 要控件，所以它拿到的一族必然配套；但接口本身不阻止调用方分别从两个主题取按钮和勾选框再组合，要连这种用法也禁止，需要额外的类型约束。

## 使用示例

主题通常在启动时根据用户设置选定一次，之后只传 `Theme`：

```go
--8<-- "codes/creational/abstractfactory/example_test.go:example"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/TLKqcXGSbRb){ .md-button .md-button--primary target=_blank } [:material-github: 查看完整源码](https://github.com/caticat/design-pattern/tree/master/codes/creational/abstractfactory){ .md-button target=_blank }
