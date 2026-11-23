# 20.1 问题与思路

## 为什么旧稿也变了？

文本编辑器用一个字符串存标题、一个 `[]string` 存正文。用户想把标题从「周末计划」改成「周日计划」，把第一行从「去公园」改成「去书店」，还要能一次撤销这两处改动。

![左：吉祥物修改文稿，托盘上敞口信封的旧稿上方还画着两轮太阳，下方却跟屏幕一样变成了一本书；右：屏幕上画着一轮太阳和一本书，吉祥物从托盘取回独立封存的旧稿，刚打开的信封里仍是两轮太阳和一棵树](../../assets/images/memento-problem.webp){ .illus }

<p class="caption">旧稿不是指向当前文稿的另一条路，而是一份独立保管的记录</p>

最直觉的做法是把字段赋给两个变量，撤销时原样赋回。标题恢复了，正文却还是「去书店」：

```go
--8<-- "codes/behavioral/memento/bad.go:bad"
```

```mermaid
flowchart LR
    A[编辑器的行切片] --> C[同一块底层数组]
    B[旧稿的行切片] --> C
    C --> D[原地改第一行：两边都读到去书店]
```

<p class="caption">赋值只复制切片描述，没有复制装着各行的数组</p>

## 编辑器自己封存旧稿

==保存和恢复由编辑器负责，历史只保管快照==。编辑器知道哪些字段要一起撤销，也能在保存时复制数组；外部可以传递快照值，却不用知道它里面存的是切片还是别的结构。

```mermaid
classDiagram
    direction LR
    class Editor {
        +Save() Memento
        +Restore(m)
    }
    class Memento {
        -title string
        -lines []string
    }
    class History {
        +Checkpoint()
        +Undo() bool
    }
    History --> Editor : 请求保存或恢复
    History o-- Memento : 保管
    Editor ..> Memento : 创建和读取
```

<p class="caption">History 按顺序存取快照，不负责解释标题和正文</p>

| 角色 | 职责 | 本例中 |
|---|---|---|
| 原发器（Originator） | 生成快照并用它恢复自身状态 | `Editor` |
| 备忘录（Memento） | 封存需要恢复的状态 | `Memento` |
| 管理者（Caretaker） | 决定何时留档、用哪份快照撤销 | `History` |
