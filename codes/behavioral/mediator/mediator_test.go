package mediator

import "testing"

func TestCheckoutRefresh(t *testing.T) {
	c := NewCheckout()
	steps := []struct {
		name    string
		change  func()
		visible bool
		enabled bool
	}{
		{"初始自取", func() {}, false, true},
		{"切换配送", func() { c.SetDelivery(true) }, true, false},
		{"填写地址", func() { c.SetAddress("松林路 8 号") }, true, true},
		{"仅留空白", func() { c.SetAddress("  \t") }, true, false},
		{"切回自取", func() { c.SetDelivery(false) }, false, true},
		{"自取时编辑地址", func() { c.SetAddress("松林路 8 号") }, false, true},
		{"再切配送", func() { c.SetDelivery(true) }, true, true},
	}
	for _, step := range steps {
		step.change()
		if visible, enabled := c.AddressVisible(), c.CanSubmit(); visible != step.visible || enabled != step.enabled {
			t.Errorf("%s：得到 (%t, %t)，期望 (%t, %t)", step.name, visible, enabled, step.visible, step.enabled)
		}
	}
}

func TestBadCheckoutKeepsStaleButton(t *testing.T) {
	a, b := &addressField{text: "松林路 8 号"}, &submitButton{}
	(&badDeliveryToggle{address: a, submit: b}).set(true)
	badEditAddress(a, "")
	if !b.enabled || a.text != "" {
		t.Fatal("反例应保留过期的按钮状态")
	}
}
