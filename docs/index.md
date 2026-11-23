# Go 设计模式图解

用**图**建立直觉，用 **Go** 落地，用**测试**保证代码可运行。讲解 GoF 23 个经典设计模式。

![吉祥物 Gopher 用转换头解决插头不匹配的问题](assets/images/adapter-problem.webp){ .cover }

## 这本书适合谁

| 读者 | 怎么读 |
|---|---|
| 初学者 | 按章节顺序读：封面类比 → 问题与思路 → Go 实现 → 专题 → 小结 |
| 有经验的开发者 | 直接看每章的 Go 实现、专题（标准库实例、相近写法对比）和小结 Q & A |

前置知识：Go 基础语法（结构体、接口、方法）。

## 每章结构

一章讲一个模式，分成几个短页面，每页一两屏读完：

| 页面 | 内容 |
|---|---|
| 章首页 | 一张类比大图 + 一句话定义 |
| N.1 问题与思路 | 不用模式时哪里出错，模式怎么解决，结构图 |
| N.2 Go 实现 | 可运行代码、Example 测试、在线运行 |
| N.3 专题 | 与相近写法的对比、标准库实例、什么时候用 |
| N.4 小结 | 重点回顾与常见问题 |

## 目录

图例：:material-check: 已完成　:material-clock-outline: 编写中

### 创建型：怎么创建对象

| 章 | 模式 | 一句话 | 状态 |
|---|---|---|---|
| 1 | [单例 Singleton](creational/singleton/index.md) | 全局只有一个实例 | :material-check: |
| 2 | [工厂方法 Factory Method](creational/factory-method/index.md) | 让子类型决定创建什么 | :material-check: |
| 3 | [抽象工厂 Abstract Factory](creational/abstract-factory/index.md) | 成套地创建相关对象 | :material-check: |
| 4 | [建造者 Builder](creational/builder/index.md) | 一步步组装复杂对象 | :material-check: |
| 5 | [原型 Prototype](creational/prototype/index.md) | 复制已有对象来创建新对象 | :material-check: |

### 结构型：怎么组合对象

| 章 | 模式 | 一句话 | 状态 |
|---|---|---|---|
| 6 | [适配器 Adapter](structural/adapter/index.md) | 转换接口，让不兼容的类型一起工作 | :material-check: |
| 7 | [装饰器 Decorator](structural/decorator/index.md) | 层层包装，动态叠加功能 | :material-check: |
| 8 | [代理 Proxy](structural/proxy/index.md) | 用替身控制对真实对象的访问 | :material-check: |
| 9 | [外观 Facade](structural/facade/index.md) | 给复杂子系统一个简单入口 | :material-check: |
| 10 | [组合 Composite](structural/composite/index.md) | 把对象组成树，统一对待单个和整体 | :material-check: |
| 11 | [桥接 Bridge](structural/bridge/index.md) | 抽象与实现分开，各自变化 | :material-check: |
| 12 | [享元 Flyweight](structural/flyweight/index.md) | 共享细粒度对象，节省内存 | :material-check: |

### 行为型：对象之间怎么协作

| 章 | 模式 | 一句话 | 状态 |
|---|---|---|---|
| 13 | [策略 Strategy](behavioral/strategy/index.md) | 封装可互换的算法，运行时切换 | :material-check: |
| 14 | [观察者 Observer](behavioral/observer/index.md) | 状态变化时自动通知订阅者 | :material-check: |
| 15 | [模板方法 Template Method](behavioral/template-method/index.md) | 固定流程骨架，步骤可替换 | :material-check: |
| 16 | 迭代器 Iterator | 不暴露内部结构地遍历集合 | :material-clock-outline: |
| 17 | 责任链 Chain of Responsibility | 请求沿链传递，直到有人处理 | :material-clock-outline: |
| 18 | 命令 Command | 把请求封装成对象 | :material-clock-outline: |
| 19 | 状态 State | 状态改变时，行为跟着改变 | :material-clock-outline: |
| 20 | 备忘录 Memento | 保存并恢复对象状态 | :material-clock-outline: |
| 21 | 中介者 Mediator | 用中间人替代对象间的直接通信 | :material-clock-outline: |
| 22 | 访问者 Visitor | 不改类型，为其添加新操作 | :material-clock-outline: |
| 23 | 解释器 Interpreter | 为简单语言定义文法并解释执行 | :material-clock-outline: |

## 运行代码

```bash
git clone https://github.com/caticat/design-pattern.git
cd design-pattern
go test ./...
```

所有代码只依赖 Go 1.23+ 标准库。
