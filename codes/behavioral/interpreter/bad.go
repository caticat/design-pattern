package interpreter

import "fmt"

// --8<-- [start:bad]
// 反例：用整条规则文本查找写死的判断逻辑。
func matchByText(rule string, p Parcel) (bool, error) {
	switch rule {
	case "fragile && cold":
		return p["fragile"] && p["cold"], nil
	case "fragile && express":
		return p["fragile"] && p["express"], nil
	default:
		return false, fmt.Errorf("未登记的规则：%q", rule)
	}
}

// --8<-- [end:bad]
