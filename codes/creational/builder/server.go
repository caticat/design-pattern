// Package builder 演示建造者模式：一步步设置参数，最后由 Build 统一校验、一次性创建对象。
package builder

import (
	"fmt"
	"time"
)

// --8<-- [start:product]
// Server 是要创建的产品。字段不导出，包外无法设置，只能通过建造者得到配好的 Server。
type Server struct {
	addr     string
	port     int
	timeout  time.Duration
	maxConns int
	tls      bool
}

// --8<-- [end:product]

func (s *Server) String() string {
	return fmt.Sprintf("%s:%d timeout=%s maxConns=%d tls=%t",
		s.addr, s.port, s.timeout, s.maxConns, s.tls)
}

// --8<-- [start:defaults]
// defaults 返回只填了必填项、其余都是默认值的 Server。
func defaults(addr string) Server {
	return Server{addr: addr, port: 80, timeout: 30 * time.Second, maxConns: 100}
}

// --8<-- [end:defaults]

// --8<-- [start:validate]
// validate 集中检查参数，以及参数之间的组合是否合理。
func (s *Server) validate() error {
	switch {
	case s.addr == "":
		return fmt.Errorf("地址不能为空")
	case s.port < 1 || s.port > 65535:
		return fmt.Errorf("端口 %d 超出范围", s.port)
	case s.maxConns < 1:
		return fmt.Errorf("最大连接数必须大于 0")
	case s.tls && s.port == 80:
		return fmt.Errorf("开启 TLS 时不能使用 80 端口")
	}
	return nil
}

// --8<-- [end:validate]
