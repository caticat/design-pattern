package bridge

import (
	"fmt"
	"os"
)

// 反例输出相同，只是每种消息里都要 switch 一遍渠道。
func Example_bad() {
	for _, n := range []Notice{badNormal{"sms"}, badNormal{"email"}, badUrgent{"sms"}, badUrgent{"email"}} {
		n.Notify("磁盘快满了")
	}
	// Output:
	// 短信：磁盘快满了
	// 邮件：磁盘快满了
	// 短信：【紧急】磁盘快满了，请立即处理
	// 邮件：【紧急】磁盘快满了，请立即处理
}

// wechat 是新增的渠道：Normal、Urgent 一行不改就能用。
type wechat struct{}

func (wechat) Send(text string) { fmt.Println("微信：" + text) }

func Example_newSender() {
	NewNormal(wechat{}).Notify("今晚发版")
	NewUrgent(wechat{}).Notify("今晚发版")
	// Output:
	// 微信：今晚发版
	// 微信：【紧急】今晚发版，请立即处理
}

func Example_slog() {
	logBoth(os.Stdout)
	// Output:
	// level=INFO msg=用户登录 user=tom
	// {"level":"INFO","msg":"用户登录","user":"tom"}
}
