# 写作规范与章节模板

## 写作原则

- **图先行**：先给类比图建立直觉，再讲结构和代码。
- **短**：正文（不含代码、图）1000–1500 字；段落不超过 4 行；能用表格不用段落。
- **一章一个模式**：不在正文展开其它模式，对比放进阶折叠块。
- **术语**：首次出现写「中文（English）」，之后只用中文。
- **不编造**：标准库 / 开源项目实例必须真实存在，写出包名和类型/函数名，可在 pkg.go.dev 查证。
- **不过时**：讲原理，不追新版本特性；用到 Go 1.21+ 特性时注明版本。
- **语气**：直接陈述，不用「我们来看看」「众所周知」等填充语。

## 章节模板

**样板章：`docs/creational/singleton.md`**。新章节照它的结构与组件写，模板如下。

文件：`docs/<分类>/<模式>.md`，分类为 `creational` / `structural` / `behavioral`。

```markdown
# X 模式

<div class="pattern-meta" markdown>
<span>:material-shape-outline: 创建型</span>          <!-- 创建型 / 结构型 / 行为型 -->
<span>:material-signal-cellular-1: 难度 入门</span>   <!-- 入门 cellular-1 / 进阶 cellular-2 / 较难 cellular-3 -->
<span>:material-fire: 使用频率 高</span>              <!-- 高 / 中 / 低 -->
</div>

!!! abstract "一句话"
    <意图，一句话，不超过 40 字，关键词加粗>

![<图名>](../assets/images/<模式>.webp)

<p class="caption">图注：用一句话把类比映射到模式角色</p>

## 问题

<2–4 行说明痛点> + 反例代码（来自 codes/.../bad.go）

!!! bug "<一句话说出后果>"
    <反例为什么出错 / 出什么问题>

## 方案

<核心思路 2–3 行，最关键的一句用 ==荧光笔==>

<Mermaid 类图或时序图>

| 角色 | 职责 | 本例中 |
|---|---|---|

## Go 实现

<代码片段，来自 codes/>，每段前一行说明在做什么；同一问题有多种写法时用标签页：

=== "经典写法"
    ...
=== "Go 1.21+ 惯用写法"
    ...

使用示例：<example_test.go 整文件，调用的是默认标签页中的写法>

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/TODO){ .md-button .md-button--primary target=_blank }

### X vs <最易混淆的写法>（可选，读者最常问的对比）

## 什么时候用

<div class="grid" markdown>

!!! success "适合"
    - 2–4 条

!!! failure "不适合"
    - 2–4 条

</div>

## 小结

!!! takeaway "划重点"
    - 3 条以内要点

??? note "进阶：Go 惯用写法"
??? note "进阶：标准库中的 X 模式"
??? note "进阶：X 与相似模式的区别"
??? warning "常见误用"
```

- 时序图（Mermaid sequenceDiagram）仅在协作顺序是理解关键时添加。
- 进阶折叠块按需取舍，至少保留「标准库实例」与「相似模式对比」之一。

### 视觉组件用法

| 组件 | 用途 | 限制 |
|---|---|---|
| `pattern-meta` 信息条 | 分类 / 难度 / 使用频率 | 每章必有，放标题下 |
| `!!! abstract` | 一句话定义 | 每章一个 |
| `!!! bug` | 反例的后果 | 只用在「问题」节 |
| `==文字==` 荧光笔 | 全章最关键的一句 | 每章 1–2 处 |
| `=== "标签"` | 多种等价写法 | 默认标签放最易懂的写法 |
| `!!! tip` / `!!! warning` | 易错点 | 正文中每章不超过 2 个 |
| `grid` + success/failure | 什么时候用 | 固定用法 |
| `!!! takeaway` | 小结（带吉祥物） | 每章一个 |

不新增自定义样式；确需新组件时先改样板章与本表，再推广。

## 代码引用

- 正文代码**全部**通过 snippets 引用 `codes/` 下的文件，不手写代码块（保证代码可编译、可测试）。
- 引用整文件：`--8<-- "codes/creational/singleton/singleton.go"`
- 引用片段：源码中用 `// --8<-- [start:name]` / `// --8<-- [end:name]` 标记，正文写 `--8<-- "path.go:name"`。
- 在线运行：每章一个 Playground 按钮（模板见上），一章只引用一个代码包。Playground 从公共仓库拉取模式包，故代码改动后须**先推送**，再运行 `python -I scripts/playground.py --update` 刷新链接并提交；不带参数为离线校验。

## 自检清单（提交人工审查前逐条确认）

- [ ] 一句话定义准确，类比图与图注对应
- [ ] 反例能说明痛点，方案确实解决了该痛点
- [ ] 类图角色与代码中的类型名一致
- [ ] 标准库实例真实存在（已核对包名/类型名）
- [ ] 字数在范围内，无填充语、无重复
- [ ] `go test ./...`、`go vet ./...`、`mkdocs build --strict`、`python -I scripts/playground.py` 全部通过
- [ ] 本地页面渲染正常（Mermaid、折叠块、图片、标签页），浅色 / 深色模式都看过
