# 4.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 建造者 = **分步设置参数，最后 Build 一次创建**；参数有名字、可省略。
    - 校验集中在 Build：组合不合理就不创建，调用方拿不到半成品。
    - Go 中可选参数多时常用函数式选项；需要分多步组装时用建造者。

## Q & A

**Q**：必填参数应该放在哪？

放进 `NewServerBuilder(addr)` 或 `NewServer(addr, ...)` 的普通参数里，漏写会编译失败。只有可选参数才放进设置方法或选项。

**Q**：GoF 里的指挥者（Director）去哪了？

指挥者按固定顺序调用建造步骤，用来复用常见的配置组合。Go 中写一个普通函数即可，例如 `func Production(addr string) *ServerBuilder`，在里面调好 `TLS()`、`Port(443)` 再返回，调用方还能继续改。

**Q**：建造者和抽象工厂有什么区别？

抽象工厂关注「创建**哪一族**对象」，每个方法一步就返回一个产品；建造者关注「**一个**复杂对象怎么一步步组装」，最后才返回产品。

**Q**：一个建造者可以 Build 多次吗？

本例可以：`Build` 返回副本，改几项再 `Build` 就得到另一个 `Server`，互不影响（如果新增可变引用字段，要重新检查[复制与共享](../prototype/copy.md)）。但建造者本身不是并发安全的，不要让多个 goroutine 同时修改同一个建造者（见[附录 A](../../appendix/concurrency.md)）。

**Q**：函数式选项能报错吗？

本例的选项只修改字段，错误统一在 `NewServer` 末尾校验。如果某个选项需要自己检查（如解析证书文件），可以把类型改成 `func(*Server) error`，`NewServer` 遇到第一个错误就返回。

## 动手试试

1. **组合出一个新选项 `WithHTTPS`**：在包外写 `func WithHTTPS() builder.Option`，一次开启 TLS 并把端口设为 443。可直接在本章 Playground 里完成。

    提示：返回的函数里依次调用 `builder.WithTLS()(s)` 和 `builder.WithPort(443)(s)`。检验：`builder.NewServer("localhost", WithHTTPS())` 输出 `localhost:443 timeout=30s maxConns=100 tls=true`，`NewServer` 一行没改；在它后面再加 `builder.WithPort(80)`，仍会被校验拦下。

2. **把建造者当模板复用会怎样**：先 `base := builder.NewServerBuilder("localhost").Timeout(5 * time.Second)`，再分别 `a, _ := base.Port(8080).Build()` 和 `b, err := base.TLS().Build()`。可直接在本章 Playground 里完成。

    提示：设置方法改的是 `base` 指向的同一个建造者，`Build` 只复制产品。检验：`a` 为 `localhost:8080 ... tls=false`，`b` 为 `localhost:8080 ... tls=true`、`err` 为 `nil`——`b` 继承了给 `a` 设的端口；想要两份独立配置，就各自从 `NewServerBuilder` 开始。
