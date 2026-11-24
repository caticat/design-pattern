package abstractfactory_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/creational/abstractfactory"
)

// HighContrast 是在包外新增的具体工厂，产出一族「高对比」组件。
type HighContrast struct{}

type hcButton struct{ label string }
type hcCheckbox struct{ label string }

func (b hcButton) Render() string   { return "[高对比按钮 " + b.label + "]" }
func (c hcCheckbox) Render() string { return "[高对比勾选框 " + c.label + "]" }

func (HighContrast) NewButton(label string) abstractfactory.Button {
	return hcButton{label}
}

func (HighContrast) NewCheckbox(label string) abstractfactory.Checkbox {
	return hcCheckbox{label}
}

// 练习 1：新增高对比主题
func TestExercise1_HighContrast(t *testing.T) {
	got := abstractfactory.LoginForm(HighContrast{})
	want := "[高对比勾选框 记住我] [高对比按钮 登录]"
	if got != want {
		t.Fatalf("得到 %q，期望 %q", got, want)
	}
}

// Mixed 嵌入 Light，按钮沿用白底，只把勾选框换成黑底。
type Mixed struct{ abstractfactory.Light }

func (Mixed) NewCheckbox(label string) abstractfactory.Checkbox {
	return abstractfactory.Dark{}.NewCheckbox(label)
}

// 练习 2：造一个「混搭」工厂
func TestExercise2_Mixed(t *testing.T) {
	got := abstractfactory.LoginForm(Mixed{})
	want := "[黑底勾选框 记住我] [白底按钮 登录]" // 能编译、能运行，接口不检查是否配套
	if got != want {
		t.Fatalf("得到 %q，期望 %q", got, want)
	}
}
