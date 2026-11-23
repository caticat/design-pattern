# 23.2 Go 实现

## 三步写完

分拣程序先读规则卡，再把它拼成可以反复使用的表达式树。

![三步：吉祥物拿起印着玻璃杯、雪花和闪电图标的规则卡；它照卡片把玻璃杯叶子接到根积木左口，把雪花与闪电接到右侧的双口子积木；它用同一套积木检查两个包裹，带玻璃杯和雪花的入托盘，只有雪花的走旁路](../../assets/images/interpreter-steps.webp){ .illus }

<p class="caption">规则卡只拼一次；换一个包裹时，积木不动，只换要检查的标签</p>

**第 1 步**：定义叶子。`Parcel` 保存当前包裹的标签；`tagExpr` 只查一个键。没有这个键时，Go 的 map 读出 `false`，不用另外写分支。

```go
--8<-- "codes/behavioral/interpreter/interpreter.go:leaves"
```

**第 2 步**：定义运算节点。`notExpr` 取反；`andExpr` 与 `orExpr` 用 Go 的短路求值（左边已足够决定结果，就不算右边）。

```go
--8<-- "codes/behavioral/interpreter/interpreter.go:operators"
```

**第 3 步**：把文本拼成树。`Parse` 先用 `go/parser.ParseExpr` 识别括号和优先级，再把结果转成前两步的节点。具体转换见 [23.3](tradeoffs.md)。

## 使用示例

`Parse` 失败时先报告错误，不拿无效规则分拣。解析成功后把规则放在循环外，每次只传入当前包裹。

```go
--8<-- "codes/behavioral/interpreter/example_test.go:example"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/3Vgjd6QYHQe){ .md-button .md-button--primary target=_blank }
