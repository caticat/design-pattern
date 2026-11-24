# 15.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 模板方法 = **流程写在一处**，会变的步骤交给接口实现。
    - 可选步骤做成钩子：单独一个小接口，用类型断言判断有没有实现。
    - Go 没有继承，嵌入不会让「父类」调用到「子类」的方法；模板方法写成接收接口的函数。

## Q & A

**Q**：步骤只有一两个，也要定义接口吗？

不必。直接把步骤作为函数参数传进去更简单，比如 `fs.WalkDir` 的回调、`slices.SortFunc` 的比较函数。步骤多、而且总是成组出现时，再用接口。

**Q**：想让某些步骤有默认实现，怎么办？

写一个提供默认方法的结构体，让具体类型嵌入它，再只写想改的那几个方法。这样可行，是因为 `Export` 通过**接口**调用步骤，调用的是外层类型的方法；和 15.3 的陷阱不同，流程不在被嵌入的结构体里。

**Q**：具体类型能改变步骤的顺序吗？

不能，这正是模板方法的目的：顺序由模板函数保证。如果每种实现都想要不同的顺序，说明流程本身就不同，用[策略模式](../strategy/index.md)替换整个算法更合适。

**Q**：钩子除了可选接口，还有别的写法吗？

也可以在接口里放一个方法，让默认实现什么都不做。可选接口的好处是不实现的类型完全不用写它，标准库也常这样做：`io.Copy` 会检查参数是否实现了 `io.WriterTo`，实现了就走更快的路径。

## 动手试试

1. **新增 HTML 格式**：在包外写 `HTML` 类型，`Header` 返回 `<table>`，`Row` 返回 `<tr><td>张三</td><td>95</td></tr>` 这样的一行，再实现钩子 `Footer` 返回 `</table>`。

    提示：仿照示例里的 `Text`，可直接在本章 Playground 里完成。检验：`Export(HTML{}, rows)` 首行 `<table>`、末行 `</table>`；传入空切片仍只输出「（无数据）」，`Export` 没改。

2. **亲手踩一次嵌入陷阱**：在包外写 `Base`（实现 `Header`、`Row`，再加方法 `Run(rows)`，内部调用 `templatemethod.Export(b, rows)`），让 `Fancy` 嵌入 `Base` 并只重写 `Header`；分别打印 `Fancy{}.Run(rows)` 和 `templatemethod.Export(Fancy{}, rows)`。

    提示：可直接在本章 Playground 里完成。检验：前者表头仍是 `Base` 的，后者才是 `Fancy` 的——流程挂在被嵌入的类型上就会掉进 15.3 的陷阱，流程放在接收接口的函数里才不会。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/behavioral/templatemethod/exercise_test.go)
