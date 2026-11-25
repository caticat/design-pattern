# 11.3 Go 里的桥接

## slog：Logger 与 Handler

Go 1.21 的 `log/slog` 就是一座桥：`Logger` 管「记什么」（`Info`、`Warn`、`With` 附加字段），`Handler` 管「怎么写」（文本、JSON、发到远端）：

```go
--8<-- "codes/structural/bridge/stdlib.go:slog"
```

`Logger` 的方法和 `Handler` 的种类各自增加，互不牵连；自己写一个 `Handler`，所有用 `Logger` 的代码立刻能用上。

## 标准库中的桥接

| 位置 | 抽象（调用方用的） | 实现（可替换的） |
|---|---|---|
| `log/slog` | `*slog.Logger` | `slog.Handler`：`TextHandler`、`JSONHandler` |
| `database/sql` | `*sql.DB`：查询、事务、连接池 | `driver.Driver`：MySQL、PostgreSQL、SQLite 等驱动 |
| `net/http` | `*http.Client`：重定向、Cookie、超时 | `http.RoundTripper`：真正发出请求的那一层 |

`database/sql` 最典型：连接池、事务这些逻辑只写一次，新数据库只需实现驱动接口。

## 桥接、适配器与策略

三者都是「持有一个接口字段」，区别在用意：

| | 桥接 | 适配器 | 策略 |
|---|---|---|---|
| 时机 | **设计时**就拆成两边 | 事后补救，接口对不上 | 设计时预留 |
| 解决什么 | 两个维度都要扩展 | 把已有类型转成要的接口 | 一个算法可以替换 |
| 两边 | 抽象一边**也有多个变体** | 只有一个适配器 | 使用方通常只有一个 |

## 什么时候用

从简单到复杂，够用就停：

| 情况 | 写法 |
|---|---|
| 消息不用记住渠道 | 直接传参：`Notify(s Sender, msg string)`，连字段都不要 |
| 消息创建时定好渠道，之后反复用 | 放一个 `Sender` 字段，这就是桥接 |
| 只有一边在变 | 一个接口就够了，谈不上桥接 |

<div class="grid" markdown>

!!! success "适合"

    - 类型按两个独立维度变化，组合数在增长
    - 高层逻辑和底层实现想分别演进，如驱动、渲染后端
    - 想在运行时切换底层实现

!!! failure "不适合"

    - 只有一个维度在变，一个接口就够了
    - 两个维度其实紧密相关，拆开后要互相暴露大量细节
    - 组合只有两三种，且不会再增加

</div>
