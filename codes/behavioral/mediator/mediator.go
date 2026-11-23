package mediator

import "strings"

// --8<-- [start:widgets]
type deliveryToggle struct {
	delivery bool
	changed  func()
}

func (t *deliveryToggle) set(delivery bool) {
	t.delivery = delivery
	t.changed()
}

type addressField struct {
	text    string
	visible bool
	changed func()
}

func (a *addressField) set(text string) {
	a.text = text
	a.changed()
}

type submitButton struct{ enabled bool }

// --8<-- [end:widgets]

// --8<-- [start:checkout]
// Checkout 协调配送开关、地址框和提交按钮；请用 NewCheckout 创建，创建后不要复制。
type Checkout struct {
	toggle  deliveryToggle
	address addressField
	submit  submitButton
}

// NewCheckout 创建默认为自取的结账表单。
func NewCheckout() *Checkout {
	c := &Checkout{}
	c.toggle.changed = c.refresh
	c.address.changed = c.refresh
	c.refresh()
	return c
}

func (c *Checkout) refresh() {
	delivery := c.toggle.delivery
	c.address.visible = delivery
	c.submit.enabled = !delivery || strings.TrimSpace(c.address.text) != ""
}

// --8<-- [end:checkout]

// --8<-- [start:controls]
// SetDelivery 切换配送（true）或自取（false）。
func (c *Checkout) SetDelivery(delivery bool) { c.toggle.set(delivery) }

// SetAddress 修改地址框的文本。
func (c *Checkout) SetAddress(text string) { c.address.set(text) }

// AddressVisible 报告地址框是否可见。
func (c *Checkout) AddressVisible() bool { return c.address.visible }

// CanSubmit 报告提交按钮是否可用。
func (c *Checkout) CanSubmit() bool { return c.submit.enabled }

// --8<-- [end:controls]
