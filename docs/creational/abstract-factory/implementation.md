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

    `lightButton`、`darkCheckbox` 都是小写。包外拿不到具体类型，只能通过工厂创建，也就不可能自己拼出一个「白底勾选框 + 黑底按钮」。

## 使用示例

主题通常在启动时根据用户设置选定一次，之后只传 `Theme`：

```go
--8<-- "codes/creational/abstractfactory/example_test.go:example"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/FDmna2IXdTS){ .md-button .md-button--primary target=_blank }
