# 模式目录

23 个模式分为三类。可以按顺序读，也可以先找眼前的问题。

## 创建型：怎么创建对象

| 章 | 模式 | 一句话 |
|---|---|---|
| 1 | [单例 Singleton](creational/singleton/index.md) | 只创建一个实例，并提供统一的访问入口 |
| 2 | [工厂方法 Factory Method](creational/factory-method/index.md) | 在固定流程中替换创建方式 |
| 3 | [抽象工厂 Abstract Factory](creational/abstract-factory/index.md) | 成套地创建相关对象 |
| 4 | [建造者 Builder](creational/builder/index.md) | 一步步组装复杂对象 |
| 5 | [原型 Prototype](creational/prototype/index.md) | 复制已有对象来创建新对象 |

## 结构型：怎么组合对象

| 章 | 模式 | 一句话 |
|---|---|---|
| 6 | [适配器 Adapter](structural/adapter/index.md) | 转换接口，让不兼容的类型一起工作 |
| 7 | [装饰器 Decorator](structural/decorator/index.md) | 层层包装，动态叠加功能 |
| 8 | [代理 Proxy](structural/proxy/index.md) | 用替身控制对真实对象的访问 |
| 9 | [外观 Facade](structural/facade/index.md) | 给复杂子系统一个简单入口 |
| 10 | [组合 Composite](structural/composite/index.md) | 把对象组成树，统一对待单个和整体 |
| 11 | [桥接 Bridge](structural/bridge/index.md) | 把两个可变维度拆开，各自变化 |
| 12 | [享元 Flyweight](structural/flyweight/index.md) | 共享对象中可复用的部分，节省内存 |

## 行为型：对象之间怎么协作

| 章 | 模式 | 一句话 |
|---|---|---|
| 13 | [策略 Strategy](behavioral/strategy/index.md) | 让同一件事可以换一种算法来做 |
| 14 | [观察者 Observer](behavioral/observer/index.md) | 一处变化，通知所有订阅者 |
| 15 | [模板方法 Template Method](behavioral/template-method/index.md) | 固定流程骨架，替换个别步骤 |
| 16 | [迭代器 Iterator](behavioral/iterator/index.md) | 不暴露内部结构地遍历集合 |
| 17 | [责任链 Chain of Responsibility](behavioral/chain-of-responsibility/index.md) | 请求沿链传递，直到有人处理或到达链尾 |
| 18 | [命令 Command](behavioral/command/index.md) | 把操作封装起来，便于排队或稍后执行 |
| 19 | [状态 State](behavioral/state/index.md) | 状态改变时，行为跟着改变 |
| 20 | [备忘录 Memento](behavioral/memento/index.md) | 保存并恢复对象状态 |
| 21 | [中介者 Mediator](behavioral/mediator/index.md) | 由一个中间人协调对象之间的交互 |
| 22 | [访问者 Visitor](behavioral/visitor/index.md) | 让已有类型接待访问者，以便添加新操作 |
| 23 | [解释器 Interpreter](behavioral/interpreter/index.md) | 组合并解释一门简单语言中的表达式 |
