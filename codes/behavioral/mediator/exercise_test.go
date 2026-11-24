package mediator_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/mediator"
)

// 练习 2：复制协调器会发生什么
func TestExercise2_Copy(t *testing.T) {
	a := mediator.NewCheckout()
	a.SetDelivery(true)
	b := *a // 复制后，b 里控件的回调仍指向 a.refresh
	b.SetAddress("松林路 8 号")

	if got := b.AddressVisible(); got != true {
		t.Errorf("AddressVisible：got %v, want true", got)
	}
	if got := b.CanSubmit(); got != false {
		t.Errorf("CanSubmit：got %v, want false（刷新的是 a 的按钮）", got)
	}
}
