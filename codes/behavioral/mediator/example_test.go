package mediator_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/mediator"
)

// --8<-- [start:example]
func Example() {
	c := mediator.NewCheckout()
	show := func(stage string) {
		fmt.Printf("%s：显示地址=%t，可提交=%t\n", stage, c.AddressVisible(), c.CanSubmit())
	}
	show("自取")
	c.SetDelivery(true)
	show("配送，地址未填")
	c.SetAddress("松林路 8 号")
	show("配送，已填地址")
	c.SetAddress("")
	show("配送，清空地址")
	c.SetDelivery(false)
	show("切回自取")
	// Output:
	// 自取：显示地址=false，可提交=true
	// 配送，地址未填：显示地址=true，可提交=false
	// 配送，已填地址：显示地址=true，可提交=true
	// 配送，清空地址：显示地址=true，可提交=false
	// 切回自取：显示地址=false，可提交=true
}

// --8<-- [end:example]
