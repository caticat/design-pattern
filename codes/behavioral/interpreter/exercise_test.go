package interpreter_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/interpreter"
)

// 练习 2：什么样的标签名能用
func TestExercise2_TagNames(t *testing.T) {
	// 中文是合法的 Go 标识符。
	e, err := interpreter.Parse("易碎 && 冷藏")
	if err != nil {
		t.Fatal(err)
	}
	if !e.Eval(interpreter.Parcel{"易碎": true, "冷藏": true}) {
		t.Error("易碎 && 冷藏：want true")
	}

	// 连字符被读成减法。
	_, err = interpreter.Parse("same-day && cold")
	if err == nil || err.Error() != "不支持的运算符：-" {
		t.Errorf("same-day：got %v, want 不支持的运算符：-", err)
	}

	// range 是 Go 关键字，连语法都过不了。
	if _, err = interpreter.Parse("range || cold"); err == nil {
		t.Error("range：want 语法错误")
	}
}
