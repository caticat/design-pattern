# 工厂方法模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 创建型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 高</span>
</div>

![工厂方法模式：吉祥物按下自动售货机的按钮，一杯饮料从出货口送出](../../assets/images/factory-method.webp){ .cover }

!!! abstract ""

    售货机的出货流程不变，换一个制作模块，就能卖另一种饮料。

    工厂方法模式（Factory Method）把**创建步骤留给可替换的实现**。Go 没有子类重写，可以把创建函数传给使用产品的流程。

## 本章内容

| 节 | 内容 |
|---|---|
| [2.1 问题与思路](problem.md) | 到处 `new` 具体类型有什么问题，工厂函数怎么解决 |
| [2.2 Go 实现](implementation.md) | 产品接口、工厂函数、调用方三步写完 |
| [2.3 注册表与标准库](registry.md) | 新增类型不改工厂、标准库中的例子、什么时候用 |
| [2.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 B 接口与方法](../../appendix/interfaces.md) · [附录 C 函数值与回调](../../appendix/callbacks.md) · [附录 A 并发](../../appendix/concurrency.md)　**相关对比**：[创建对象](../../compare.md#creators)
