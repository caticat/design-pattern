# 12.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 享元 = 把**不变、可共享**的内部状态抽出来只存一份，大量对象共用。
    - 每个对象独有的外部状态留在对象里，或在调用时作为参数传入。
    - 享元工厂负责缓存：相同的只创建一次；享元本身必须不可变。

## Q & A

**Q**：怎么分辨内部状态和外部状态？

问一句：**这个数据在同一种对象之间是否都一样**？都一样（树的名字、贴图）就是内部状态，放进享元；每个对象不同（位置）就是外部状态，留在外面。

**Q**：多个 goroutine 同时调用 `TreeFactory.Get` 安全吗？

不安全，map 并发读写会直接崩溃。给 `TreeFactory` 加一个 `sync.Mutex` 保护 `Get`；如果享元只是简单的值，直接用 `unique.Make` 更省事（见[附录 A](../../appendix/concurrency.md)）。

**Q**：Go 里的指针和切片本来就是共享的，还需要享元吗？

指针只解决了「怎么共享」，享元模式还管「**同一份数据只创建一次**」：工厂保证十万棵松树拿到的是同一个 `*TreeType`，而不是十万个内容相同的副本。

**Q**：享元工厂里的对象会一直占着内存吗？

本章的 map 会。种类很少时没问题；种类很多且会变化时，要么定期清理，要么用 `unique`，没人引用的值会被 GC 自动回收。

## 动手试试

1. **加一种上下文：树苗**：写类型 `Sapling`，字段为 `x, y int` 和 `kind *flyweight.TreeType`，`Draw` 返回 `kind.Draw(x, y) + " 树苗"`；两棵树苗都从同一个 `flyweight.NewTreeFactory()` 的 `Get("松树")` 拿享元。可直接在本章 Playground 里完成。

    提示：外部状态（位置）放在 `Sapling` 里，享元和工厂原样复用。检验：「加载贴图：松树」只打印一次，两棵的 `kind` 用 `==` 比较为 `true`，`TreeType`、`TreeFactory` 没改。

2. **让工厂撞上并发，再从外面加锁**：100 个 goroutine 同时对一个 `TreeFactory` 调 `Get("松树")` 或 `Get("橡树")`；再写 `SafeFactory`（`sync.Mutex` + `*flyweight.TreeFactory`），在 `Get` 里加锁后转调。包装可在 Playground 写，竞态需在本机写成外部测试后运行 `go test -race ./codes/structural/flyweight`。

    提示：用 `sync.WaitGroup` 等所有 goroutine 结束。检验：直接用工厂时报 `DATA RACE`，「加载贴图」常常多于 2 行，享元被重复创建；换成 `SafeFactory` 后不再报错，恰好 2 行。
