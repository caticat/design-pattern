# 附录 C &nbsp; 函数值与回调

## 把函数交出去，还是现在就执行？

`fan.Start` 没有括号，是一个绑定了目标风扇的**方法值**；`fan.Start()` 才是立即启动。把函数值交给另一个函数或对象，让它在合适的时候调用，这样传入的函数就叫**回调**。例如[命令章](../behavioral/command/implementation.md)的队列先存下 `fan.Start`，到 `Run` 才执行。

## 闭包捕获的是变量

匿名函数可以使用外层变量，这样的函数叫**闭包**。下面先排队、再改 `title`，执行时看到的是改后的值。若想保存排队那一刻的内容，用一个新变量保存它，再捕获新变量。

```go
--8<-- "codes/appendix/callbacks/example_test.go:example"
```

## 怎么让函数实现接口？

直接接收 `func()` 就能调用回调。如果已有接口要求 `Execute`，给**具名函数类型**加一个转发方法即可：

```go
--8<-- "codes/behavioral/command/command.go:func"
```

`command.Func(fan.Start)` 是类型转换，不会启动风扇；调用 `Execute` 才会。标准库的 `http.HandlerFunc` 也用这个办法，把普通函数接到 `http.Handler` 接口上。
