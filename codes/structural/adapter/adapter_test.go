package adapter

import (
	"errors"
	"testing"
)

func TestWalletAdapter(t *testing.T) {
	tests := []struct {
		name string
		yuan int
		want string
	}{
		{"正常支付", 99, "支付成功"},
		{"金额无效", 0, "支付失败: 钱包错误码 INVALID_AMOUNT"},
		{"恰好限额", 5000, "支付成功"},
		{"超出限额", 5001, "支付失败: 钱包错误码 LIMIT_EXCEEDED"},
	}
	p := WalletAdapter{Client: &WalletClient{}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checkout(p, "A001", tt.yuan); got != tt.want {
				t.Fatalf("Checkout(%d) = %q, 期望 %q", tt.yuan, got, tt.want)
			}
		})
	}
}

func TestPayFunc(t *testing.T) {
	var gotID string
	var gotYuan int
	fake := PayFunc(func(orderID string, yuan int) error {
		gotID, gotYuan = orderID, yuan
		return errors.New("余额不足")
	})
	if got := Checkout(fake, "T001", 10); got != "支付失败: 余额不足" {
		t.Fatalf("Checkout = %q", got)
	}
	if gotID != "T001" || gotYuan != 10 {
		t.Fatalf("伪造支付收到 (%q, %d), 期望 (\"T001\", 10)", gotID, gotYuan)
	}
}

// 反例与新写法结果一致，说明重构没有改变行为。
func TestCheckoutBadMatches(t *testing.T) {
	c := &WalletClient{}
	for _, yuan := range []int{-1, 0, 1, 5000, 5001} {
		if a, b := CheckoutBad(c, "X", yuan), Checkout(WalletAdapter{c}, "X", yuan); a != b {
			t.Fatalf("yuan=%d: 反例 %q, 适配器 %q", yuan, a, b)
		}
	}
}
