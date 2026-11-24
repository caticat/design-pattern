package chain_test

import (
	"fmt"
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/chain"
)

// Blocker 是练习 1 新加的一环：只拦私人消费，其余交给 Next。
type Blocker struct {
	Next chain.Handler
}

func (b Blocker) Handle(r chain.Request) string {
	if r.Name == "私人消费" {
		return fmt.Sprintf("驳回：%s %d 元", r.Name, r.Amount)
	}
	return b.Next.Handle(r)
}

// 练习 1：在链头加一道拦截
func TestExercise1_Blocker(t *testing.T) {
	head := Blocker{Next: chain.NewChain()}

	if got, want := head.Handle(chain.Request{Name: "私人消费", Amount: 500}), "驳回：私人消费 500 元"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if got, want := head.Handle(chain.Request{Name: "差旅", Amount: 3000}), "经理批准：差旅 3000 元"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// 练习 2：把链的顺序串反
func TestExercise2_Reversed(t *testing.T) {
	// 从链尾往前建：组长是链尾，总监是链头。
	lead := &chain.Approver{Title: "组长", Limit: 1000}
	manager := &chain.Approver{Title: "经理", Limit: 5000, Next: lead}
	head := &chain.Approver{Title: "总监", Limit: 50000, Next: manager}

	cases := []struct {
		amount int
		want   string
	}{
		{800, "总监批准：差旅 800 元"},
		{3000, "总监批准：差旅 3000 元"},
		{80000, "无人能批：差旅 80000 元"},
	}
	for _, c := range cases {
		if got := head.Handle(chain.Request{Name: "差旅", Amount: c.amount}); got != c.want {
			t.Errorf("%d 元：got %q, want %q", c.amount, got, c.want)
		}
	}
}
