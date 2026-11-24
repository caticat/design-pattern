# 6.3 函数适配器与标准库

## 把函数适配成接口

目标接口只有一个方法时，可以定义一个函数类型，把普通函数「转换」成这个接口（与 [13.3](../../behavioral/strategy/functions.md) 是同一个技巧）：

```go
--8<-- "codes/structural/adapter/adapter.go:func"
```

最常见的用途是测试：不调用真实的第三方，用一个函数伪造支付结果。

```go
--8<-- "codes/structural/adapter/example_test.go:example-func"
```

标准库的 `http.HandlerFunc` 就是这样写的，它的文档原话称自己为「adapter」。

## 标准库中的适配器

| 适配器 | 把 | 适配成 |
|---|---|---|
| `http.HandlerFunc(f)` | `func(http.ResponseWriter, *http.Request)` | `http.Handler` |
| `strings.NewReader(s)` | `string` | `io.Reader` |
| `bytes.NewReader(b)` | `[]byte` | `io.Reader` |
| `io.NopCloser(r)` | `io.Reader` | `io.ReadCloser`（`Close` 什么也不做） |
| `slog.NewLogLogger(h, level)` | `slog.Handler` | `*log.Logger`（Go 1.21+） |

例如某个函数要求 `io.Reader`，而手头只有一个字符串，用 `strings.NewReader(s)` 转一下就能传进去。

## 什么时候用

<div class="grid" markdown>

!!! success "适合"

    - 接入第三方库或旧系统，接口和自己的不一致
    - 想把业务与外部依赖隔离，方便替换和测试
    - 多个来源要统一成一个接口（多家支付、多种存储）

!!! failure "不适合"

    - 两边代码都归你管，直接改接口更简单
    - 只有一个调用点，近期不会替换，也不需要隔离测试
    - 要的是叠加功能而不是转换接口（用装饰器）
    - 差异大到要写很多业务逻辑才能对上

</div>
