package templatemethod_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/templatemethod"
)

var scores = []templatemethod.Record{{Name: "张三", Score: 95}, {Name: "李四", Score: 82}}

// HTML 实现两个必选步骤和钩子 Footer。
type HTML struct{}

func (HTML) Header() string { return "<table>" }
func (HTML) Row(r templatemethod.Record) string {
	return fmt.Sprintf("<tr><td>%s</td><td>%d</td></tr>", r.Name, r.Score)
}
func (HTML) Footer(n int) string { return "</table>" }

// 练习 1：新增 HTML 格式
func TestExercise1_HTML(t *testing.T) {
	want := "<table>\n" +
		"<tr><td>张三</td><td>95</td></tr>\n" +
		"<tr><td>李四</td><td>82</td></tr>\n" +
		"</table>\n"
	if got := templatemethod.Export(HTML{}, scores); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if got := templatemethod.Export(HTML{}, nil); got != "（无数据）\n" {
		t.Errorf("空数据：got %q", got)
	}
}

// Base 把流程挂在自己身上——这正是陷阱所在。
type Base struct{}

func (Base) Header() string { return "姓名 分数" }
func (Base) Row(r templatemethod.Record) string {
	return fmt.Sprintf("%s %d", r.Name, r.Score)
}
func (b Base) Run(rows []templatemethod.Record) string {
	return templatemethod.Export(b, rows) // b 永远是 Base
}

// Fancy 嵌入 Base，只重写 Header。
type Fancy struct{ Base }

func (Fancy) Header() string { return "== 成绩单 ==" }

// 练习 2：亲手踩一次嵌入陷阱
func TestExercise2_EmbedTrap(t *testing.T) {
	header := func(s string) string { return strings.SplitN(s, "\n", 2)[0] }

	if got := header(Fancy{}.Run(scores)); got != "姓名 分数" {
		t.Errorf("Fancy{}.Run：表头 got %q, want Base 的「姓名 分数」", got)
	}
	if got := header(templatemethod.Export(Fancy{}, scores)); got != "== 成绩单 ==" {
		t.Errorf("Export(Fancy{})：表头 got %q, want Fancy 的「== 成绩单 ==」", got)
	}
}
