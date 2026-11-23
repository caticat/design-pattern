package callbacks_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/command"
)

// --8<-- [start:example]
func Example() {
	var q command.Queue
	title := "初稿"
	saved := title
	q.Add(command.Func(func() { fmt.Println(title) }))
	q.Add(command.Func(func() { fmt.Println(saved) }))
	title = "终稿"
	q.Run()
	// Output:
	// 终稿
	// 初稿
}

// --8<-- [end:example]
