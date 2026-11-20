package strategy

// --8<-- [start:bad]
// PayBad 是反例：所有优惠规则挤在一个 switch 里。
// 每加一种优惠都要修改这个函数，规则越多越难读、越难测。
func PayBad(kind string, total int) int {
	switch kind {
	case "full_reduction":
		if total >= 200 {
			return total - 30
		}
		return total
	case "percent":
		return total * 80 / 100
	default:
		return total
	}
}

// --8<-- [end:bad]
