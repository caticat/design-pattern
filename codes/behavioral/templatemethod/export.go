package templatemethod

import (
	"fmt"
	"strings"
)

// --8<-- [start:steps]
// Record 是一行要导出的数据。
type Record struct {
	Name  string
	Score int
}

// Exporter 是导出流程中会变的步骤，每种格式各自实现。
type Exporter interface {
	Header() string
	Row(r Record) string
}

// Footer 是可选步骤（钩子）：实现了才会被调用。
type Footer interface {
	Footer(n int) string
}

// --8<-- [end:steps]

// --8<-- [start:template]
// Export 是模板方法：步骤和顺序固定在这里，每步怎么写交给 e。
func Export(e Exporter, rows []Record) string {
	if len(rows) == 0 {
		return "（无数据）\n"
	}
	var b strings.Builder
	b.WriteString(e.Header() + "\n")
	for _, r := range rows {
		b.WriteString(e.Row(r) + "\n")
	}
	if f, ok := e.(Footer); ok { // 钩子：可选
		b.WriteString(f.Footer(len(rows)) + "\n")
	}
	return b.String()
}

// --8<-- [end:template]

// --8<-- [start:concrete]
// CSV 只实现必选的两步。
type CSV struct{}

func (CSV) Header() string      { return "name,score" }
func (CSV) Row(r Record) string { return fmt.Sprintf("%s,%d", r.Name, r.Score) }

// Markdown 额外实现了钩子 Footer。
type Markdown struct{}

func (Markdown) Header() string      { return "| 姓名 | 分数 |\n|---|---|" }
func (Markdown) Row(r Record) string { return fmt.Sprintf("| %s | %d |", r.Name, r.Score) }
func (Markdown) Footer(n int) string { return fmt.Sprintf("共 %d 行", n) }

// --8<-- [end:concrete]
