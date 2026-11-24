# 命令模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 行为型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 高</span>
</div>

![吉祥物把任务卡放进托盘，机器人取出一张卡，按卡片要求打开台灯](../../assets/images/command.webp){ .cover }

!!! abstract ""

    要开灯，不必亲手按开关。把「打开这盏灯」写在任务卡上，先放进托盘，等机器人取走执行。

    命令模式（Command）把**一次请求封装成可执行的值**。提交方可以先存起来，到合适的时机再执行，不必知道设备如何工作。

## 本章内容

| 节 | 内容 |
|---|---|
| [18.1 问题与思路](problem.md) | 把立即调用变成可排队的命令 |
| [18.2 Go 实现](implementation.md) | 命令接口、设备操作和执行队列 |
| [18.3 Go 里的命令](stdlib.md) | 用闭包表示命令、什么时候用 |
| [18.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 C 函数值与回调](../../appendix/callbacks.md) · [附录 D 怎么读类图](../../appendix/class-diagram.md)　**相关对比**：[换一种做法](../../compare.md#behaviors) · [撤销](../../compare.md#undo)
