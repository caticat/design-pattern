# 2.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 工厂方法把**创建步骤留给可替换的实现**，使用产品的流程不用跟着改。
    - Go 可以直接传创建函数；返回接口本身不是工厂方法的判据。
    - 简单工厂集中选择逻辑；实现需要独立登记时，再用注册表。

## Q & A

**Q**：简单工厂、工厂方法、抽象工厂有什么区别？

简单工厂把「选哪个」集中在一个地方，例如 `NewNotifier` 的 `switch`。工厂方法让使用方的流程留出一个创建扩展点，例如 `SignUpWith` 的 `create` 参数。抽象工厂则提供**一组配套产品**的创建方法（第 3 章）。

**Q**：GoF 原版要求「让子类决定创建什么」，Go 没有子类怎么办？

原版的目的是让框架在不知道具体类型时也能创建对象。Go 里把工厂函数作为参数或字段传进去即可，例如 `hmac.New(sha256.New, key)`：`hmac` 不知道具体用哪种哈希，由调用方传入的函数决定。

**Q**：注册表为什么重复登记就 panic？

重复登记几乎总是两个包用了同一个名字，属于程序写错了。`init` 阶段 panic 会让程序一启动就失败，比运行中悄悄用错实现好排查。`sql.Register` 也是这样做的。

**Q**：程序运行中可以再登记新类型吗？

本例的 `factories` 是普通 map，并发读写会出错（见[附录 A](../../appendix/concurrency.md)）。`init` 在 `main` 之前顺序执行，只在这里登记最安全；运行中还要登记，就得像 `database/sql` 那样用互斥锁保护。

**Q**：工厂返回 `nil` 而不是错误，可以吗？

不推荐。调用方拿到 `nil` 接口后直接调用方法会 panic，出错位置离原因很远。返回 `(Notifier, error)`，让调用方在源头处理。

## 动手试试

1. **接入站内信**：写一个实现 `Notifier` 的 `Inbox` 类型，在 `init` 中用 `factorymethod.Register("inbox", ...)` 登记，再调用 `factorymethod.SignUpWith("inbox", "小明", factorymethod.Create)`。可直接在本章 Playground 里完成。

    提示：`Send` 打印 `[站内信] 发给 %s：%s`。检验：输出 `[站内信] 发给 小明：欢迎注册`，`Create` 和 `SignUpWith` 一行没改；而 `factorymethod.SignUp("inbox", "小明")` 仍返回 `未知的通知方式 "inbox"`，因为简单工厂 `NewNotifier` 的名单是写死的。

2. **让工厂返回 nil 看看会怎样**：给 `SignUpWith` 传一个返回 `(nil, nil)` 的创建函数。可直接在本章 Playground 里完成。

    提示：用 `defer func() { fmt.Println(recover()) }()` 接住 panic。检验：打印 `runtime error: invalid memory address or nil pointer dereference`，出错位置在 `SignUpWith` 内部的 `n.Send`，而不是写错的那个工厂；改成返回 `error` 后，错误在源头就能看到。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/creational/factorymethod/exercise_test.go)
