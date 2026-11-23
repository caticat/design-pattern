package bridge

import "fmt"

// --8<-- [start:bad]
// 反例：每种「消息 × 渠道」的组合写一个类型。
type normalSMS struct{}
type normalEmail struct{}
type urgentSMS struct{}
type urgentEmail struct{}

func (normalSMS) Notify(msg string)   { fmt.Println("短信：" + msg) }
func (normalEmail) Notify(msg string) { fmt.Println("邮件：" + msg) }
func (urgentSMS) Notify(msg string)   { fmt.Println("短信：" + urgentText(msg)) }
func (urgentEmail) Notify(msg string) { fmt.Println("邮件：" + urgentText(msg)) }

// 加一个「微信」渠道：再写 normalWechat、urgentWechat；
// 加一种「定时」消息：每个渠道再各写一个。M 种消息 × N 种渠道 = M×N 个类型。
// --8<-- [end:bad]

func urgentText(msg string) string { return "【紧急】" + msg + "，请立即处理" }
