# 装饰器模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 结构型</span>
<span>:material-signal-cellular-2: 难度 进阶</span>
<span>:material-fire: 使用频率 高</span>
</div>

![装饰器模式：一杯咖啡层层加料](../../assets/images/decorator.webp){ .cover }

!!! abstract ""

    一杯美式，加牛奶，再加奶油：杯子还是那个杯子，价格一层层往上加。

    装饰器模式在**不修改原对象**的前提下，用一层层**包装**给它动态叠加新功能。

## 本章内容

| 节 | 内容 |
|---|---|
| [7.1 问题与思路](problem.md) | 「每种组合一个类型」为什么会爆炸，包装怎么解决 |
| [7.2 Go 实现](implementation.md) | 组件接口、装饰器结构体与几种写法 |
| [7.3 HTTP 中间件](middleware.md) | Go 中最常见的装饰器、标准库中的例子、什么时候用 |
| [7.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 D 怎么读类图](../../appendix/class-diagram.md)　**相关对比**：[包一层](../../compare.md#wrappers) · [树与遍历](../../compare.md#trees)
