package builder

import "time"

// --8<-- [start:bad]
// 反例：所有参数都挤在构造函数里。
func NewServerBad(addr string, port int, timeout time.Duration, maxConns int, tls bool) *Server {
	return &Server{addr, port, timeout, maxConns, tls}
}

// 调用处：只想改端口，其余参数也得全写；
// 把 port 和 maxConns 写反了，编译器也不会报错。
var badServer = NewServerBad("localhost", 100, 30*time.Second, 8080, false)

// --8<-- [end:bad]
