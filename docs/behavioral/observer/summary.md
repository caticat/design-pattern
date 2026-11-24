# 14.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 观察者 = 关心某件事的对象**自己来订阅**，事件发生时主题挨个通知。
    - 主题只认识观察者接口，加减下游不用改主题的代码。
    - Go 里常用函数当观察者，订阅时返回取消函数用来退订。

## Q & A

**Q**：某个观察者很慢，会拖慢支付吗？

会，通知是同步的。可以在 `OnPaid` 里开 goroutine 去做慢的事，或者把事件放进 channel，由后台 goroutine 慢慢处理。改成异步后，要另外考虑执行顺序和失败重试。

**Q**：多个 goroutine 同时 `Subscribe` 和 `Pay` 安全吗？

不安全，切片会被并发读写。用 `sync.RWMutex` 保护观察者列表；通知时先在锁内复制一份列表，再在锁外逐个调用，避免观察者在回调里订阅时死锁（见[附录 A](../../appendix/concurrency.md)）。

**Q**：一个观察者出错或 panic，后面的还能收到通知吗？

panic 会中断整个循环，后面的都收不到。需要隔离时，在循环里对每次调用 `recover`；如果需要汇报错误，让接口返回 `error`，再用 `errors.Join` 收集。

**Q**：忘了退订会怎样？

主题一直引用着观察者，它就不会被 GC 回收，也会一直收到通知，这是常见的内存泄漏。订阅时返回取消函数，在观察者用完时调用它，或把它和 `defer` 放在一起。

## 动手试试

1. **让函数也能订阅 Shop**：在包外定义 `type ObserverFunc func(observer.Order)` 并实现 `OnPaid`，再用它订阅一个「金额 ≥ 1000 时打印 `风控：大额订单 <ID>`」的回调。

    提示：仿照 13.3 的 `DiscountFunc`，可直接在本章 Playground 里完成。检验：`Pay(Order{ID: "A002", Amount: 1500})` 在「支付成功：A002」后多一行「风控：大额订单 A002」，`Shop` 没改。

2. **一个观察者 panic，其余受连累**：依次订阅 `Stock{}`、一个 `OnPaid` 里直接 `panic` 的观察者、`Points{}`，然后 `Pay`；再写一个包装类型 `Safe{ob}`，在转发 `OnPaid` 前 `defer recover()`，把中间那个包起来重试。

    提示：可直接在本章 Playground 里完成，第一次 `Pay` 要包在带 `recover` 的函数里才能继续往下跑。检验：第一次只打印到「库存」，「积分」丢了；包上 `Safe` 后「积分：+29」照常出现。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/behavioral/observer/exercise_test.go)
