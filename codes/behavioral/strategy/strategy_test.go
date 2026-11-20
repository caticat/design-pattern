package strategy

import "testing"

func TestDiscounts(t *testing.T) {
	tests := []struct {
		name  string
		d     Discount
		total int
		want  int
	}{
		{"不打折", NoDiscount{}, 300, 300},
		{"满减-达到门槛", FullReduction{Threshold: 200, Off: 30}, 300, 270},
		{"满减-未达门槛", FullReduction{Threshold: 200, Off: 30}, 150, 150},
		{"8折", Percent{Rate: 80}, 300, 240},
		{"函数策略", DiscountFunc(func(t int) int { return t - 5 }), 300, 295},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCheckout(tt.d).Pay(tt.total); got != tt.want {
				t.Fatalf("Pay(%d) = %d, 期望 %d", tt.total, got, tt.want)
			}
		})
	}
}

func TestSetDiscount(t *testing.T) {
	c := NewCheckout(NoDiscount{})
	c.SetDiscount(Percent{Rate: 50})
	if got := c.Pay(100); got != 50 {
		t.Fatalf("切换策略后 Pay(100) = %d, 期望 50", got)
	}
}

// 反例与新写法结果一致，说明重构没有改变行为。
func TestPayBadMatches(t *testing.T) {
	cases := map[string]Discount{
		"":               NoDiscount{},
		"full_reduction": FullReduction{Threshold: 200, Off: 30},
		"percent":        Percent{Rate: 80},
	}
	for kind, d := range cases {
		for _, total := range []int{100, 200, 300} {
			if PayBad(kind, total) != d.Apply(total) {
				t.Fatalf("kind=%q total=%d 结果不一致", kind, total)
			}
		}
	}
}
