package observer_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/observer"
)

// --8<-- [start:example]
// Coupon 是新加的观察者：写在包外，Shop 的代码一行不改。
type Coupon struct{}

func (Coupon) OnPaid(o observer.Order) {
	if o.Amount >= 200 {
		fmt.Println("优惠券：满 200 送 20 元券")
	}
}

func Example() {
	shop := &observer.Shop{}
	shop.Subscribe(observer.Stock{})
	shop.Subscribe(observer.SMS{})
	shop.Subscribe(observer.Points{})
	shop.Subscribe(Coupon{})

	shop.Pay(observer.Order{ID: "A001", Amount: 299})
	// Output:
	// 支付成功：A001
	// 库存：扣减 A001 的商品
	// 短信：A001 已支付 299 元
	// 积分：+29
	// 优惠券：满 200 送 20 元券
}

// --8<-- [end:example]
