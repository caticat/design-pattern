package state

// --8<-- [start:bad]
// 反例：用两个标志位记录状态，各入口自己检查放行条件。
type flagMachine struct {
	stocked bool
	paid    bool
}

func (m *flagMachine) insert() string {
	if !m.stocked {
		return "已售罄，退回硬币"
	}
	if m.paid {
		return "已投币，退回硬币"
	}
	m.paid = true
	return "已收币"
}

func (m *flagMachine) press() string {
	if !m.stocked {
		return "已售罄"
	}
	// 漏查 paid：有货就出货，没投币也能拿到饮料。
	m.stocked, m.paid = false, false
	return "已出货"
}

// --8<-- [end:bad]
