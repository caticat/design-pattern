package builder

import "time"

// --8<-- [start:options]
// Option 修改 Server 的某一项配置。
type Option func(*Server)

func WithPort(p int) Option              { return func(s *Server) { s.port = p } }
func WithTimeout(d time.Duration) Option { return func(s *Server) { s.timeout = d } }
func WithMaxConns(n int) Option          { return func(s *Server) { s.maxConns = n } }
func WithTLS() Option                    { return func(s *Server) { s.tls = true } }

// NewServer 先填默认值，再依次应用选项，最后校验。
func NewServer(addr string, opts ...Option) (*Server, error) {
	s := defaults(addr)
	for _, opt := range opts {
		opt(&s)
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return &s, nil
}

// --8<-- [end:options]
