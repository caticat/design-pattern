package bridge

// --8<-- [start:notice]
// Notice 是「抽象」这一边：负责发什么。
type Notice interface {
	Notify(msg string)
}

// Normal 是普通消息：原样发出。
type Normal struct {
	sender Sender // 桥：怎么发，交给另一边
}

func NewNormal(s Sender) *Normal { return &Normal{sender: s} }

func (n *Normal) Notify(msg string) { n.sender.Send(msg) }

// --8<-- [end:notice]

// --8<-- [start:urgent]
// Urgent 是紧急消息：加前缀和处理要求。它不关心用什么渠道发。
type Urgent struct {
	sender Sender
}

func NewUrgent(s Sender) *Urgent { return &Urgent{sender: s} }

func (u *Urgent) Notify(msg string) {
	u.sender.Send("【紧急】" + msg + "，请立即处理")
}

// --8<-- [end:urgent]
