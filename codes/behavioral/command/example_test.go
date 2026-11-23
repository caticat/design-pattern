package command_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/command"
)

// --8<-- [start:example]
func Example() {
	lamp, fan := &command.Lamp{}, &command.Fan{}
	var q command.Queue
	q.Add(command.Func(fan.Start))
	q.Add(command.LightOn{Lamp: lamp})
	fmt.Println("执行前：", lamp.On, fan.Running)
	q.Run()
	fmt.Println("执行后：", lamp.On, fan.Running)
	// Output:
	// 执行前： false false
	// 执行后： true true
}

// --8<-- [end:example]
