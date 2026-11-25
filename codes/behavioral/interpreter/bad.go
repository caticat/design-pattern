package interpreter

// --8<-- [start:bad]
// 反例：规则直接写成 Go 代码。
func toTray(p Parcel) bool {
	return p["fragile"] && p["cold"] // 上周的规则：易碎且冷藏
}

// 这周运营要改成「易碎，且冷藏或加急」：改这一行、跑测试、发版。
// 下周再改，再来一遍。把规则写进配置也没用：程序读不懂一行文字。
// --8<-- [end:bad]
