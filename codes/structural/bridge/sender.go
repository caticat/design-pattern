// Package bridge 演示桥接模式：两个独立变化的维度各自扩展，用字段接起来。
package bridge

import "fmt"

// --8<-- [start:sender]
// Sender 是「实现」这一边：负责怎么发。
type Sender interface {
	Send(text string)
}

// SMS 用短信发。
type SMS struct{}

func (SMS) Send(text string) { fmt.Println("短信：" + text) }

// Email 用邮件发。
type Email struct{}

func (Email) Send(text string) { fmt.Println("邮件：" + text) }

// --8<-- [end:sender]
