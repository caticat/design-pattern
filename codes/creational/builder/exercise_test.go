package builder_test

import (
	"testing"
	"time"

	"github.com/caticat/design-pattern/codes/creational/builder"
)

// WithHTTPS 把两个已有选项组合成一个：开启 TLS，端口设为 443。
func WithHTTPS() builder.Option {
	return func(s *builder.Server) {
		builder.WithTLS()(s)
		builder.WithPort(443)(s)
	}
}

// 练习 1：组合出一个新选项 WithHTTPS
func TestExercise1_WithHTTPS(t *testing.T) {
	s, err := builder.NewServer("localhost", WithHTTPS())
	if err != nil {
		t.Fatal(err)
	}
	if got, want := s.String(), "localhost:443 timeout=30s maxConns=100 tls=true"; got != want {
		t.Fatalf("得到 %q，期望 %q", got, want)
	}

	// 后面再把端口改回 80：TLS 仍开着，校验会拦下
	_, err = builder.NewServer("localhost", WithHTTPS(), builder.WithPort(80))
	if err == nil || err.Error() != "开启 TLS 时不能使用 80 端口" {
		t.Fatalf("期望被校验拦下，得到 %v", err)
	}
}

// 练习 2：把建造者当模板复用会怎样
func TestExercise2_SharedBuilder(t *testing.T) {
	base := builder.NewServerBuilder("localhost").Timeout(5 * time.Second)

	a, _ := base.Port(8080).Build()
	b, err := base.TLS().Build() // 改的仍是 base 指向的同一个建造者

	if err != nil {
		t.Fatalf("err = %v，期望 nil", err)
	}
	if got, want := a.String(), "localhost:8080 timeout=5s maxConns=100 tls=false"; got != want {
		t.Fatalf("a = %q，期望 %q", got, want)
	}
	// b 继承了给 a 设的 8080 端口
	if got, want := b.String(), "localhost:8080 timeout=5s maxConns=100 tls=true"; got != want {
		t.Fatalf("b = %q，期望 %q", got, want)
	}
}
