# 代码规范

## 目录

```text
codes/<分类>/<模式>/
  <模式>.go          # 推荐实现（含经典写法与 Go 惯用写法）
  bad.go             # 反例（可编译，仅用于说明问题，注释标明「反例」）
  <模式>_test.go     # 单元测试
  example_test.go    # Example 测试：既是文档示例，又由 go test 校验输出
```

- 一个模式一个包，包名 = 模式英文名小写（如 `singleton`、`strategy`、`decorator`）。
- 模块路径：`github.com/caticat/design-pattern`。

## 约束

- Go 1.23+；**只用标准库**，不引入第三方依赖。
- 通过 `gofmt`、`go vet ./...`、`go test ./...`；涉及并发的包另跑 `go test -race`（需 CGO，环境不支持时可跳过并注明）。
- 注释用中文；导出的类型和函数必须有注释，实现接口的方法可省略（避免片段噪音）。
- 每个被正文引用的片段不超过 40 行；超出则拆片段。
- 示例业务选生活化场景（支付、咖啡、配置、通知等），不用 `Foo/Bar/Component` 这类抽象名。
- 金额用 `int`（单位：分/元），避免浮点误差干扰阅读。
- Example 测试必须带 `// Output:`，正文展示的运行结果以它为准。
