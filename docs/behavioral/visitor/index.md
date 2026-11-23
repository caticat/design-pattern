# 访问者模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 行为型</span>
<span>:material-signal-cellular-2: 难度 进阶</span>
<span>:material-fire: 使用频率 低</span>
</div>

![博物馆里挂着画、立着雕塑；戴安全帽的吉祥物拿着空白展签，旁边的运输机器人推着画箱和木箱，两人沿同一条展线工作](../../assets/images/visitor.webp){ .cover }

!!! abstract ""

    同一批展品，策展人要写展签，运输人员要备箱子。画作和雕塑不用换一套，只要换一位来处理它们的人。

    访问者（Visitor）模式把**对一组不同类型对象的操作独立出来**，让新增操作不必修改这些对象的类型。

## 本章内容

| 节 | 内容 |
|---|---|
| [22.1 问题与思路](problem.md) | 写展签和备箱子，为什么总在辨认同一批展品 |
| [22.2 Go 实现](implementation.md) | 用 `Accept` 和 `Visit` 方法选中正确的处理逻辑 |
| [22.3 新增操作还是新增类型](tradeoffs.md) | 访问者的取舍与标准库里的树遍历 |
| [22.4 小结](summary.md) | 重点回顾与常见问题 |
