# 备忘录模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 行为型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 中</span>
</div>

![吉祥物坐在文本编辑器前，屏幕上方用一轮太阳代表新标题，下方用书本代表新正文；一份装着旧标题和旧正文的密封信封放在撤销托盘上](../../assets/images/memento.webp){ .cover }

!!! abstract ""

    改周末计划前，把当前稿装进信封交给撤销托盘。后悔时把信封交还编辑器，由编辑器拆开恢复。托盘只管保管，不用知道信封里有哪些字段。

    备忘录模式（Memento）让对象在**不公开内部状态**的前提下保存快照，并在需要时恢复到保存时的样子。

## 本章内容

| 节 | 内容 |
|---|---|
| [20.1 问题与思路](problem.md) | 为什么存一份切片还不够 |
| [20.2 Go 实现](implementation.md) | 拷贝快照、保管快照、撤销编辑 |
| [20.3 撤销历史怎么存](undo-history.md) | 完整快照的成本、增量与保留上限 |
| [20.4 小结](summary.md) | 重点回顾与常见问题 |
