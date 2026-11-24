# 桥接模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 结构型</span>
<span>:material-signal-cellular-2: 难度 进阶</span>
<span>:material-fire: 使用频率 中</span>
</div>

![桥接模式：左岸是几封信，右岸是小货车、纸飞机和信鸽，吉祥物拿着一封信走过桥，虚线从信连到每一种送信工具](../../assets/images/bridge.webp){ .cover }

!!! abstract ""

    发通知时有两件事在变：**发什么**（普通、紧急）和**怎么发**（短信、邮件）。把每种组合都写一遍，类型数会成倍增长。

    桥接模式把两件事拆成两边，各自扩展，再用一个字段把它们**接起来**：M 种加 N 种，而不是 M 乘 N 种。

## 本章内容

| 节 | 内容 |
|---|---|
| [11.1 问题与思路](problem.md) | 组合爆炸是怎么来的，拆成两边怎么解决 |
| [11.2 Go 实现](implementation.md) | 实现一边、抽象一边、用字段接起来 |
| [11.3 Go 里的桥接](stdlib.md) | slog 的 Logger 与 Handler、标准库中的例子、什么时候用 |
| [11.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 A 并发](../../appendix/concurrency.md) · [附录 D 怎么读类图](../../appendix/class-diagram.md)　**相关对比**：[换一种做法](../../compare.md#behaviors)
