package factorymethod

import (
	"errors"
	"testing"
)

func TestNewNotifier(t *testing.T) {
	if n, err := NewNotifier("email"); err != nil {
		t.Fatal(err)
	} else if _, ok := n.(*EmailNotifier); !ok {
		t.Errorf("email: 得到 %T", n)
	}
	if n, err := NewNotifier("sms"); err != nil {
		t.Fatal(err)
	} else if _, ok := n.(*SMSNotifier); !ok {
		t.Errorf("sms: 得到 %T", n)
	}
	if _, err := NewNotifier("pigeon"); err == nil {
		t.Error("未知类型应返回错误")
	}
}

func TestSignUpUnknown(t *testing.T) {
	if err := SignUp("pigeon", "小明"); err == nil {
		t.Error("未知类型应返回错误")
	}
}

type fake struct{ sent []string }

func (f *fake) Send(to, msg string) { f.sent = append(f.sent, to+":"+msg) }

func TestSignUpWith(t *testing.T) {
	n := &fake{}
	calls := 0
	err := SignUpWith("custom", "小明", func(kind string) (Notifier, error) {
		calls++
		if kind != "custom" {
			t.Errorf("创建参数 = %q，想要 custom", kind)
		}
		return n, nil
	})
	if err != nil || calls != 1 || len(n.sent) != 1 || n.sent[0] != "小明:欢迎注册" {
		t.Fatalf("错误 = %v，创建次数 = %d，通知 = %v", err, calls, n.sent)
	}
}

func TestSignUpWithFactoryError(t *testing.T) {
	n := &fake{}
	want := errors.New("创建失败")
	err := SignUpWith("custom", "小明", func(string) (Notifier, error) { return n, want })
	if err != want || len(n.sent) != 0 {
		t.Fatalf("错误 = %v，通知 = %v", err, n.sent)
	}
}

func TestSignUpWithNilFactory(t *testing.T) {
	if err := SignUpWith("custom", "小明", nil); err == nil {
		t.Fatal("空工厂应返回错误")
	}
}

func TestRegistry(t *testing.T) {
	f := &fake{}
	Register("test-fake", func() Notifier { return f })
	defer delete(factories, "test-fake")
	n, err := Create("test-fake")
	if err != nil {
		t.Fatal(err)
	}
	n.Send("a", "b")
	if len(f.sent) != 1 || f.sent[0] != "a:b" {
		t.Errorf("sent = %v", f.sent)
	}
	if _, err := Create("pigeon"); err == nil {
		t.Error("未登记的类型应返回错误")
	}
}

func TestRegisterDuplicatePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("重复登记应 panic")
		}
	}()
	Register("email", func() Notifier { return &EmailNotifier{} })
}

func TestCreateMatchesNewNotifier(t *testing.T) {
	for _, kind := range []string{"email", "sms"} {
		a, _ := NewNotifier(kind)
		b, _ := Create(kind)
		if a == nil || b == nil {
			t.Fatalf("%s: 创建失败", kind)
		}
		switch x := a.(type) {
		case *EmailNotifier:
			if *x != *b.(*EmailNotifier) {
				t.Errorf("%s 配置不一致", kind)
			}
		case *SMSNotifier:
			if *x != *b.(*SMSNotifier) {
				t.Errorf("%s 配置不一致", kind)
			}
		}
	}
}
