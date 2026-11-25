package interpreter_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/interpreter"
)

// --8<-- [start:example]
func Example() {
	// 「易碎，且冷藏或加急」拼成一棵树。
	rule := interpreter.And(
		interpreter.Tag("fragile"),
		interpreter.Or(interpreter.Tag("cold"), interpreter.Tag("express")),
	)
	fmt.Println(rule.Eval(interpreter.Parcel{"fragile": true, "cold": true}))
	fmt.Println(rule.Eval(interpreter.Parcel{"cold": true}))

	// 规则来自配置时，Parse 把同样的文字变成同样的树。
	parsed, err := interpreter.Parse("fragile && (cold || express)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parsed.Eval(interpreter.Parcel{"fragile": true, "express": true}))
	// Output:
	// true
	// false
	// true
}

// --8<-- [end:example]
