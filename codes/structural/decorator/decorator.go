// Package decorator 演示装饰器模式：在不修改原对象的前提下，层层包装、动态叠加功能。
package decorator

// --8<-- [start:component]
// Beverage 是饮品：所有咖啡和配料都实现这个接口。
type Beverage interface {
	Cost() int           // 价格（元）
	Description() string // 描述
}

// Coffee 是被装饰的基础对象：一杯美式。
type Coffee struct{}

func (Coffee) Cost() int           { return 15 }
func (Coffee) Description() string { return "美式" }

// --8<-- [end:component]

// --8<-- [start:decorators]
// Milk 是装饰器：包装任意 Beverage，加 3 元并追加描述。
type Milk struct{ B Beverage }

func (m Milk) Cost() int           { return m.B.Cost() + 3 }
func (m Milk) Description() string { return m.B.Description() + " + 牛奶" }

// Cream 是另一个装饰器：加 5 元。
type Cream struct{ B Beverage }

func (c Cream) Cost() int           { return c.B.Cost() + 5 }
func (c Cream) Description() string { return c.B.Description() + " + 奶油" }

// --8<-- [end:decorators]
