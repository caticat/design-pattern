package factorymethod_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/creational/factorymethod"
)

func Example() {
	// 通知方式通常来自配置，业务代码只拿到 Notifier 接口
	for _, kind := range []string{"email", "sms", "pigeon"} {
		if err := factorymethod.SignUp(kind, "小明"); err != nil {
			fmt.Println("出错:", err)
		}
	}
	// Output:
	// [邮件 smtp.example.com] 发给 小明：欢迎注册
	// [短信]【Go 商城】发给 小明：欢迎注册
	// 出错: 未知的通知方式 "pigeon"
}

// --8<-- [start:example]
func ExampleSignUpWith() {
	// 装配时选择创建方式，注册流程不变。
	if err := factorymethod.SignUpWith("sms", "小明", factorymethod.NewNotifier); err != nil {
		fmt.Println(err)
	}
	if err := factorymethod.SignUpWith("wechat", "小明", factorymethod.Create); err != nil {
		fmt.Println(err)
	}
	// Output:
	// [短信]【Go 商城】发给 小明：欢迎注册
	// [微信] 发给 小明：欢迎注册
}

// --8<-- [end:example]

// --8<-- [start:example-register]
// WeChat 是新增的通知方式：只加代码，不改 Create。
type WeChat struct{}

func (WeChat) Send(to, msg string) { fmt.Printf("[微信] 发给 %s：%s\n", to, msg) }

func init() {
	factorymethod.Register("wechat", func() factorymethod.Notifier { return WeChat{} })
}

func ExampleCreate() {
	n, _ := factorymethod.Create("wechat")
	n.Send("小明", "订单已发货")
	// Output:
	// [微信] 发给 小明：订单已发货
}

// --8<-- [end:example-register]
