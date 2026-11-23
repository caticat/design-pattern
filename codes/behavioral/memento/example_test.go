package memento_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/memento"
)

// --8<-- [start:example]
func Example() {
	e := &memento.Editor{}
	e.Rename("周末计划")
	e.Append("去公园")
	h := memento.NewHistory(e)

	h.Checkpoint()
	e.Rename("周日计划")
	e.ReplaceLine(0, "去书店")
	fmt.Println(e.Title() + "：" + e.Text())

	h.Undo()
	fmt.Println(e.Title() + "：" + e.Text())
	fmt.Println(h.Undo())
	// Output:
	// 周日计划：去书店
	// 周末计划：去公园
	// false
}

// --8<-- [end:example]
