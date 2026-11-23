package state_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/state"
)

// --8<-- [start:example]
func Example() {
	m := state.NewMachine()
	fmt.Println(m.Press())
	fmt.Println(m.Insert())
	fmt.Println(m.Press())
	fmt.Println(m.Insert())
	fmt.Println(m.Refill())
	fmt.Println(m.Insert())
	// Output:
	// 请先投币
	// 已收币
	// 已出货
	// 已售罄，退回硬币
	// 补货完成
	// 已收币
}

// --8<-- [end:example]
