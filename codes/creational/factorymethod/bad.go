package factorymethod

// --8<-- [start:bad]
// 反例：每个业务函数都自己判断类型、自己创建通知器。
func SignUpBad(kind, user string) {
	switch kind {
	case "email":
		n := &EmailNotifier{SMTPHost: "smtp.example.com"}
		n.Send(user, "欢迎注册")
	case "sms":
		n := &SMSNotifier{Sign: "Go 商城"}
		n.Send(user, "欢迎注册")
	}
}

func ResetPasswordBad(kind, user string) {
	switch kind { // 同样的判断和创建细节，又写了一遍
	case "email":
		n := &EmailNotifier{SMTPHost: "smtp.example.com"}
		n.Send(user, "验证码 1234")
	case "sms":
		n := &SMSNotifier{Sign: "Go 商城"}
		n.Send(user, "验证码 1234")
	}
}

// --8<-- [end:bad]
