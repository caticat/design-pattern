# 19.4 小结

## 重点回顾

!!! takeaway "划重点"

    - 状态模式把各阶段的行为放到各自的类型里，机器负责委托。
    - 测试不只看能否切换，还要验证操作被拒绝后是否留在原状态。

```mermaid
sequenceDiagram
    actor Customer as 顾客
    participant Machine
    participant Paid as paid
    Customer->>Machine: 再投一枚硬币
    Machine->>Paid: insert(m)
    Paid-->>Machine: 已投币，退回硬币
    Machine-->>Customer: 已投币，退回硬币
    Customer->>Machine: 按出货键
    Machine->>Paid: press(m)
    Paid->>Machine: 切到 soldOut
    Paid-->>Machine: 已出货
    Machine-->>Customer: 已出货
```

<p class="caption">重复投币被拒绝后仍是已投币，原来买到的一罐还能取走</p>

## Q & A

**Q**：状态模式和策略模式都在换接口实现，怎么区分？

策略通常由调用方或配置选，用来决定一件事「怎么做」。状态随对象处理事件而变化，决定「现在能做什么、做完去哪」。

**Q**：状态一定要自己决定下一状态吗？

不一定。本例在状态方法里切换，读到动作就能看到去向；如果跨阶段的约束很多，也可由机器统一检查后再切换，让转换关系集中可查。

**Q**：真实售货机能放很多罐，难道每个库存数字都要一个类型？

不用。库存是数量，放在机器上；「能否继续卖」才是阶段。成功出货后减库存：还有货转回待投币，卖完才转售罄。

**Q**：并发调用 `Insert`、`Press` 安全吗？

不安全。本例只演示顺序调用。需要并发时，可用[互斥锁](../../appendix/concurrency.md)保护整个入口，让「读状态、执行、切换」一次做完，别的操作不能插进来。
