# 代理模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 结构型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 中</span>
</div>

![代理模式：片库门口的售票窗口](../../assets/images/proxy.webp){ .cover }

!!! abstract ""

    想看片的人不进片库，只在窗口递卡；窗口先验卡，再进库取片，取过的片放在手边下次直接给。

    代理模式用一个**同接口的替身**挡在真实对象前面，**控制对它的访问**：先检查、晚创建、可复用。

## 本章内容

| 节 | 内容 |
|---|---|
| [8.1 问题与思路](problem.md) | 一次加载全部、到处写检查的麻烦，替身怎么解决 |
| [8.2 Go 实现](implementation.md) | 公共接口、代理结构体与延迟加载 |
| [8.3 代理的用途](uses.md) | 常见的几种代理、标准库中的例子、什么时候用 |
| [8.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 A 并发](../../appendix/concurrency.md) · [附录 D 怎么读类图](../../appendix/class-diagram.md)　**相关对比**：[包一层](../../compare.md#wrappers)
