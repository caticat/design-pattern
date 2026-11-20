// Package strategy 演示策略模式：把一组可互换的算法封装起来，运行时按需切换。
package strategy

// --8<-- [start:interface]
// Discount 是优惠策略：输入原价，返回优惠后的价格（单位：元）。
type Discount interface {
	Apply(total int) int
}

// --8<-- [end:interface]

// --8<-- [start:concrete]
// NoDiscount 不打折。
type NoDiscount struct{}

func (NoDiscount) Apply(total int) int { return total }

// FullReduction 满减：满 Threshold 减 Off。
type FullReduction struct{ Threshold, Off int }

func (f FullReduction) Apply(total int) int {
	if total >= f.Threshold {
		return total - f.Off
	}
	return total
}

// Percent 折扣：Rate=80 表示打 8 折。
type Percent struct{ Rate int }

func (p Percent) Apply(total int) int { return total * p.Rate / 100 }

// --8<-- [end:concrete]

// --8<-- [start:context]
// Checkout 是收银台（上下文）：只依赖 Discount 接口，不关心具体怎么算。
type Checkout struct {
	discount Discount
}

// NewCheckout 创建收银台，并指定初始策略。
func NewCheckout(d Discount) *Checkout { return &Checkout{discount: d} }

// SetDiscount 在运行时切换策略。
func (c *Checkout) SetDiscount(d Discount) { c.discount = d }

// Pay 计算应付金额。
func (c *Checkout) Pay(total int) int { return c.discount.Apply(total) }

// --8<-- [end:context]

// --8<-- [start:func]
// DiscountFunc 让普通函数也能当作策略使用，
// 与标准库 http.HandlerFunc 是同一个技巧。
type DiscountFunc func(total int) int

func (f DiscountFunc) Apply(total int) int { return f(total) }

// --8<-- [end:func]
