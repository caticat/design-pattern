# 附录 A &nbsp; Go 并发速览

正文涉及并发时，默认读者了解下面三个概念。不熟悉的话，读完这一页就能看懂各章。

## goroutine

goroutine 是 Go 的轻量级线程：在函数调用前加 `go`，这个函数就在新的 goroutine 里**并发**执行，调用方不等它结束，继续往下走。

Go 程序天然是并发的。例如 `net/http` 的服务器为每个连接启动一个 goroutine，同一时刻可能有成百上千个 goroutine 在调用同一个函数。

## 数据竞争

多个 goroutine 无同步地同时读写同一个变量，叫数据竞争（data race）。下面 1000 个 goroutine 各给 `count` 加 1，结果却常常小于 1000：

```go
--8<-- "codes/appendix/concurrency/example_test.go:race"
```

`count++` 实际是「读出 → 加 1 → 写回」三步：两个 goroutine 同时读到 5，各自写回 6，就少加了一次。用 `go test -race` 或 `go run -race` 运行，Go 会报告发生竞争的位置。

## 用 sync 包同步

给 `count++` 加上互斥锁，同一时刻只有一个 goroutine 能执行它：

```go
--8<-- "codes/appendix/concurrency/example_test.go:mutex"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/2KWR9DB6WdR){ .md-button .md-button--primary target=_blank }

运行可以看到：`Example_race` 每次输出不同，`Example_mutex` 总是 1000。

| 工具 | 作用 | 出现在 |
|---|---|---|
| `sync.Mutex` | 互斥锁：同一时刻只有一个 goroutine 能执行 `Lock` 与 `Unlock` 之间的代码 | 上例 |
| `sync.WaitGroup` | 等一组 goroutine 全部结束 | 上例 |
| `sync.Once` | 一段代码只执行一次，并发调用也一样 | [1.2 单例的 Go 实现](../creational/singleton/implementation.md) |
