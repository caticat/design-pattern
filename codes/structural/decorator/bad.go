package decorator

// --8<-- [start:bad]
// 反例：用「组合出一个新类型」的方式扩展功能。
// 2 种配料就要 3 个新类型，n 种配料最多要 2ⁿ-1 个，这就是「类爆炸」。

// CoffeeWithMilk 是加牛奶的美式。
type CoffeeWithMilk struct{}

func (CoffeeWithMilk) Cost() int { return 15 + 3 }

// CoffeeWithCream 是加奶油的美式。
type CoffeeWithCream struct{}

func (CoffeeWithCream) Cost() int { return 15 + 5 }

// CoffeeWithMilkAndCream 是加牛奶和奶油的美式。
type CoffeeWithMilkAndCream struct{}

func (CoffeeWithMilkAndCream) Cost() int { return 15 + 3 + 5 }

// --8<-- [end:bad]
