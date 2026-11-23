package adapter

import "fmt"

// --8<-- [start:adaptee]
// WalletClient 模拟第三方钱包 SDK（被适配者），我们不能修改它的代码。
// 它的习惯和商城不同：金额单位是分，结果用字符串状态码表示。
type WalletClient struct{}

// Transfer 转账，返回 "OK" 表示成功，其余为错误码。
func (*WalletClient) Transfer(fen int64, memo string) string {
	fmt.Printf("[钱包] 转账 %d 分，备注：%s\n", fen, memo)
	switch {
	case fen <= 0:
		return "INVALID_AMOUNT"
	case fen > 500000: // 单笔限额 5000 元
		return "LIMIT_EXCEEDED"
	}
	return "OK"
}

// --8<-- [end:adaptee]
