package decorator

// --8<-- [start:bad]
// 反例：所有配料都做成订单上的开关，每个方法里挨个 if。
type Order struct{ Milk, Cream bool }

func (o Order) Cost() int {
	cost := 15
	if o.Milk {
		cost += 3
	}
	if o.Cream {
		cost += 5
	}
	return cost
}

func (o Order) Description() string {
	desc := "美式"
	if o.Milk { // 和 Cost 里一样的 if，又写了一遍
		desc += " + 牛奶"
	}
	if o.Cream {
		desc += " + 奶油"
	}
	return desc
}

// 加一种「糖浆」：Order 加字段，Cost、Description 各加一个 if；
// 想要「双份牛奶」：bool 表达不了，得改成计数。
// --8<-- [end:bad]
