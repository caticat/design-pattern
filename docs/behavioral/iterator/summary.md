# 16.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 迭代器 = **集合自己提供遍历方式**，调用方只拿「下一个」，看不到内部结构。
    - Go 1.23 起首选 `iter.Seq`：返回一个接收 `yield` 的函数，调用方直接 `for range`。
    - `yield` 返回 `false` 必须立刻停止；需要暂停、继续时再用经典迭代器或 `iter.Pull`。

## Q & A

**Q**：`yield` 返回 `false` 之后还继续调用，会怎样？

程序直接 panic。`false` 表示调用方的循环已经 `break` 或 `return` 了，循环体不能再执行。所以每次调用 `yield` 都要检查返回值，像 16.2 第 3 步那样 `if !yield(...) { return }`。

**Q**：遍历的时候修改集合，会怎样？

结果取决于迭代器何时读下一项。本例取出最后一首歌后再追加，经典迭代器的游标已到 `nil`，看不到新歌；但 `All` 在 `yield` 返回后才读下一项，继续遍历时能看到。稳妥的做法是先用 `slices.Collect` 取一份快照，再遍历修改。

**Q**：遍历过程中可能出错（读文件、查数据库），怎么告诉调用方？

两种常见写法：用 `iter.Seq2[T, error]`，每一步都带一个错误；或者像 `bufio.Scanner` 那样，遍历完再调用 `Err()` 查一次。

**Q**：什么时候还需要经典迭代器？

当遍历要「暂停、稍后再继续」时，比如合并两个有序列表，要交替从两边各取一个。`for range` 一旦开始就要跑完或 `break`，这时用经典迭代器，或用 `iter.Pull` 把 `iter.Seq` 转成拉式，记得用 `defer stop()` 收尾，见 [16.3](stdlib.md)。

## 动手试试

1. **再写一个组合器 Take**：在包外写 `Take[T any](seq iter.Seq[T], n int) iter.Seq[T]`，只放行前 `n` 个元素，再和 `iterator.Filter` 叠起来用。

    提示：照 `Filter` 的写法，数够 `n` 个就 `return`，`yield` 返回 `false` 也要 `return`；可直接在本章 Playground 里完成。检验：歌单「晨光、雨夜、海风、雨巷、雨后」上，`slices.Collect(Take(Filter(p.All(), 以「雨」开头), 2))` 得到 `[雨夜 雨巷]`，`Playlist` 没改。

2. **忽略 yield 的返回值会怎样**：复制 `Filter` 写一个 `BadFilter`，把 `if keep(v) && !yield(v) { return }` 改成只调用 `yield(v)`；在歌单「晨光、雨夜、海风、雨巷」上 `for range BadFilter(...)`，取到第一首就 `break`。

    提示：可直接在本章 Playground 里完成；要让 `break` 之后还有符合条件的歌，否则看不到问题。检验：打印「雨夜」后程序 panic，信息为 `range function continued iteration after function for loop body returned false`。
