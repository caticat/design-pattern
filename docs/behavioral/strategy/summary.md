# 13.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 策略 = **把「怎么做」抽成可替换的对象**，上下文只管「什么时候做」。
    - 用接口取代 `switch`：新增算法只加代码，不改旧代码。
    - Go 中单方法策略优先用函数或函数类型，更轻量。

## Q & A

**Q**：使用方还是要写 `switch` 来选策略，问题不是没解决吗？

选择逻辑仍在，但只剩一处，而且只负责「选」，不再负责「算」。可以进一步用注册表收敛：`map[string]Discount{"percent": Percent{Rate: 80}, ...}`，新增策略时只加一行。

**Q**：只有两个分支，也要用策略吗？

不必。接口加几个类型，比一个 `if` 难读得多。等分支确实在增加、或需要运行时切换时再重构，13.1 的反例改成策略只需几分钟。

**Q**：策略和状态、模板方法有什么区别？

策略由**使用方**选、替换**整个算法**；状态由对象**自己**切换；模板方法只替换流程中的**某几步**。详见[易混模式对比](../../compare.md#behaviors)。

**Q**：策略里可以保存状态吗？

尽量不要。像 `Percent{Rate: 80}` 这样只读的配置没问题；如果策略在 `Apply` 中修改自己的字段，多个收银台共用同一个策略、或被多个 goroutine 同时调用时（见[附录 A](../../appendix/concurrency.md)），结果会互相影响。

## 动手试试

1. **叠加优惠**：在包外定义 `type Combo []strategy.Discount`，让它的 `Apply` 依次调用每个策略，再用 `Combo{FullReduction{Threshold: 200, Off: 30}, Percent{Rate: 80}}` 结算 300 元。

    提示：`Combo` 只要实现 `Apply(total int) int` 就是 `Discount`，可直接在本章 Playground 里完成。检验：`Pay(300)` 输出 216（300 → 270 → 216），`Checkout` 和已有优惠一行没改。

2. **让策略带上状态**：写一个「首单立减 50」策略 `type firstOrder struct{ used bool }`，用指针接收者在 `Apply` 里把 `used` 置为 `true`；把同一个 `&firstOrder{}` 交给两个 `NewCheckout`，各付 300 元。

    提示：可直接在本章 Playground 里完成。检验：第一个收银台得 250，第二个得 300——它的「首单」被另一个收银台用掉了，这就是 Q & A 里说的共享状态问题。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/behavioral/strategy/exercise_test.go)
