# 23.3 规则语言的边界

## 从文字到树：借用 Go 自己的解析器

规则写在配置里是一行文字，例如 `fragile && (cold || express)`。这种写法和 Go 的布尔表达式一模一样，所以不必自己写解析：`go/parser.ParseExpr` 把文字读成 Go 的语法树，`build` 再逐个节点翻译成我们的 `Expr`。括号、先与后或这些规矩都由 Go 处理好了。

```go
--8<-- "codes/behavioral/interpreter/interpreter.go:parse"
```

只翻译认识的节点：标签、括号、`!`、`&&`、`||`。数字、函数调用之类一律报错，不会把配置当成 Go 代码执行。

规则写错（比如少了括号）时 `Parse` 返回错误，调用方应拒绝这条规则，而不是把所有包裹都当成「不符合」。

```go
--8<-- "codes/behavioral/interpreter/interpreter.go:binary"
```

标准库里的同类例子是 `text/template`：`Parse` 把模板文字变成树，`Execute` 拿数据求值。思路一样，先解析、再反复执行。

## 什么时候用

<div class="grid" markdown>

!!! success "适合"

    - 规则要由别人（运营、配置）经常改，组合又多。
    - 运算种类少而稳定，比如只有与、或、非。

!!! failure "不适合"

    - 只有几条固定判断：直接写 `if` 更好读。
    - 只要「这几个标签都要有」：一张标签列表加 `for` 循环就够了。
    - 要变量、函数、循环：那已经是脚本语言，用现成的。

</div>
