# 附录 B &nbsp; 接口与方法

## 能调用方法，不一定能赋给接口

Go 不用写 `implements`。一个类型拥有接口要求的所有方法，且方法名和签名一致，就实现了这个接口。用[第 2 章](../creational/factory-method/implementation.md) 的 `Notifier` 举例：

```go
--8<-- "codes/appendix/interfaces/example_test.go:example"
```

| 方法的接收者 | 哪些类型拥有该方法 |
|---|---|
| `T` | `T` 和 `*T` |
| `*T` | 只有 `*T` |

`EmailNotifier.Send` 用的是指针接收者。变量 `email` 能直接调用 `email.Send`，是因为 Go 能帮它取地址；但赋给 `Notifier` 时不会自动取地址，所以要写 `&email`。

## 装着空指针的接口不等于 nil

接口值记录**动态类型和动态值**。`var n Notifier` 两者都没有，`n == nil`。若把一个值为 `nil` 的 `*EmailNotifier` 赋给 `n`，它仍记录着类型 `*EmailNotifier`，所以 `n != nil`；这不代表背后的对象可用。工厂失败时，返回 `nil, err`，别把空指针当作成功的结果。

## 要判断具体类型时

用 `email, ok := n.(*EmailNotifier)`：类型不匹配时 `ok` 是 `false`，不会 panic；匹配时仍可能拿到空指针。要区分多种类型时，用类型分支 `switch v := n.(type)`，如 [访问者的反例](../behavioral/visitor/problem.md)。只需要 `Send` 时，直接调用接口方法，不必检查具体类型。
