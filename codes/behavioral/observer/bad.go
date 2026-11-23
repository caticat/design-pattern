package observer

import "fmt"

// --8<-- [start:bad]
// 反例：支付成功后，直接调用每个相关模块。
func payBad(o Order) {
	fmt.Println("支付成功：" + o.ID)
	deductStock(o)
	sendSMS(o)
	addPoints(o)
	// 要加「发优惠券」：又得回来改 payBad。
	// 支付代码依赖所有下游模块，少调一个就出 bug。
}

// --8<-- [end:bad]

func deductStock(o Order) { Stock{}.OnPaid(o) }
func sendSMS(o Order)     { SMS{}.OnPaid(o) }
func addPoints(o Order)   { Points{}.OnPaid(o) }
