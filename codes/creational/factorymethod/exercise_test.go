package factorymethod_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/caticat/design-pattern/codes/creational/factorymethod"
)

// captureOutput 运行 f，返回它打印到标准输出的内容。
func captureOutput(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = old }()

	f()
	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

// Inbox 是新接入的站内信通知方式。
type Inbox struct{}

func (Inbox) Send(to, msg string) { fmt.Printf("[站内信] 发给 %s：%s\n", to, msg) }

func init() {
	factorymethod.Register("inbox", func() factorymethod.Notifier { return Inbox{} })
}

// 练习 1：接入站内信
func TestExercise1_Inbox(t *testing.T) {
	out := captureOutput(t, func() {
		if err := factorymethod.SignUpWith("inbox", "小明", factorymethod.Create); err != nil {
			t.Error(err)
		}
	})
	if want := "[站内信] 发给 小明：欢迎注册\n"; out != want {
		t.Fatalf("输出 %q，期望 %q", out, want)
	}

	// 简单工厂 NewNotifier 的名单是写死的，不认识 inbox
	err := factorymethod.SignUp("inbox", "小明")
	if err == nil || err.Error() != `未知的通知方式 "inbox"` {
		t.Fatalf("得到 %v，期望「未知的通知方式」", err)
	}
}

// 练习 2：让工厂返回 nil 看看会怎样
func TestExercise2_NilFactory(t *testing.T) {
	badCreate := func(string) (factorymethod.Notifier, error) { return nil, nil }

	var p any
	func() {
		defer func() { p = recover() }()
		factorymethod.SignUpWith("inbox", "小明", badCreate) // 在内部的 n.Send 处 panic
	}()

	err, ok := p.(error)
	want := "runtime error: invalid memory address or nil pointer dereference"
	if !ok || err.Error() != want {
		t.Fatalf("recover 得到 %v，期望 %q", p, want)
	}
}
