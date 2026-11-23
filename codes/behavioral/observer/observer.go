// Package observer 演示观察者模式：订单支付成功后，通知所有订阅者。
package observer

import "fmt"

// --8<-- [start:observer]
// Order 是要通知的事件内容。
type Order struct {
	ID     string
	Amount int // 金额（元）
}

// Observer 是观察者：谁想知道「支付成功」，就实现它。
type Observer interface {
	OnPaid(o Order)
}

// --8<-- [end:observer]

// --8<-- [start:subject]
// Shop 是被观察的主题：只维护一份观察者列表。
type Shop struct {
	observers []Observer
}

// Subscribe 登记一个观察者。
func (s *Shop) Subscribe(ob Observer) {
	s.observers = append(s.observers, ob)
}

// Pay 完成支付后，挨个通知，不关心对方是谁。
func (s *Shop) Pay(o Order) {
	fmt.Println("支付成功：" + o.ID)
	for _, ob := range s.observers {
		ob.OnPaid(o)
	}
}

// --8<-- [end:subject]

// --8<-- [start:concrete]
// Stock 扣减库存。
type Stock struct{}

func (Stock) OnPaid(o Order) { fmt.Println("库存：扣减 " + o.ID + " 的商品") }

// SMS 给用户发短信。
type SMS struct{}

func (SMS) OnPaid(o Order) { fmt.Printf("短信：%s 已支付 %d 元\n", o.ID, o.Amount) }

// Points 按金额加积分。
type Points struct{}

func (Points) OnPaid(o Order) { fmt.Printf("积分：+%d\n", o.Amount/10) }

// --8<-- [end:concrete]
