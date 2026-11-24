# 6.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 适配器 = **转换接口**，让已有类型满足调用方期望的接口。
    - 适配器只做翻译（方法、单位、错误表示），不放业务规则。
    - 业务代码依赖**自己定义**的接口，第三方的差异关在适配器里。

## Q & A

**Q**：适配器和装饰器、代理、外观有什么区别？

适配器把已有接口**转换**成另一个已有接口；装饰器、代理接口不变，外观为子系统**新设计**一个简化入口。详见[易混模式对比](../../compare.md#wrappers)。

**Q**：目标接口应该照抄第三方 SDK 的方法吗？

不要。目标接口由使用方按自己的需要定义，方法越少越好。这也是 Go 的惯用法：接口定义在使用它的包里，而不是实现它的包里。这样 SDK 怎么变，影响都止于适配器。

**Q**：书上的「类适配器」在 Go 里怎么写？

类适配器依靠继承，Go 没有继承。用结构体嵌入 `type WalletAdapter struct{ *WalletClient }` 可以近似，但 `Transfer` 等被适配者的方法也会暴露给使用方。一般用本章持有字段的「对象适配器」即可。

## 动手试试

1. **再接一家银行卡支付**：自己写一个「第三方」类型 `CardClient`，方法为 `Charge(amount float64, ref string) bool`（金额以元计，返回是否成功），再写 `CardAdapter` 让它满足 `adapter.Payer`。可直接在本章 Playground 里完成。

    提示：仿照 `WalletAdapter`，在 `Pay` 里把 `int` 元转成 `float64`、把 `false` 转成 `error`，并加一行 `var _ adapter.Payer = CardAdapter{}`。检验：`adapter.Checkout(CardAdapter{}, "B001", 50)` 输出「支付成功」，`Checkout` 和 `WalletAdapter` 一行没改。

2. **观察业务规则放进适配器的后果**：写一个 `StrictWalletAdapter`，在调用钱包前先拒绝超过 1000 元的订单，然后把它换回 `adapter.WalletAdapter`。可直接在本章 Playground 里完成。

    提示：`Pay` 里先判断 `yuan > 1000` 返回错误，否则转调 `adapter.WalletAdapter{Client: a.Client}.Pay`。检验：同一笔 2000 元订单，用 `StrictWalletAdapter` 返回「支付失败: 超过 1000 元」，换回 `WalletAdapter` 却「支付成功」——限额随适配器一起丢了。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/structural/adapter/exercise_test.go)
