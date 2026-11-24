package command_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/command"
)

// Macro 把多条命令打包成一条：它本身也是 Command。
type Macro []command.Command

func (m Macro) Execute() {
	for _, c := range m {
		c.Execute()
	}
}

// 练习 1：一键「开工」宏命令
func TestExercise1_Macro(t *testing.T) {
	lamp := &command.Lamp{}
	fan := &command.Fan{}
	q := &command.Queue{}
	q.Add(Macro{command.LightOn{Lamp: lamp}, command.Func(fan.Start)})

	if lamp.On || fan.Running {
		t.Fatalf("Run 前：got %v %v, want false false", lamp.On, fan.Running)
	}
	q.Run()
	if !lamp.On || !fan.Running {
		t.Errorf("Run 后：got %v %v, want true true", lamp.On, fan.Running)
	}
}

// UndoableLightOn 执行前记下灯的原状态，撤销时恢复它。
type UndoableLightOn struct {
	Lamp *command.Lamp
	prev bool
}

// 指针接收者：Execute 里记下的 prev，Undo 才看得到。
func (c *UndoableLightOn) Execute() {
	c.prev = c.Lamp.On
	c.Lamp.TurnOn()
}

func (c *UndoableLightOn) Undo() { c.Lamp.On = c.prev }

// 练习 2：写一个正确的撤销
func TestExercise2_Undo(t *testing.T) {
	for _, before := range []bool{false, true} {
		lamp := &command.Lamp{On: before}
		c := &UndoableLightOn{Lamp: lamp}
		c.Execute()
		c.Undo()
		if lamp.On != before {
			t.Errorf("原本 %v：撤销后 got %v, want %v", before, lamp.On, before)
		}
	}
}
