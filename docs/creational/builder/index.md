# 建造者模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 创建型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 高</span>
</div>

![建造者模式：吉祥物照着只有图标的点单卡，在托盘上一层层叠汉堡，正放上番茄片](../../assets/images/builder.webp){ .cover }

!!! abstract ""

    点汉堡时一项项说：加芝士、多放生菜、不要洋葱，店员照单一层层做好，最后一次递给你。

    建造者模式把复杂对象的创建**拆成一步步设置**，最后由 **Build 统一校验、一次性交付**。

## 本章内容

| 节 | 内容 |
|---|---|
| [4.1 问题与思路](problem.md) | 参数一长调用就看不懂，分步设置怎么解决 |
| [4.2 Go 实现](implementation.md) | 产品、建造者、Build 校验三步写完 |
| [4.3 函数式选项与标准库](options.md) | Go 更常见的写法，三种写法怎么选，什么时候用 |
| [4.4 小结](summary.md) | 重点回顾、常见问题与动手试试 |

**前置知识**：[附录 C 函数值与回调](../../appendix/callbacks.md) · [附录 A 并发](../../appendix/concurrency.md)　**相关对比**：[创建对象](../../compare.md#creators)
