# 7.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 装饰器 = **同一接口的层层包装**，每层只叠加一件事。
    - 用组合代替「每种组合一个类型」，n 种功能只需 n 个装饰器。
    - Go 中最常见的形式是 HTTP 中间件 `func(http.Handler) http.Handler` 和 `io.Reader` 包装。

## Q & A

**Q**：装饰器和适配器、代理有什么区别？

适配器**改变接口**，让不兼容的类型能一起用；装饰器**保持接口不变**，只增强功能。代理和装饰器结构几乎一样，但代理侧重**控制访问**（权限、延迟加载），通常只有一层；装饰器侧重**叠加功能**，常常多层嵌套。

**Q**：中间件可以不调用 `next`，这还是装饰器吗？

`RequireToken` 在缺少 token 时直接返回、不再往里调用，这是责任链的特点：某一环可以中断。HTTP 中间件同时具有装饰器和责任链两种特点。

**Q**：包了很多层之后，出错了怎么排查？

每层只做一件事，并起清楚的名字（如 `Logging`、`RequireToken`），调用栈里就能直接看出卡在哪一层。层数确实很多时，用 `Chain` 集中在一处组装，顺序一目了然。

**Q**：被包装后，`b.(Coffee)` 为什么断言失败？

`b` 现在是最外层的 `Cream`，不再是 `Coffee`。使用方应只依赖 `Beverage` 接口，不要依赖被包装在里面的具体类型。

**Q**：能用结构体嵌入 `type Milk struct{ Beverage }` 来写装饰器吗？

能。嵌入后，没有重写的方法会自动转发给里面那层，接口方法很多时可以少写代码。代价是忘记重写某个方法时不会报错，它会悄悄走到里面那层。

## 动手试试

1. **加一种配料：糖浆**：写装饰器 `Syrup`（加 2 元，描述追加「 + 糖浆」），包在 `decorator.Milk{B: decorator.Coffee{}}` 外面。可直接在本章 Playground 里完成。

    提示：照 `Milk` 写，持有一个 `decorator.Beverage` 字段，实现 `Cost` 和 `Description` 两个方法。检验：输出「美式 + 牛奶 + 糖浆 = 20 元」，`Coffee`、`Milk`、`Cream` 都没改。

2. **把中间件顺序反过来**：在 `ExampleChain` 里去掉 `req.Header.Set("X-Token", "demo")`，分别用 `Chain(hello, Logging(os.Stdout), RequireToken)` 和 `Chain(hello, RequireToken, Logging(os.Stdout))` 处理请求。可直接在本章 Playground 里完成。

    提示：用 `httptest.NewRecorder()` 接住响应，打印它的 `Code`。检验：两种顺序都返回 403，但只有前一种打印出 `-> GET /hi` 和 `<- /hi`——被拒的请求在后一种顺序下没进日志。
