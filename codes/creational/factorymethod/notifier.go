// Package factorymethod 演示工厂方法模式：把「创建哪个类型」集中到一个函数里，
// 调用方只说要什么，拿到的是接口。
package factorymethod

import "fmt"

// --8<-- [start:product]
// Notifier 是所有通知方式的统一接口（产品）。
type Notifier interface {
	Send(to, msg string)
}

// EmailNotifier 通过邮件发送通知。
type EmailNotifier struct {
	SMTPHost string // 邮件服务器地址
}

func (e *EmailNotifier) Send(to, msg string) {
	fmt.Printf("[邮件 %s] 发给 %s：%s\n", e.SMTPHost, to, msg)
}

// SMSNotifier 通过短信发送通知。
type SMSNotifier struct {
	Sign string // 短信签名
}

func (s *SMSNotifier) Send(to, msg string) {
	fmt.Printf("[短信]【%s】发给 %s：%s\n", s.Sign, to, msg)
}

// --8<-- [end:product]

// --8<-- [start:factory]
// NewNotifier 是工厂方法：按 kind 决定创建哪种通知器，返回接口。
// 创建细节（服务器地址、签名）只写在这一处。
func NewNotifier(kind string) (Notifier, error) {
	switch kind {
	case "email":
		return &EmailNotifier{SMTPHost: "smtp.example.com"}, nil
	case "sms":
		return &SMSNotifier{Sign: "Go 商城"}, nil
	default:
		return nil, fmt.Errorf("未知的通知方式 %q", kind)
	}
}

// --8<-- [end:factory]

// --8<-- [start:client]
// SignUp 只认 Notifier 接口，不知道也不关心具体类型。
func SignUp(kind, user string) error {
	n, err := NewNotifier(kind)
	if err != nil {
		return err
	}
	n.Send(user, "欢迎注册")
	return nil
}

// --8<-- [end:client]
