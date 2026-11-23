package interpreter

import "testing"

func TestEval(t *testing.T) {
	cold := Parcel{"cold": true}
	express := Parcel{"express": true}
	fragileCold := Parcel{"fragile": true, "cold": true}
	fragileExpress := Parcel{"fragile": true, "express": true}

	tests := []struct {
		name   string
		rule   string
		parcel Parcel
		want   bool
	}{
		{"标签存在", "cold", cold, true},
		{"标签缺失", "fragile", cold, false},
		{"标签显式为假", "cold", Parcel{"cold": false}, false},
		{"否定缺失标签", "!cold", nil, true},
		{"否定整个子树", "!(cold || express)", cold, false},
		{"括号左支", "fragile && (cold || express)", fragileCold, true},
		{"括号右支", "fragile && (cold || express)", fragileExpress, true},
		{"括号改变结果", "fragile && (cold || express)", express, false},
		{"与先于或", "fragile && cold || express", express, true},
		{"非先于与", "!fragile && cold", nil, false},
		{"多余空格", "fragile&&cold ", fragileCold, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := Parse(tt.rule)
			if err != nil {
				t.Fatal(err)
			}
			if got := expr.Eval(tt.parcel); got != tt.want {
				t.Errorf("Eval(%v) = %v，期望 %v", tt.parcel, got, tt.want)
			}
		})
	}
}

func TestParseRejectsInvalidRule(t *testing.T) {
	for _, rule := range []string{
		"", "cold &&", "(cold", "cold express",
		"cold | express", "-cold", "cold()", "1", "parcel.cold", "cold || check()",
	} {
		t.Run(rule, func(t *testing.T) {
			if expr, err := Parse(rule); err == nil || expr != nil {
				t.Errorf("Parse(%q) = (%v, %v)，期望仅返回错误", rule, expr, err)
			}
		})
	}
}

func TestMatchByTextNeedsExactRule(t *testing.T) {
	parcel := Parcel{"fragile": true, "cold": true}
	if got, err := matchByText("fragile && cold", parcel); err != nil || !got {
		t.Fatalf("已登记规则结果为 (%v, %v)，期望 (true, nil)", got, err)
	}
	for _, rule := range []string{"fragile&&cold", "fragile && (cold || express)"} {
		if _, err := matchByText(rule, parcel); err == nil {
			t.Errorf("反例不应识别 %q", rule)
		}
	}
}
