package adapter_test

import (
	"errors"
	"testing"

	"github.com/caticat/design-pattern/codes/structural/adapter"
)

// CardClient 扮演第三方银行卡 SDK：金额以元计（float64），返回是否成功。
type CardClient struct{}

func (CardClient) Charge(amount float64, ref string) bool { return amount > 0 }

// CardAdapter 把 CardClient 适配成 adapter.Payer。
type CardAdapter struct{ Client CardClient }

var _ adapter.Payer = CardAdapter{}

func (a CardAdapter) Pay(orderID string, yuan int) error {
	if !a.Client.Charge(float64(yuan), orderID) { // int → float64
		return errors.New("银行卡扣款失败") // false → error
	}
	return nil
}

// 练习 1：再接一家银行卡支付
func TestExercise1_CardAdapter(t *testing.T) {
	if got := adapter.Checkout(CardAdapter{}, "B001", 50); got != "支付成功" {
		t.Fatalf("得到 %q，期望「支付成功」", got)
	}
}

// StrictWalletAdapter 把「限额」这条业务规则塞进了适配器（反面示范）。
type StrictWalletAdapter struct{ Client *adapter.WalletClient }

func (a StrictWalletAdapter) Pay(orderID string, yuan int) error {
	if yuan > 1000 {
		return errors.New("超过 1000 元")
	}
	return adapter.WalletAdapter{Client: a.Client}.Pay(orderID, yuan)
}

// 练习 2：观察业务规则放进适配器的后果
func TestExercise2_RuleInAdapter(t *testing.T) {
	client := &adapter.WalletClient{}

	got := adapter.Checkout(StrictWalletAdapter{Client: client}, "B002", 2000)
	if got != "支付失败: 超过 1000 元" {
		t.Fatalf("Strict 得到 %q", got)
	}

	// 换回普通适配器：同一笔订单通过了，限额随适配器一起丢了
	got = adapter.Checkout(adapter.WalletAdapter{Client: client}, "B002", 2000)
	if got != "支付成功" {
		t.Fatalf("WalletAdapter 得到 %q", got)
	}
}
