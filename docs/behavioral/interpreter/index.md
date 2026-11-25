# 解释器模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 行为型</span>
<span>:material-signal-cellular-2: 难度 进阶</span>
<span>:material-fire: 使用频率 低</span>
</div>

![分拣台上，吉祥物用可拼接的规则积木检查带有玻璃杯和雪花贴纸的包裹，符合规则的包裹滑入托盘](../../assets/images/interpreter.webp){ .cover }

!!! abstract ""

    分拣员收到一张新规则卡：易碎包裹只有在冷藏或加急时才入托盘。他不用改造分拣台，只要照卡片拼好规则积木，再拿包裹逐项检查。

    解释器（Interpreter）模式把一条规则**拼成一棵小结构体组成的树**：叶子读一个数据，「与、或、非」节点合并子节点的结果；从根节点求值，就得到整条规则的答案。

## 本章内容

| 节 | 内容 |
|---|---|
| [23.1 问题与思路](problem.md) | 规则写死在代码里的麻烦，把规则拼成一棵树 |
| [23.2 Go 实现](implementation.md) | 四个小结构体，拼树、求值 |
| [23.3 规则语言的边界](tradeoffs.md) | 把配置里的文字变成树、什么时候用 |
| [23.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 D 怎么读类图](../../appendix/class-diagram.md)　**相关对比**：[树与遍历](../../compare.md#trees)
