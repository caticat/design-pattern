package observer

import (
	"fmt"
	"testing"
)

// 反例输出与正例前四行相同，但所有下游都写死在 payBad 里。
func Example_bad() {
	payBad(Order{ID: "A001", Amount: 299})
	// Output:
	// 支付成功：A001
	// 库存：扣减 A001 的商品
	// 短信：A001 已支付 299 元
	// 积分：+29
}

type recorder struct{ got []string }

func (r *recorder) OnPaid(o Order) { r.got = append(r.got, o.ID) }

func TestNotifyAll(t *testing.T) {
	var shop Shop
	a, b := &recorder{}, &recorder{}
	shop.Subscribe(a)
	shop.Subscribe(b)
	shop.Pay(Order{ID: "X"})
	if len(a.got) != 1 || len(b.got) != 1 {
		t.Fatalf("每个观察者应收到 1 次，got %v %v", a.got, b.got)
	}
}

func Example_bus() {
	var bus Bus
	bus.On(func(o Order) { fmt.Println("日志：" + o.ID) })
	cancel := bus.On(func(o Order) { fmt.Println("短信：" + o.ID) })

	bus.Emit(Order{ID: "A001"})
	cancel() // 取消短信订阅
	bus.Emit(Order{ID: "A002"})
	// Output:
	// 日志：A001
	// 短信：A001
	// 日志：A002
}
