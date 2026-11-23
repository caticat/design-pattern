package builder

import "time"

// --8<-- [start:builder]
// ServerBuilder 是建造者：暂存参数，Build 时才真正创建 Server。
type ServerBuilder struct {
	s Server
}

// NewServerBuilder 只要求必填项，其余参数先用默认值。
func NewServerBuilder(addr string) *ServerBuilder {
	return &ServerBuilder{s: defaults(addr)}
}

// 每个设置方法都返回建造者自己，所以可以连着写。
func (b *ServerBuilder) Port(p int) *ServerBuilder {
	b.s.port = p
	return b
}

func (b *ServerBuilder) TLS() *ServerBuilder {
	b.s.tls = true
	return b
}

// --8<-- [end:builder]

// Timeout、MaxConns 与 Port 写法相同。
func (b *ServerBuilder) Timeout(d time.Duration) *ServerBuilder {
	b.s.timeout = d
	return b
}

func (b *ServerBuilder) MaxConns(n int) *ServerBuilder {
	b.s.maxConns = n
	return b
}

// --8<-- [start:build]
// Build 先整体校验，通过后返回一个新的 Server；不通过就不创建。
func (b *ServerBuilder) Build() (*Server, error) {
	s := b.s // 复制一份：之后再改建造者，不影响已创建的 Server
	if err := s.validate(); err != nil {
		return nil, err
	}
	return &s, nil
}

// --8<-- [end:build]
