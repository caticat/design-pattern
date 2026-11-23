# 17.2 Go 实现

## 三步写完

**第 1 步**：请求和统一接口。链上每一环都只需要实现 `Handle`。

```go
--8<-- "codes/behavioral/chain/chain.go:handler"
```

**第 2 步**：具体处理者。额度内就批；批不了，有下一环就交出去，没有就说明到头了。

```go
--8<-- "codes/behavioral/chain/chain.go:approver"
```

**第 3 步**：把链串起来。从链尾往前建，最后返回链头。

```go
--8<-- "codes/behavioral/chain/chain.go:chain"
```

![三步：吉祥物举起一段只有一个孔的滑道；吉祥物把三段孔从小到大的滑道首尾拼接，每个孔下站着端杯子的机器人；一枚硬币滚过小孔、掉进中孔下机器人的杯子，机器人欢呼，吉祥物竖起大拇指](../../assets/images/chain-of-responsibility-steps.webp){ .illus }

<p class="caption">每段滑道只管一个孔；首尾拼起来；硬币自己滚到装得下它的那一格</p>

!!! tip "改规则只改串链的地方"

    加一级「副总」、调整额度、让某个部门跳过经理，只需修改 `NewChain` 中组装链的代码（不同部门也可组装不同的链），`Approver` 和调用方一行不改。

## 使用示例

```go
--8<-- "codes/behavioral/chain/example_test.go:example"
```

[:material-play: 在 Go Playground 运行](https://go.dev/play/p/TODO){ .md-button .md-button--primary target=_blank }
