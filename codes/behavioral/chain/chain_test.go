package chain

import (
	"strings"
	"testing"
)

func TestChainRoutes(t *testing.T) {
	c := NewChain()
	cases := map[int]string{
		1000:  "组长",
		1001:  "经理",
		5000:  "经理",
		50000: "总监",
		50001: "无人能批",
	}
	for amount, want := range cases {
		got := c.Handle(Request{Name: "测试", Amount: amount})
		if !strings.HasPrefix(got, want) {
			t.Errorf("%d 元：得到 %q，想要以 %q 开头", amount, got, want)
		}
	}
}

func TestSingleApprover(t *testing.T) {
	a := &Approver{Title: "组长", Limit: 100}
	if got := a.Handle(Request{Name: "x", Amount: 200}); !strings.HasPrefix(got, "无人能批") {
		t.Errorf("链尾批不了应返回无人能批，得到 %q", got)
	}
}

func TestFirst(t *testing.T) {
	small := func(r Request) (string, bool) { return "小额", r.Amount <= 100 }
	large := func(r Request) (string, bool) { return "大额", r.Amount <= 1000 }
	handle := First(small, large)
	for amount, want := range map[int]string{50: "小额", 500: "大额", 5000: "无人处理"} {
		if got := handle(Request{Amount: amount}); got != want {
			t.Errorf("%d：得到 %q，想要 %q", amount, got, want)
		}
	}
}

func TestBadMatchesChain(t *testing.T) {
	c := NewChain()
	for _, amount := range []int{1, 1000, 1001, 5000, 5001, 50000, 50001} {
		r := Request{Name: "对照", Amount: amount}
		if got, want := approve(r), c.Handle(r); got != want {
			t.Errorf("%d：反例 %q，责任链 %q", amount, got, want)
		}
	}
}
