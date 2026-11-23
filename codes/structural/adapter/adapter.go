// Package adapter 演示适配器模式：把一个已有类型的接口，转换成调用方期望的接口。
package adapter

import "fmt"

// --8<-- [start:target]
// Payer 是商城期望的支付接口（目标接口）：金额单位是元，失败返回 error。
type Payer interface {
	Pay(orderID string, yuan int) error
}

// Checkout 是商城的结算逻辑，只认 Payer，不知道背后是哪家支付。
func Checkout(p Payer, orderID string, yuan int) string {
	if err := p.Pay(orderID, yuan); err != nil {
		return "支付失败: " + err.Error()
	}
	return "支付成功"
}

// --8<-- [end:target]

// --8<-- [start:adapter]
// WalletAdapter 把第三方的 WalletClient 适配成商城的 Payer。
type WalletAdapter struct {
	Client *WalletClient
}

// 编译期检查：WalletAdapter 必须实现 Payer，否则这一行编译不过。
var _ Payer = WalletAdapter{}

// Pay 只做「翻译」，不加业务逻辑。
func (a WalletAdapter) Pay(orderID string, yuan int) error {
	code := a.Client.Transfer(int64(yuan)*100, "订单 "+orderID) // 元 → 分，订单号 → 备注
	if code != "OK" {
		return fmt.Errorf("钱包错误码 %s", code) // 状态码 → error
	}
	return nil
}

// --8<-- [end:adapter]

// --8<-- [start:func]
// PayFunc 让普通函数也能当作 Payer 使用，常用来在测试中伪造支付。
type PayFunc func(orderID string, yuan int) error

func (f PayFunc) Pay(orderID string, yuan int) error { return f(orderID, yuan) }

// --8<-- [end:func]
