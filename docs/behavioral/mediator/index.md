# 中介者模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 行为型</span>
<span>:material-signal-cellular-2: 难度 进阶</span>
<span>:material-fire: 使用频率 中</span>
</div>

![安全帽吉祥物坐在一块结账面板前，面板有自取与配送拨杆、地址卡槽和提交灯；每个控件的连线都接到中央的协调盒，控件之间没有连线](../../assets/images/mediator.webp){ .cover }

!!! abstract ""

    点餐时切到配送，地址框就出现；填好地址，提交按钮才亮。像一块有中央控制盒的面板：开关和输入框报出变化，控制盒决定该更新哪一处。

    中介者（Mediator）模式把**一组对象之间的协作规则交给一个协调对象**，让这些对象不必彼此引用。

## 本章内容

| 节 | 内容 |
|---|---|
| [21.1 问题与思路](problem.md) | 清空配送地址，提交按钮为什么还亮 |
| [21.2 Go 实现](implementation.md) | 用回调把控件接到同一个结账协调器 |
| [21.3 协调逻辑放在哪](coordination.md) | 直接调用与集中协调的取舍 |
| [21.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 D 怎么读类图](../../appendix/class-diagram.md)　**相关对比**：[对象之间传消息](../../compare.md#messaging)
