package builder

import (
	"testing"
	"time"
)

func TestBuilderDefaults(t *testing.T) {
	s, err := NewServerBuilder("localhost").Build()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := s.String(), "localhost:80 timeout=30s maxConns=100 tls=false"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBuilderValidate(t *testing.T) {
	cases := map[string]*ServerBuilder{
		"空地址":      NewServerBuilder(""),
		"端口越界":     NewServerBuilder("h").Port(70000),
		"连接数为 0":   NewServerBuilder("h").MaxConns(0),
		"TLS 用 80": NewServerBuilder("h").TLS(),
	}
	for name, b := range cases {
		if s, err := b.Build(); err == nil {
			t.Errorf("%s：应当失败，却创建了 %v", name, s)
		}
	}
}

// Build 返回副本：之后继续修改建造者，已创建的 Server 不受影响。
func TestBuildReturnsCopy(t *testing.T) {
	b := NewServerBuilder("h").Port(8080)
	s1, _ := b.Build()
	s2, _ := b.Port(9090).Build()
	if s1.port != 8080 || s2.port != 9090 {
		t.Errorf("s1.port=%d s2.port=%d", s1.port, s2.port)
	}
}

// 函数式选项与建造者共用默认值和校验，结果一致。
func TestOptionsMatchBuilder(t *testing.T) {
	a, err1 := NewServerBuilder("h").Port(8443).TLS().Timeout(time.Second).MaxConns(5).Build()
	b, err2 := NewServer("h", WithPort(8443), WithTLS(), WithTimeout(time.Second), WithMaxConns(5))
	if err1 != nil || err2 != nil || a.String() != b.String() {
		t.Errorf("builder=%v,%v options=%v,%v", a, err1, b, err2)
	}
	if _, err := NewServer("h", WithTLS()); err == nil {
		t.Error("NewServer 应当拒绝 TLS + 80 端口")
	}
}

// 反例的参数写反了，编译器发现不了，只能靠运行结果。
func TestBadSwapsArgs(t *testing.T) {
	if badServer.port != 100 || badServer.maxConns != 8080 {
		t.Errorf("badServer = %v", badServer)
	}
}
