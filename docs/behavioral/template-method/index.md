# 模板方法模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 行为型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 中</span>
</div>

![模板方法模式：吉祥物在柜台后往三明治上盖顶层面包，柜台上三个三明治面包一样、馅料不同，三个机器人各自举着番茄、奶酪、生菜等着放进去](../../assets/images/template-method.webp){ .cover }

!!! abstract ""

    报表要导出成 CSV、Markdown 等多种格式。每种格式的流程都一样：检查数据、写表头、逐行写、写结尾，只有每一步写出来的样子不同。为每种格式各写一个完整函数，流程就被复制了很多份。

    模板方法模式把**流程固定在一处**，只把会变的步骤交出去，就像做三明治：面包的顺序不变，只换中间的馅。

## 本章内容

| 节 | 内容 |
|---|---|
| [15.1 问题与思路](problem.md) | 流程被复制多份，怎么只换会变的步骤 |
| [15.2 Go 实现](implementation.md) | 步骤接口、模板函数、钩子 |
| [15.3 Go 里的模板方法](stdlib.md) | 嵌入不是继承、标准库中的例子、什么时候用 |
| [15.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 D 怎么读类图](../../appendix/class-diagram.md)　**相关对比**：[换一种做法](../../compare.md#behaviors)
