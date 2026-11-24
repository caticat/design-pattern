package decorator_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caticat/design-pattern/codes/structural/decorator"
)

// Syrup 是新配料：加 2 元，描述追加「 + 糖浆」。
type Syrup struct{ B decorator.Beverage }

func (s Syrup) Cost() int           { return s.B.Cost() + 2 }
func (s Syrup) Description() string { return s.B.Description() + " + 糖浆" }

// 练习 1：加一种配料：糖浆
func TestExercise1_Syrup(t *testing.T) {
	var b decorator.Beverage = Syrup{B: decorator.Milk{B: decorator.Coffee{}}}

	got := fmt.Sprintf("%s = %d 元", b.Description(), b.Cost())
	if want := "美式 + 牛奶 + 糖浆 = 20 元"; got != want {
		t.Fatalf("得到 %q，期望 %q", got, want)
	}
}

// 练习 2：把中间件顺序反过来
// 日志写进 bytes.Buffer 而不是 os.Stdout，方便断言。
func TestExercise2_MiddlewareOrder(t *testing.T) {
	hello := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "hello")
	})

	// serve 用给定顺序处理一个不带 X-Token 的请求，返回状态码和日志
	serve := func(order func(log *bytes.Buffer) http.Handler) (int, string) {
		var log bytes.Buffer
		rec := httptest.NewRecorder()
		order(&log).ServeHTTP(rec, httptest.NewRequest("GET", "/hi", nil))
		return rec.Code, log.String()
	}

	// Logging 在外：被拒的请求也进了日志
	code, log := serve(func(log *bytes.Buffer) http.Handler {
		return decorator.Chain(hello, decorator.Logging(log), decorator.RequireToken)
	})
	if code != http.StatusForbidden || log != "-> GET /hi\n<- /hi\n" {
		t.Fatalf("Logging 在外：code=%d log=%q", code, log)
	}

	// RequireToken 在外：请求在进日志前就被拒了
	code, log = serve(func(log *bytes.Buffer) http.Handler {
		return decorator.Chain(hello, decorator.RequireToken, decorator.Logging(log))
	})
	if code != http.StatusForbidden || log != "" {
		t.Fatalf("RequireToken 在外：code=%d log=%q", code, log)
	}
}
