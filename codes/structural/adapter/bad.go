package adapter

// --8<-- [start:bad]
// CheckoutBad 是反例：结算逻辑直接调用第三方 SDK。
// 元转分、错误码判断都混在业务代码里；
// 换一家支付或 SDK 升级时，每个这样的调用处都要改。
func CheckoutBad(c *WalletClient, orderID string, yuan int) string {
	code := c.Transfer(int64(yuan)*100, "订单 "+orderID)
	if code != "OK" {
		return "支付失败: 钱包错误码 " + code
	}
	return "支付成功"
}

// --8<-- [end:bad]
