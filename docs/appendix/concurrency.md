# 附录 A &nbsp; Go 并发速览

## goroutine：启动后不等它结束

在函数调用前加 `go`，这次调用就在新的 goroutine 中执行，调用方继续往下走。多个 goroutine 可以交错执行，不一定要在不同 CPU 核上同时跑。

## 数据竞争：等待结束不等于保护数据

两个 goroutine 访问同一内存位置，至少有一次写入，而且没有同步为它们建立先后关系，这就叫数据竞争（data race）。下面用 `sync.WaitGroup` 等 1000 个任务结束，但它不会规定各次 `count++` 的先后：

```go
--8<-- "codes/appendix/concurrency/example_test.go:race"
```

`count++` 是「读出 → 加 1 → 写回」：两边都读到 5，再各自写回 6，就少加了一次。

## 用同一把锁保护共享变量

并发执行的 `count++` 用同一把锁保护，一次只允许一个 goroutine 修改计数。`Wait` 返回后，这里才读取最终结果。

```go
--8<-- "codes/appendix/concurrency/example_test.go:mutex"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/ByZloUFm9do){ .md-button .md-button--primary target=_blank }

在线可比较两个结果：反例可能小于 1000，也可能恰好是 1000；加锁后一定是 1000。普通测试不运行故意有竞争的 Example（它没有 `Output`）。要看检测报告，在仓库根目录执行：

```bash
go test -race -tags=race_demo -count=1 -run '^TestIntentionalRace$' ./codes/appendix/concurrency
```

需要支持 `-race` 的平台、开启 CGO 并装有兼容的 C 编译器。看到 `WARNING: DATA RACE` 且命令失败，才是这个反例的预期结果。

| 遇到什么 | 怎么处理 |
|---|---|
| 普通 map 有并发写 | 用同一把锁保护相关读写，如[享元缓存](../structural/flyweight/summary.md) |
| 外部回调可能再次进入当前对象 | 锁内复制订阅列表，释放锁后调用，如[观察者](../behavioral/observer/summary.md) |
| 多处调用，但只想初始化一次 | 用 `sync.Once`，见[单例实现](../creational/singleton/implementation.md) |
