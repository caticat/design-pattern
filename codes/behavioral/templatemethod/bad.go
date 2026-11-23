package templatemethod

import (
	"fmt"
	"strings"
)

// --8<-- [start:bad]
// 反例：每种格式各写一个完整的导出函数，流程整段复制。
func exportCSV(rows []Record) string {
	if len(rows) == 0 {
		return "（无数据）\n"
	}
	var b strings.Builder
	b.WriteString("name,score\n")
	for _, r := range rows {
		fmt.Fprintf(&b, "%s,%d\n", r.Name, r.Score)
	}
	return b.String()
}

func exportMarkdown(rows []Record) string {
	if len(rows) == 0 { // 同样的检查，再写一遍
		return "（无数据）\n"
	}
	var b strings.Builder
	b.WriteString("| 姓名 | 分数 |\n|---|---|\n")
	for _, r := range rows { // 同样的循环，再写一遍
		fmt.Fprintf(&b, "| %s | %d |\n", r.Name, r.Score)
	}
	fmt.Fprintf(&b, "共 %d 行\n", len(rows))
	return b.String()
}

// --8<-- [end:bad]
