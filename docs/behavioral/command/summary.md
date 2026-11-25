# 18.4 小结

## 重点回顾

!!! takeaway "划重点"
    - 命令把「要做的事」存成一个值，先提交、后执行；Go 里最简单的命令就是 `func()`，队列就是 `[]func()`。
    - 队列通过统一接口执行；不关心哪种设备，也不需要按名称分派。
    - 只有一个动作时用函数；要保存元数据或支持撤销时再定义具体类型。

## Q & A

**Q**：把函数传进去，它会立刻执行吗？

不会。`fan.Start` 是绑定到该风扇的方法值；`fan.Start()` 才是直接调用。队列保存前者，等到 `Run` 时执行。

**Q**：用了命令模式就能撤销吗？

不能。以开灯命令为例，执行前灯可能已经亮了；撤销应恢复之前的状态，而不是一律关灯。有些操作无法撤销，只能做补偿操作。

**Q**：能把队列存到数据库，下次启动再执行吗？

本例不行：命令里有指向内存对象的指针，闭包也不是持久化格式。可以保存操作名、目标 ID 和参数，重启后再查找设备、重建命令。

**Q**：命令和策略有什么区别？

命令把「一件事」打包待执行，策略是「同一件事换算法」。详见[易混模式对比](../../compare.md#behaviors)。

## 动手试试

1. **一键「开工」宏命令**：在包外定义 `type Macro []command.Command`，让它的 `Execute` 依次执行每条命令；把 `Macro{command.LightOn{Lamp: lamp}, command.Func(fan.Start)}` 作为一条命令交给 `Queue`。

    提示：`Macro` 实现了 `Execute` 就是 `Command`，可直接在本章 Playground 里完成。检验：`Run` 前打印 `false false`，`Run` 后打印 `true true`；`Queue` 只看到一条命令，代码没改。

2. **写一个正确的撤销**：在包外写 `UndoableLightOn`（字段 `Lamp *command.Lamp` 和未导出的 `prev bool`），`Execute` 先记下 `Lamp.On` 再开灯，`Undo` 把 `Lamp.On` 恢复成 `prev`。

    提示：用指针接收者，否则 `prev` 记不住；可直接在本章 Playground 里完成。检验：灯原本关着，执行再撤销后为 `false`；灯原本亮着，执行再撤销后仍为 `true`，而不是一律关灯。

参考答案：[exercise_test.go](https://github.com/caticat/design-pattern/blob/master/codes/behavioral/command/exercise_test.go)
