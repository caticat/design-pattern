package interfaces_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/creational/factorymethod"
)

// --8<-- [start:example]
func Example() {
	var email factorymethod.EmailNotifier
	var n factorymethod.Notifier = &email
	_, ok := n.(*factorymethod.EmailNotifier)
	fmt.Println(ok)

	var empty *factorymethod.EmailNotifier
	n = empty
	fmt.Println(n == nil)
	// Output:
	// true
	// false
}

// --8<-- [end:example]
