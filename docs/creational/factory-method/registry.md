# 2.3 注册表与标准库

## 新增类型不改工厂

2.2 的工厂里有一个 `switch`：每加一种通知方式，还是要改 `NewNotifier`。把 `switch` 换成「名字 → 工厂函数」的映射，就能做到**只加代码、不改工厂**。

![三步：售货机只有两个按钮、第三个空着，吉祥物拿着珊瑚色卡盒；把卡盒插进机器侧面；第三个按钮亮起，吉祥物拿到新饮料](../../assets/images/factory-method-register.webp){ .illus }

<p class="caption">新口味插一个卡盒就上架，售货机本身不用拆开</p>

每种通知方式在 `init`（包初始化时自动执行的函数）中登记自己：

```go
--8<-- "codes/creational/factorymethod/registry.go:registry"
--8<-- "codes/creational/factorymethod/registry.go:init"
```

新增微信通知时，`Create` 一行不改，甚至可以写在别的包里：

```go
--8<-- "codes/creational/factorymethod/example_test.go:example-register"
```

## 标准库中的工厂

| 写法 | 用法 | 说明 |
|---|---|---|
| 注册表 | `sql.Register` + `sql.Open("mysql", dsn)` | 驱动在 `init` 中登记；换数据库只改驱动名 |
| 注册表 | `image.RegisterFormat` + `image.Decode(r)` | 按文件头找到解码器，返回 `image.Image` |
| 传入工厂函数 | `hmac.New(sha256.New, key)` | `sha256.New` 本身就是工厂，`hmac` 只认 `hash.Hash` |
| 工厂函数返回接口 | `sha256.New() hash.Hash`、`errors.New` | 隐藏具体类型，只暴露接口 |

## 什么时候用

<div class="grid" markdown>

!!! success "适合"

    - 要在多个实现中选一个，且选择来自配置或输入
    - 创建过程有细节（地址、参数、默认值），不想散落各处
    - 需要插件式扩展：第三方包自己登记新类型

!!! failure "不适合"

    - 只有一个具体类型：直接 `NewXxx` 返回结构体
    - 创建很简单、只在一处使用
    - 要成套创建彼此配套的对象（用第 3 章的抽象工厂）

</div>
