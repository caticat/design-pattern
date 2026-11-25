package decorator

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBeverage(t *testing.T) {
	tests := []struct {
		b        Beverage
		wantCost int
		wantDesc string
	}{
		{Coffee{}, 15, "美式"},
		{Milk{Coffee{}}, 18, "美式 + 牛奶"},
		{Cream{Milk{Coffee{}}}, 23, "美式 + 牛奶 + 奶油"},
		{Milk{Milk{Coffee{}}}, 21, "美式 + 牛奶 + 牛奶"},
	}
	for _, tt := range tests {
		if got := tt.b.Cost(); got != tt.wantCost {
			t.Errorf("%s: Cost() = %d, 期望 %d", tt.wantDesc, got, tt.wantCost)
		}
		if got := tt.b.Description(); got != tt.wantDesc {
			t.Errorf("Description() = %q, 期望 %q", got, tt.wantDesc)
		}
	}
}

// 反例与装饰器结果一致。
func TestBadMatches(t *testing.T) {
	bad, good := Order{Milk: true, Cream: true}, Cream{Milk{Coffee{}}}
	if bad.Cost() != good.Cost() || bad.Description() != good.Description() {
		t.Fatal("反例与装饰器结果不一致")
	}
}

func hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello"))
	})
}

func TestChain(t *testing.T) {
	var log bytes.Buffer
	h := Chain(hello(), Logging(&log), RequireToken)

	// 无 token：被 RequireToken 拦截，但外层 Logging 仍会记录。
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/hi", nil))
	if rec.Code != http.StatusForbidden {
		t.Fatalf("无 token 状态码 = %d, 期望 403", rec.Code)
	}
	if log.String() != "-> GET /hi\n<- /hi\n" {
		t.Fatalf("日志 = %q", log.String())
	}

	// 有 token：放行。
	rec = httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/hi", nil)
	req.Header.Set("X-Token", "t")
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK || rec.Body.String() != "hello" {
		t.Fatalf("有 token: code=%d body=%q", rec.Code, rec.Body.String())
	}
}
