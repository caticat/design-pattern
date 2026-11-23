package abstractfactory

import (
	"strings"
	"testing"
)

func TestLoginFormMatches(t *testing.T) {
	cases := []struct {
		theme      Theme
		want, deny string
	}{
		{Light{}, "白底", "黑底"},
		{Dark{}, "黑底", "白底"},
	}
	for _, c := range cases {
		got := LoginForm(c.theme)
		if strings.Count(got, c.want) != 2 || strings.Contains(got, c.deny) {
			t.Errorf("LoginForm(%T) = %q，控件没有配套", c.theme, got)
		}
	}
}

// 反例确实会混搭：这正是抽象工厂要消除的问题。
func TestLoginFormBadMixes(t *testing.T) {
	got := LoginFormBad(true)
	if !strings.Contains(got, "黑底") || !strings.Contains(got, "白底") {
		t.Errorf("LoginFormBad(true) = %q，预期出现混搭", got)
	}
}

// fakeTheme 演示测试时整族替换成假对象。
type fakeTheme struct{}
type fakeWidget string

func (w fakeWidget) Render() string             { return string(w) }
func (fakeTheme) NewButton(l string) Button     { return fakeWidget("B:" + l) }
func (fakeTheme) NewCheckbox(l string) Checkbox { return fakeWidget("C:" + l) }

func TestLoginFormWithFake(t *testing.T) {
	if got, want := LoginForm(fakeTheme{}), "C:记住我 B:登录"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
