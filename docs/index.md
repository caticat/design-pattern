# Go 设计模式图解

用**图**建立直觉，用 **Go** 落地，用**测试**保证代码可运行。讲解 GoF 23 个经典设计模式。

![吉祥物 Gopher 用转换头解决插头不匹配的问题](assets/images/adapter.webp)

<p class="caption">设计模式就像工具箱里的转换头：把反复出现的问题，用一种经过验证的方式解决掉</p>

## 这本书适合谁

| 读者 | 怎么读 |
|---|---|
| 初学者 | 按顺序读正文：类比图 → 问题 → 方案 → 代码 |
| 有经验的开发者 | 直接展开每章末尾的「进阶」折叠块：Go 惯用写法、标准库实例、相似模式对比 |

## 每章结构

1. **一句话**：模式的意图
2. **类比图**：用生活场景建立直觉
3. **问题**：不用模式时，代码哪里难受
4. **方案**：结构图和角色
5. **Go 实现**：可运行代码，配 Example 测试
6. **什么时候用 / 小结 / 进阶**

## 目录

图例：:material-check: 已完成　:material-clock-outline: 编写中

### 创建型：怎么创建对象

| 模式 | 一句话 | 状态 |
|---|---|---|
| [单例 Singleton](creational/singleton.md) | 全局只有一个实例 | :material-check: |
| 工厂方法 Factory Method | 让子类型决定创建什么 | :material-clock-outline: |
| 抽象工厂 Abstract Factory | 成套地创建相关对象 | :material-clock-outline: |
| 建造者 Builder | 一步步组装复杂对象 | :material-clock-outline: |
| 原型 Prototype | 复制已有对象来创建新对象 | :material-clock-outline: |

### 结构型：怎么组合对象

| 模式 | 一句话 | 状态 |
|---|---|---|
| 适配器 Adapter | 转换接口，让不兼容的类型一起工作 | :material-clock-outline: |
| 桥接 Bridge | 抽象与实现分开，各自变化 | :material-clock-outline: |
| 组合 Composite | 把对象组成树，统一对待单个和整体 | :material-clock-outline: |
| [装饰器 Decorator](structural/decorator.md) | 层层包装，动态叠加功能 | :material-check: |
| 外观 Facade | 给复杂子系统一个简单入口 | :material-clock-outline: |
| 享元 Flyweight | 共享细粒度对象，节省内存 | :material-clock-outline: |
| 代理 Proxy | 用替身控制对真实对象的访问 | :material-clock-outline: |

### 行为型：对象之间怎么协作

| 模式 | 一句话 | 状态 |
|---|---|---|
| 责任链 Chain of Responsibility | 请求沿链传递，直到有人处理 | :material-clock-outline: |
| 命令 Command | 把请求封装成对象 | :material-clock-outline: |
| 解释器 Interpreter | 为简单语言定义文法并解释执行 | :material-clock-outline: |
| 迭代器 Iterator | 不暴露内部结构地遍历集合 | :material-clock-outline: |
| 中介者 Mediator | 用中间人替代对象间的直接通信 | :material-clock-outline: |
| 备忘录 Memento | 保存并恢复对象状态 | :material-clock-outline: |
| 观察者 Observer | 状态变化时自动通知订阅者 | :material-clock-outline: |
| 状态 State | 状态改变时，行为跟着改变 | :material-clock-outline: |
| [策略 Strategy](behavioral/strategy.md) | 封装可互换的算法，运行时切换 | :material-check: |
| 模板方法 Template Method | 固定流程骨架，步骤可替换 | :material-clock-outline: |
| 访问者 Visitor | 不改类型，为其添加新操作 | :material-clock-outline: |

## 运行代码

```bash
git clone https://github.com/caticat/design-pattern.git
cd design-pattern
go test ./...
```

所有代码只依赖 Go 1.23+ 标准库。
