# 1.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 单例 = **受控创建**的全局变量：只创建一次，全局可访问。
    - Go 中用 `sync.Once` / `sync.OnceValue` 实现，天然并发安全，不要手写加锁和双重检查。
    - 单例是全局状态，能用参数传递（依赖注入）时优先传递。

## Q & A

**Q**：初始化失败了，能重试吗？

不能直接重试。`sync.Once` 只执行一次，即使失败或 panic 也不会再执行；`sync.OnceValues` 会缓存首次返回的 error。需要重试时，用互斥锁自己控制：成功才保存实例，失败下次再试。

**Q**：单例让单元测试很难写，怎么办？

多个测试共享同一个全局实例，状态会互相影响，也无法换成假对象。做法是让业务代码通过参数（或接口）接收实例，只在 `main` 里调用 `GetInstance()`；测试时直接传入假对象。

**Q**：什么都放进单例，是不是更方便？

短期方便，长期有害。单例很容易变成什么都往里塞的「上帝对象」，任何代码都能读写它，依赖关系完全隐藏。只把真正全局唯一的资源做成单例。

**Q**：拿到单例后，它的字段还能被改吗？

能。单例只保证「只有一个」，不保证「不被修改」。需要只读时，只暴露方法或返回副本。另外 `var GetConfig = sync.OnceValue(...)` 是导出变量，可被重新赋值；要彻底防替换，用经典写法的函数。

**Q**：其它语言的「双重检查加锁」在 Go 里能用吗？

不要手写。先判空、加锁、再判空的写法在 Go 中容易出现数据竞争；`sync.Once` 内部已经正确、高效地实现了它，直接用即可。

## 动手试试

1. **给日志器也做一个单例**：仿照 `GetConfig`，用 `sync.OnceValue` 写一个 `GetLogger`，用原子计数器记录创建次数，再让 100 个 goroutine 同时调用它。可直接在本章 Playground 里完成。

    提示：计数器放在传给 `sync.OnceValue` 的函数里，用 `sync.WaitGroup` 等所有 goroutine 结束。检验：所有调用拿到同一个指针，创建次数为 1；`singleton.GetInstance()` 的行为不受影响。

2. **观察初始化 panic 后的两种结局**：分别用 `sync.OnceValue` 和 `sync.Once` 包一个必定 `panic("加载失败")` 的初始化函数，各调用两次，用 `recover` 接住。可直接在本章 Playground 里完成。

    提示：写一个 `try(f func()) (p any)` 辅助函数，在 `defer` 里 `p = recover()`。检验：`OnceValue` 两次都得到 `加载失败`，`Once` 第二次不再 panic、返回 `nil`；两种写法的初始化函数都只执行 1 次。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/creational/singleton/exercise_test.go)
