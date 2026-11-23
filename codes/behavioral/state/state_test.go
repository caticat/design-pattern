package state_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/state"
)

func TestMachineCycle(t *testing.T) {
	m := state.NewMachine()
	steps := []struct {
		name string
		act  func() string
		want string
	}{
		{"未投币按键", m.Press, "请先投币"},
		{"未投币补货", m.Refill, "已有货，无需补货"},
		{"收币", m.Insert, "已收币"},
		{"重复投币", m.Insert, "已投币，退回硬币"},
		{"已投币补货", m.Refill, "已有货，无需补货"},
		{"出货", m.Press, "已出货"},
		{"售罄按键", m.Press, "已售罄"},
		{"售罄投币", m.Insert, "已售罄，退回硬币"},
		{"补货", m.Refill, "补货完成"},
		{"补货后仍需投币", m.Press, "请先投币"},
		{"补货后收币", m.Insert, "已收币"},
		{"再次出货", m.Press, "已出货"},
	}
	for _, step := range steps {
		if got := step.act(); got != step.want {
			t.Fatalf("%s：得到 %q，想要 %q", step.name, got, step.want)
		}
	}
}

func TestMachinesKeepSeparateStates(t *testing.T) {
	a, b := state.NewMachine(), state.NewMachine()
	a.Insert()
	if got := b.Press(); got != "请先投币" {
		t.Fatalf("另一台机器不应收到付款：%q", got)
	}
	if got := a.Press(); got != "已出货" {
		t.Fatalf("已投币的机器应出货：%q", got)
	}
}
