# 附录 E &nbsp; 术语表

各章反复出现的词，一句话说明，并指向讲得最清楚的地方。

## 模式通用的角色名

| 术语 | 一句话 | 见 |
|---|---|---|
| GoF | 《设计模式》一书的四位作者（Gang of Four），本书的 23 个模式出自该书 | [模式目录](../patterns.md) |
| 调用方 / 客户端（Client） | 使用模式的那段代码；模式的目标通常是让它少改或不改 | [抽象工厂 3.1](../creational/abstract-factory/problem.md) |
| 接口 / 抽象 X | 只规定「能做什么」的类型，例如抽象产品、抽象处理者 | [附录 B](interfaces.md) |
| 具体 X | 实现了接口的普通类型，例如具体策略、具体工厂 | [策略 13.1](../behavioral/strategy/problem.md) |
| 上下文（Context） | 持有接口并在需要时调用它的一方；具体含义因模式而异 | [策略 13.1](../behavioral/strategy/problem.md)、[享元 12.1](../structural/flyweight/problem.md) |
| 委托 | 自己不做，转交给持有的另一个对象去做 | [状态 19.1](../behavioral/state/problem.md) |
| 扩展点 | 流程中留出的、可由外部替换的一步 | [工厂方法 2.1](../creational/factory-method/problem.md) |
| 钩子 | 可选的扩展点：实现了就调用，不实现就跳过 | [模板方法 15.2](../behavioral/template-method/implementation.md) |
| 注册表 | 用 `map` 按名字登记实现，新增一种只需登记一次 | [工厂方法 2.3](../creational/factory-method/registry.md) |
| 中间件 | 包住 `http.Handler` 再返回 `http.Handler` 的函数，是装饰器在 Go 中最常见的形态 | [装饰器 7.3](../structural/decorator/middleware.md) |
| 延迟加载 | 第一次用到时才创建，而不是提前创建 | [单例 1.3](../creational/singleton/tradeoffs.md)、[代理 8.3](../structural/proxy/uses.md) |
| 浅拷贝 / 深拷贝 | 只复制一层（引用仍共用） / 连引用指向的数据一起复制 | [原型 5.3](../creational/prototype/copy.md) |

## 常用的 Go 基础

| 术语 | 一句话 | 见 |
|---|---|---|
| 隐式实现接口 | 类型拥有接口要求的全部方法，就实现了接口，不用声明 | [附录 B](interfaces.md) |
| 指针接收者 | 方法写在 `*T` 上时，只有 `*T` 实现接口，`T` 不算 | [附录 B](interfaces.md) |
| 类型断言 | `v, ok := x.(T)`：从接口值里取出具体类型 | [附录 B](interfaces.md) |
| 函数值 / 方法值 | 函数可以像变量一样传递；`obj.Method` 不加括号就是一个绑定了 `obj` 的函数值 | [附录 C](callbacks.md) |
| 回调 | 交给别人、由别人在合适时机调用的函数 | [附录 C](callbacks.md) |
| 闭包 | 引用了外层变量的匿名函数；捕获的是变量本身，不是当时的值 | [附录 C](callbacks.md) |
| 函数类型 | `type F func(...)`，给它加方法后，普通函数也能实现接口（如 `http.HandlerFunc`） | [附录 C](callbacks.md) |
| 结构体嵌入 | 把一个类型直接写进结构体，它的方法会被「提升」到外层；不是继承 | [适配器 6.4](../structural/adapter/summary.md) |
| 零值 | 变量未赋值时的默认值（`0`、`""`、`nil`）；好的类型零值可直接用 | [建造者 4.3](../creational/builder/options.md) |
| goroutine / 并发安全 | Go 的轻量线程；多个 goroutine 同时访问也不出错，才叫并发安全 | [附录 A](concurrency.md) |
| UML 类图 | 用方框和连线表示类型及其关系 | [附录 D](class-diagram.md) |
