package mediator

import "strings"

// --8<-- [start:bad]
// 反例：配送开关直接修改其他控件，地址框却没有同步更新按钮。
type badDeliveryToggle struct {
	address *addressField
	submit  *submitButton
}

func (t *badDeliveryToggle) set(delivery bool) {
	t.address.visible = delivery
	t.submit.enabled = !delivery || strings.TrimSpace(t.address.text) != ""
}

// 反例：地址框只保存自己的文本，不会重新计算按钮状态。
func badEditAddress(a *addressField, text string) {
	a.text = text
}

// --8<-- [end:bad]
