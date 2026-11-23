package adapter_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/structural/adapter"
)

// --8<-- [start:example]
func Example() {
	// 第三方 SDK 外面套一层适配器，商城就能当 Payer 用
	var p adapter.Payer = adapter.WalletAdapter{Client: &adapter.WalletClient{}}

	fmt.Println(adapter.Checkout(p, "A001", 99))
	fmt.Println(adapter.Checkout(p, "A002", 6000))
	// Output:
	// [钱包] 转账 9900 分，备注：订单 A001
	// 支付成功
	// [钱包] 转账 600000 分，备注：订单 A002
	// 支付失败: 钱包错误码 LIMIT_EXCEEDED
}

// --8<-- [end:example]

// --8<-- [start:example-func]
func ExamplePayFunc() {
	// 测试时用一个函数伪造支付，不必真的调用第三方
	fake := adapter.PayFunc(func(orderID string, yuan int) error {
		fmt.Println("[伪造] 收到", orderID, yuan, "元")
		return nil
	})
	fmt.Println(adapter.Checkout(fake, "T001", 10))
	// Output:
	// [伪造] 收到 T001 10 元
	// 支付成功
}

// --8<-- [end:example-func]
