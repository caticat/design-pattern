package interpreter_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/interpreter"
)

// 练习 2：易碎但不加急
func TestExercise2_FragileNotExpress(t *testing.T) {
	hand := interpreter.And(interpreter.Tag("fragile"), interpreter.Not(interpreter.Tag("express")))
	parsed, err := interpreter.Parse("fragile && !express")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		p    interpreter.Parcel
		want bool
	}{
		{interpreter.Parcel{"fragile": true}, true},
		{interpreter.Parcel{"fragile": true, "express": true}, false},
	} {
		if hand.Eval(c.p) != c.want || parsed.Eval(c.p) != c.want {
			t.Errorf("%v：want %v", c.p, c.want)
		}
	}
}
