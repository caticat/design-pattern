# 11.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 桥接 = 把**两个独立变化的维度**拆成两边，用一个字段接起来。
    - 「抽象」一边管发什么，「实现」一边管怎么发，各自增加类型互不影响。
    - 类型数从 M × N 降到 M + N。

## Q & A

**Q**：「抽象」和「实现」这两个名字很难懂，有更直白的说法吗？

把「抽象」理解成**调用方直接用的那一边**，把「实现」理解成**被它委托去干底层活的那一边**。这里的「实现」和 Go 里「类型实现接口」不是一回事。

**Q**：桥接和策略写出来几乎一样，怎么区分？

看持有接口的那一边有没有自己的变体。策略里使用方通常只有一个，只是换算法；桥接里两边都有多个变体（`Normal`、`Urgent` × `SMS`、`Email`）。对比表见 [11.3](stdlib.md)。

**Q**：如果只有 `Urgent` 一种消息，还需要 `Notice` 接口吗？

不需要。抽象一边只有一个类型时，它就是一个普通结构体加一个接口字段。等第二种消息出现再抽出接口，Go 的隐式接口让这一步改动很小。

**Q**：运行时能换渠道吗？

可以。给 `Normal` 加一个 `SetSender(s Sender)` 方法即可；如果多个 goroutine 同时发送，需要加锁或用 `atomic.Pointer` 保护这个字段（见[附录 A](../../appendix/concurrency.md)）。

## 动手试试

1. **在抽象一边加一种「定时」消息**：写类型 `Scheduled`，字段 `At string` 和 `Sender bridge.Sender`，`Notify` 发出「【09:00 定时】周报已生成」这样的文字。可直接在本章 Playground 里完成。

    提示：`Scheduled` 实现 `bridge.Notice`，`Notify` 里拼好文字后交给 `Sender.Send`。检验：配 `bridge.Email{}` 输出「邮件：【09:00 定时】周报已生成」，换 `bridge.SMS{}` 输出「短信：…」，`SMS`、`Email` 没改。

2. **不改包，也能运行时换渠道**：写一个可切换的渠道 `Switch`（字段 `S bridge.Sender`，`Send` 转给 `S`），用 `bridge.NewNormal(sw)` 接上它，发一条后把 `sw.S` 换掉再发一条。可直接在本章 Playground 里完成。

    提示：`NewNormal` 只认 `Sender` 接口，传 `*Switch` 即可，`Send` 用指针接收者。检验：先输出「短信：值班交接」，换成 `bridge.Email{}` 后输出「邮件：值班交接」，`Normal` 没加任何方法。
