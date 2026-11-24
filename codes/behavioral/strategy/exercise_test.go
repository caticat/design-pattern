package strategy_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/strategy"
)

// Combo 依次套用多个优惠：它本身也是 Discount。
type Combo []strategy.Discount

func (c Combo) Apply(total int) int {
	for _, d := range c {
		total = d.Apply(total)
	}
	return total
}

// 练习 1：叠加优惠
func TestExercise1_Combo(t *testing.T) {
	d := Combo{strategy.FullReduction{Threshold: 200, Off: 30}, strategy.Percent{Rate: 80}}
	if got := strategy.NewCheckout(d).Pay(300); got != 216 { // 300 → 270 → 216
		t.Errorf("got %d, want 216", got)
	}
}

// firstOrder 是带状态的策略：首单立减 50，之后原价。
type firstOrder struct{ used bool }

func (f *firstOrder) Apply(total int) int {
	if f.used {
		return total
	}
	f.used = true
	return total - 50
}

// 练习 2：让策略带上状态
func TestExercise2_SharedState(t *testing.T) {
	shared := &firstOrder{}
	a := strategy.NewCheckout(shared)
	b := strategy.NewCheckout(shared)

	if got := a.Pay(300); got != 250 {
		t.Errorf("第一个收银台：got %d, want 250", got)
	}
	if got := b.Pay(300); got != 300 { // 首单已被 a 用掉
		t.Errorf("第二个收银台：got %d, want 300", got)
	}
}
