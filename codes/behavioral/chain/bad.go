package chain

import "fmt"

// --8<-- [start:bad]
// 反例：谁能批、按什么顺序问，全写死在一个函数里。
func approve(r Request) string {
	if r.Amount <= 1000 {
		return fmt.Sprintf("组长批准：%s %d 元", r.Name, r.Amount)
	} else if r.Amount <= 5000 {
		return fmt.Sprintf("经理批准：%s %d 元", r.Name, r.Amount)
	} else if r.Amount <= 50000 {
		return fmt.Sprintf("总监批准：%s %d 元", r.Name, r.Amount)
	}
	return fmt.Sprintf("无人能批：%s %d 元", r.Name, r.Amount)
}

// --8<-- [end:bad]
