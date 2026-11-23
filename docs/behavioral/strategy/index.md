# 策略模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 行为型</span>
<span>:material-signal-cellular-1: 难度 入门</span>
<span>:material-fire: 使用频率 高</span>
</div>

![策略模式：导航仪插上不同卡片，走不同路线到达同一目的地](../../assets/images/strategy.webp){ .cover }

!!! abstract ""

    去同一个地方，导航仪插上「步行」「骑行」「驾车」卡片，就换一种路线算法，目的地不变。

    策略模式把一组可互换的**算法**分别封装，让使用方在**运行时切换**。

## 本章内容

| 节 | 内容 |
|---|---|
| [13.1 问题与思路](problem.md) | 越写越长的 `switch` 哪里不好，怎么把算法抽出去 |
| [13.2 Go 实现](implementation.md) | 策略接口、具体策略、上下文与运行时切换 |
| [13.3 函数即策略](functions.md) | Go 中更轻量的函数写法、标准库中的例子、什么时候用 |
| [13.4 小结](summary.md) | 重点回顾与常见问题 |
