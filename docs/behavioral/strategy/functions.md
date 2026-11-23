# 13.3 函数即策略

## 单方法接口可以换成函数

Go 的函数可以当值传递。策略接口只有一个方法时，可以定义一个函数类型并让它实现接口，临时规则就不必新建类型：

```go
--8<-- "codes/behavioral/strategy/strategy.go:func"
```

`DiscountFunc(func(total int) int { ... })` 把普通函数转成 `Discount`，与标准库 `http.HandlerFunc` 是同一个技巧。

| 写法 | 适合 |
|---|---|
| 函数参数 `func(int) int` | 策略无状态，不需要满足某个接口 |
| 函数类型 + 方法（`DiscountFunc`） | 既想直接传函数，又要满足已有接口 |
| 接口 + 结构体（13.2 的写法） | 策略带配置参数（如满减门槛），或有多个方法 |

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
