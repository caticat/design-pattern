package factorymethod

import "fmt"

// --8<-- [start:registry]
// Factory 是创建通知器的函数。
type Factory func() Notifier

var factories = map[string]Factory{}

// Register 登记一种通知方式。只在 init 中调用；重复登记会 panic。
func Register(kind string, f Factory) {
	if _, dup := factories[kind]; dup {
		panic("重复登记通知方式 " + kind)
	}
	factories[kind] = f
}

// Create 按名字找到工厂函数，用它创建通知器。
func Create(kind string) (Notifier, error) {
	f, ok := factories[kind]
	if !ok {
		return nil, fmt.Errorf("未知的通知方式 %q", kind)
	}
	return f(), nil
}

// --8<-- [end:registry]

// --8<-- [start:init]
// 每种通知方式在 init 中登记自己。
func init() {
	Register("email", func() Notifier { return &EmailNotifier{SMTPHost: "smtp.example.com"} })
	Register("sms", func() Notifier { return &SMSNotifier{Sign: "Go 商城"} })
}

// --8<-- [end:init]
