// Package state 用一次只装一罐饮料的售货机演示状态模式。
package state

// --8<-- [start:machine]
type vendingState interface {
	insert(*Machine) string
	press(*Machine) string
	refill(*Machine) string
}

// Machine 把操作交给当前状态。请用 NewMachine 创建。
type Machine struct {
	current vendingState
}

// NewMachine 创建一台有货、未投币的售货机。
func NewMachine() *Machine { return &Machine{current: ready{}} }

// Insert 投入一枚硬币，返回机器的提示。
func (m *Machine) Insert() string { return m.current.insert(m) }

// Press 按出货按钮，返回机器的提示。
func (m *Machine) Press() string { return m.current.press(m) }

// Refill 补入一罐饮料，返回机器的提示。
func (m *Machine) Refill() string { return m.current.refill(m) }

// --8<-- [end:machine]

// --8<-- [start:ready]
type ready struct{}

func (ready) insert(m *Machine) string {
	m.current = paid{}
	return "已收币"
}

func (ready) press(*Machine) string  { return "请先投币" }
func (ready) refill(*Machine) string { return "已有货，无需补货" }

// --8<-- [end:ready]

// --8<-- [start:other_states]
type paid struct{}

func (paid) insert(*Machine) string { return "已投币，退回硬币" }

func (paid) press(m *Machine) string {
	m.current = soldOut{}
	return "已出货"
}

func (paid) refill(*Machine) string { return "已有货，无需补货" }

type soldOut struct{}

func (soldOut) insert(*Machine) string { return "已售罄，退回硬币" }
func (soldOut) press(*Machine) string  { return "已售罄" }

func (soldOut) refill(m *Machine) string {
	m.current = ready{}
	return "补货完成"
}

// --8<-- [end:other_states]
