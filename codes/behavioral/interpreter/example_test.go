package interpreter_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/interpreter"
)

// --8<-- [start:example]
func Example() {
	rule, err := interpreter.Parse("fragile && (cold || express)")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 同一棵树，换两份包裹标签。
	fmt.Println(rule.Eval(interpreter.Parcel{"fragile": true, "cold": true}))
	fmt.Println(rule.Eval(interpreter.Parcel{"cold": true}))
	// Output:
	// true
	// false
}

// --8<-- [end:example]
