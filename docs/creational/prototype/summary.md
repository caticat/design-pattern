# 5.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 原型 = **复制一个配好的对象，再改不同的部分**，而不是从零构造。
    - `Clone` 要复制需要独立修改的数据；不可变数据可以共享。
    - 原型注册表登记和取出都用副本，原型本身从不交出去。

## Q & A

**Q**：原型模式和工厂方法有什么区别？

工厂方法靠写代码决定造什么；原型靠**复制已有对象**，新增种类只需在运行时登记原型。详见[易混模式对比](../../compare.md#creators)。

**Q**：能不能用 `encoding/gob` 或 `encoding/json` 编码再解码，实现通用深拷贝？

能用，但不推荐：比手写 `Clone` 慢得多，未导出的字段会丢失，函数、chan 等字段也无法编码。字段一多就手写 `Clone`，新增字段时记得同步更新，并用测试检查副本和原型互不影响（见仓库中 `prototype_test.go`）。

**Q**：要不要定义一个 `Cloner` 接口？

只有一种原型时不需要。如果注册表要保存多种类型的原型，可以定义 `type Prototype interface{ Clone() Prototype }`，让怪物、道具等类型都实现它；客户端只调用 `Clone`，不必知道具体类型。

**Q**：对象之间有循环引用怎么办？

例如 A 指向 B、B 又指向 A，逐字段复制会无限递归。需要用一个 map 记录「原对象 → 已复制的对象」，遇到已经复制过的直接复用。这种情况下应先考虑能否去掉循环引用。

**Q**：注册表能被多个 goroutine 同时使用吗？

本例不行：`Register` 会写 map。只在启动时登记、之后只调用 `Spawn` 的话没问题；运行中还要登记，就要加锁（见[附录 A](../../appendix/concurrency.md)）。

## 动手试试

1. **运行时派生一个新原型**：从注册表 `Spawn("goblin")` 取一只，改名为精英哥布林、`HP` 设为 60、追加技能「怒吼」，再用 `Register("elite", ...)` 登记成新原型。可直接在本章 Playground 里完成。

    提示：不用写新的构造函数，也不用改 `Registry`。检验：两次 `Spawn("elite")` 得到的怪物各改各的 `Skills` 互不影响；`Spawn("goblin")` 仍是 `哥布林 HP=30 技能=[挥砍] ...`。

2. **复制一支怪物小队**：把两只怪物放进 `team := []*prototype.Monster{...}`，用 `slices.Clone(team)` 复制后改副本第一只的 `HP`，再改用逐个调用 `Clone` 的方式复制一次对比。可直接在本章 Playground 里完成。

    提示：`slices.Clone` 只复制切片本身，元素是指针。检验：用 `slices.Clone` 时原小队的 `team[0].HP` 也变了；逐个 `Clone` 后原小队保持 `HP=30`。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/creational/prototype/exercise_test.go)
