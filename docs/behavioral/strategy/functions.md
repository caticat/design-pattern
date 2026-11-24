# 13.3 函数即策略

## 单方法接口可以换成函数

Go 的函数可以当值传递。策略接口只有一个方法时，可以定义一个函数类型并让它实现接口，临时规则就不必新建类型：

```go
--8<-- "codes/behavioral/strategy/strategy.go:func"
```

`DiscountFunc(func(total int) int { ... })` 把普通函数转成 `Discount`，与标准库 `http.HandlerFunc` 是同一个技巧（见[附录 C](../../appendix/callbacks.md)）。

| 写法 | 适合 |
|---|---|
| 函数参数 `func(int) int` | 最轻量：例如写一个 `Apply(rule func(int) int, total int)`，直接把规则传进去；不满足已有接口时用 |
| 函数类型 + 方法（`DiscountFunc`） | 既想直接传函数，又要满足已有接口 |
| 接口 + 结构体（13.2 的写法） | 需要多个方法，或希望配置显式保存、便于比较时 |

表里的「无状态」指策略不保存、也不修改会变化的共享数据；固定参数用闭包捕获即可，不算状态。

## 标准库中的策略

| 例子 | 固定的部分 | 可替换的策略 |
|---|---|---|
| `slices.SortFunc(s, cmp)`（Go 1.21+） | 排序算法 | 比较函数 `cmp` |
| `sort.Sort(data)` | 排序算法 | `sort.Interface` 的比较与交换 |
| `bufio.Scanner.Split(bufio.ScanWords)` | 读取与缓冲 | 切分规则 `SplitFunc`：按行、按单词、按字节 |
| `http.Server` + `http.Handler` | 监听与连接管理 | 如何处理请求 |

## 什么时候用

<div class="grid" markdown>

!!! success "适合"

    - 同一件事有多种做法，且还会继续增加
    - 运行时按配置或用户选择切换算法
    - 想单独测试每一种算法

!!! failure "不适合"

    - 只有两三个分支，几乎不会变，`if` 更直白
    - 各分支差异很小，抽接口反而啰嗦
    - 算法需要大量访问上下文的内部数据

</div>
