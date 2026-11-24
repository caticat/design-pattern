# 单例模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 创建型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 高</span>
</div>

![单例模式：全城只有一座钟楼](../../assets/images/singleton.webp){ .cover }

!!! abstract ""

    家家户户的钟各走各的，只有钟楼的时间人人认同。

    单例模式保证一个类型**只有一个实例**，并提供一个全局访问点。

## 本章内容

| 节 | 内容 |
|---|---|
| [1.1 问题与思路](problem.md) | 「用到时再创建」为什么在并发下出错，单例怎么解决 |
| [1.2 Go 实现](implementation.md) | `sync.Once`、`sync.OnceValue` 与几种常见写法 |
| [1.3 单例与全局变量](tradeoffs.md) | 两者的区别、标准库中的例子、什么时候该用 |
| [1.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 A 并发](../../appendix/concurrency.md)　**相关对比**：[创建对象](../../compare.md#creators)
