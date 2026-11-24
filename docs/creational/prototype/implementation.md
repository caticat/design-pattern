# 5.2 Go 实现

## 三步写完

**第 1 步**：原型。除了 `Name`、`HP` 这样的普通字段，还特意放了切片、map 和指针，它们是复制时最容易出错的地方。

```go
--8<-- "codes/creational/prototype/monster.go:monster"
```

**第 2 步**：`Clone`。先 `c := *m` 整体复制一层，再把切片、map、指针指向的数据也各复制一份，这样副本才完全独立。==只复制一层是原型模式最常见的 bug==，原因见 [5.3](copy.md)。

```go
--8<-- "codes/creational/prototype/monster.go:clone"
```

**第 3 步**：原型注册表。登记时存一份副本，取出时也给一份副本，原型本身从不交出去。

```go
--8<-- "codes/creational/prototype/registry.go:registry"
```

!!! tip "登记时为什么也要复制"

    如果直接存 `m`，调用方登记之后还拿着同一个对象，随手一改，原型就跟着变了。存副本之后，注册表里的原型只有注册表自己能碰到。

## 使用示例

```go
--8<-- "codes/creational/prototype/example_test.go:example"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/DepUJZAILMe){ .md-button .md-button--primary target=_blank } [:material-github: 查看完整源码](https://github.com/caticat/design-pattern/tree/master/codes/creational/prototype){ .md-button target=_blank }
