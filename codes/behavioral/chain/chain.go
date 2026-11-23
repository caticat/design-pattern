// Package chain 演示责任链模式：报销单沿审批链传递，直到有人能批。
package chain

import "fmt"

// --8<-- [start:handler]
// Request 是一张报销单。
type Request struct {
	Name   string
	Amount int
}

// Handler 是链上每一环的统一接口。
type Handler interface {
	Handle(r Request) string
}

// --8<-- [end:handler]

// --8<-- [start:approver]
// Approver 能批 Limit 元以内的单子，批不了就交给 Next。
type Approver struct {
	Title string
	Limit int
	Next  Handler
}

func (a *Approver) Handle(r Request) string {
	if r.Amount <= a.Limit {
		return fmt.Sprintf("%s批准：%s %d 元", a.Title, r.Name, r.Amount)
	}
	if a.Next == nil {
		return fmt.Sprintf("无人能批：%s %d 元", r.Name, r.Amount)
	}
	return a.Next.Handle(r)
}

// --8<-- [end:approver]

// --8<-- [start:chain]
// NewChain 把三个审批人串起来：组长 → 经理 → 总监。
func NewChain() Handler {
	director := &Approver{Title: "总监", Limit: 50000}
	manager := &Approver{Title: "经理", Limit: 5000, Next: director}
	return &Approver{Title: "组长", Limit: 1000, Next: manager}
}

// --8<-- [end:chain]
