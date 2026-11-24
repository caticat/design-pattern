# 9.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 外观 = **一组子系统的简单入口**，把步骤和顺序写在一处。
    - 调用方只依赖外观，子系统怎么变都不用改调用方。
    - Go 里外观常常就是一个包级函数，如 `os.ReadFile`、`http.Get`。

## Q & A

**Q**：外观会不会越写越大，变成什么都管的「上帝对象」？

会，如果把所有操作都往里塞。外观只放常用的组合流程；场景不同就拆成几个小外观（如「观影」和「K 歌」），罕见需求直接用子系统。

**Q**：用了外观，还能直接调用子系统吗？

能。外观不是权限控制，只是多一个方便的入口。需要精细控制时直接用子系统，见 [9.2](implementation.md) 的提示。

**Q**：外观要定义成接口吗？

一般不用，直接用结构体或函数。只有调用方测试时想替换掉整个外观，才在调用方这边定义一个只含所需方法的小接口。

**Q**：外观和中介者有什么区别？

外观是单向的：调用方找外观，外观调子系统，子系统不认识外观。中介者是双向的：一组对象都通过中介者互相通信（第 21 章）。

## 动手试试

1. **再做一个「K 歌」外观**：写结构体 `Karaoke`，持有 `facade.Lights`、`facade.Speaker`、`facade.Player`，方法 `Sing(song)` 依次调光到 50%、开音响、音量 60、播放。可直接在本章 Playground 里完成。

    提示：子系统都是可直接使用的零值结构体，字段不用初始化。检验：输出「灯光：调到 50%」→「音响：开机」→「音响：音量 60」→「播放器：播放 …」，`HomeTheater` 一行没改。

2. **测试时替换掉整个外观**：在调用方定义接口 `Theater`（含 `Watch(movie string)` 和 `End()`），写 `MovieNight(t Theater, movie string)` 先看再关，再写一个只记录调用的 `fakeTheater`。可直接在本章 Playground 里完成。

    提示：`fakeTheater` 把调用追加到 `calls []string`，方法用指针接收者。检验：传 `fakeTheater` 时不打印任何设备输出，`calls` 为 `[Watch 星际旅行 End]`；传 `facade.NewHomeTheater()` 时输出与本章示例相同。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/structural/facade/exercise_test.go)
