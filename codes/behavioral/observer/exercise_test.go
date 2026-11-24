package observer_test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/observer"
)

// captureStdout 运行 f，返回它打印到标准输出的内容。
func captureStdout(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w
	done := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		done <- string(b)
	}()
	f()
	os.Stdout = old
	w.Close()
	return <-done
}

// ObserverFunc 让普通函数也能当观察者。
type ObserverFunc func(observer.Order)

func (f ObserverFunc) OnPaid(o observer.Order) { f(o) }

// 练习 1：让函数也能订阅 Shop
func TestExercise1_ObserverFunc(t *testing.T) {
	shop := &observer.Shop{}
	shop.Subscribe(ObserverFunc(func(o observer.Order) {
		if o.Amount >= 1000 {
			fmt.Println("风控：大额订单 " + o.ID)
		}
	}))

	got := captureStdout(t, func() { shop.Pay(observer.Order{ID: "A002", Amount: 1500}) })
	if want := "支付成功：A002\n风控：大额订单 A002\n"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// panicky 是一个出错的观察者。
type panicky struct{}

func (panicky) OnPaid(observer.Order) { panic("观察者出错") }

// Safe 把观察者的 panic 拦在自己这里，不连累后面的观察者。
type Safe struct{ ob observer.Observer }

func (s Safe) OnPaid(o observer.Order) {
	defer func() { recover() }()
	s.ob.OnPaid(o)
}

// 练习 2：一个观察者 panic，其余受连累
func TestExercise2_Safe(t *testing.T) {
	order := observer.Order{ID: "A001", Amount: 299}

	shop := &observer.Shop{}
	shop.Subscribe(observer.Stock{})
	shop.Subscribe(panicky{})
	shop.Subscribe(observer.Points{})
	got := captureStdout(t, func() {
		defer func() { recover() }() // 让测试能继续往下跑
		shop.Pay(order)
	})
	if want := "支付成功：A001\n库存：扣减 A001 的商品\n"; got != want {
		t.Errorf("未包装：got %q, want %q", got, want)
	}

	safeShop := &observer.Shop{}
	safeShop.Subscribe(observer.Stock{})
	safeShop.Subscribe(Safe{panicky{}})
	safeShop.Subscribe(observer.Points{})
	got = captureStdout(t, func() { safeShop.Pay(order) })
	if !strings.Contains(got, "积分：+29") {
		t.Errorf("包上 Safe 后：got %q, want 含「积分：+29」", got)
	}
}
