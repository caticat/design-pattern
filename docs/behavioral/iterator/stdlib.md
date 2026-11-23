# 16.3 Go 里的迭代器

## 迭代器可以组合

`iter.Seq[T]` 只是一个函数类型，所以可以写函数「包装」迭代器，像流水线一样一层层叠加：

```go
--8<-- "codes/behavioral/iterator/filter.go:filter"
```

`Filter(p.All(), keep)` 得到的还是 `iter.Seq[string]`，可以继续 `for range`，也可以再包一层。整个过程**边取边处理**，不会先把所有元素复制到一个切片里。

## 标准库中的迭代器

| 位置 | 形式 | 作用 |
|---|---|---|
| `slices.All`、`slices.Values`、`maps.Keys` | `iter.Seq` / `iter.Seq2` | 把切片、map 变成迭代器 |
| `slices.Collect`、`slices.Sorted` | 接收 `iter.Seq` | 把迭代器收集成切片 |
| `bufio.Scanner` | 经典：`Scan` + `Text` | 逐行读文件，读完用 `Err` 查错误 |
| `sql.Rows` | 经典：`Next` + `Scan` | 逐行读查询结果，用完要 `Close` |

例如 map 的键按顺序输出：`for _, k := range slices.Sorted(maps.Keys(m))`。

## 经典迭代器与 iter.Seq

| | 经典迭代器 | `iter.Seq`（Go 1.23） |
|---|---|---|
| 谁控制循环 | 调用方，想取才取（拉） | 集合，一个个推给 `yield`（推） |
| 遍历状态存在哪 | 迭代器对象的字段 | 函数里的局部变量 |
| 怎么用 | `for it.HasNext() { it.Next() }` | `for v := range seq` |
| 适合 | 暂停后继续、两个集合交替取 | 绝大多数遍历 |

两者可以互转：`iter.Pull(seq)` 把推式的 `iter.Seq` 变成拉式的 `next()` 函数。

## 什么时候用

<div class="grid" markdown>

!!! success "适合"

    - 集合内部结构复杂（链表、树、分页接口），不想让调用方知道
    - 同一个集合要支持多种遍历方式（正序、倒序、过滤）
    - 元素很多或要按需生成，不想一次全放进切片

!!! failure "不适合"

    - 集合本来就是切片或 map，直接 `range` 就行
    - 只遍历一处、结构也不会变，包一层反而多绕一步

</div>
