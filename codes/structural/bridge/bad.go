package bridge

import "fmt"

// --8<-- [start:bad]
// 反例：渠道是个字符串，每种消息里都 switch 一遍「怎么发」。
type badNormal struct{ channel string }

func (n badNormal) Notify(msg string) {
	switch n.channel {
	case "sms":
		fmt.Println("短信：" + msg)
	case "email":
		fmt.Println("邮件：" + msg)
	}
}

type badUrgent struct{ channel string }

func (u badUrgent) Notify(msg string) {
	text := urgentText(msg)
	switch u.channel { // 和上面一模一样的 switch，又抄了一遍
	case "sms":
		fmt.Println("短信：" + text)
	case "email":
		fmt.Println("邮件：" + text)
	}
}

// 加一个「微信」渠道：每种消息的 switch 都要加 case，漏改一处就发不出去。
// --8<-- [end:bad]

func urgentText(msg string) string { return "【紧急】" + msg + "，请立即处理" }
