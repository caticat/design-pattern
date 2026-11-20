package strategy_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/strategy"
)

func Example() {
	c := strategy.NewCheckout(strategy.NoDiscount{})
	fmt.Println("原价:", c.Pay(300))

	c.SetDiscount(strategy.FullReduction{Threshold: 200, Off: 30})
	fmt.Println("满200减30:", c.Pay(300))

	c.SetDiscount(strategy.Percent{Rate: 80})
	fmt.Println("8折:", c.Pay(300))

	// 临时规则：直接用函数，不必新建类型。
	c.SetDiscount(strategy.DiscountFunc(func(total int) int {
		return max(total-100, 0)
	}))
	fmt.Println("立减100:", c.Pay(300))
	// Output:
	// 原价: 300
	// 满200减30: 270
	// 8折: 240
	// 立减100: 200
}
