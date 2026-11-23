package chain_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/chain"
)

// --8<-- [start:example]
func Example() {
	c := chain.NewChain()

	// 调用方只认识链头，不知道后面有谁
	for _, amount := range []int{800, 3000, 20000, 80000} {
		fmt.Println(c.Handle(chain.Request{Name: "差旅", Amount: amount}))
	}
	// Output:
	// 组长批准：差旅 800 元
	// 经理批准：差旅 3000 元
	// 总监批准：差旅 20000 元
	// 无人能批：差旅 80000 元
}

// --8<-- [end:example]
