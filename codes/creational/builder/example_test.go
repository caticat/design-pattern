package builder_test

import (
	"fmt"
	"time"

	"github.com/caticat/design-pattern/codes/creational/builder"
)

// --8<-- [start:example]
func Example() {
	// 只写需要改的参数，每个值前面都有名字
	s, err := builder.NewServerBuilder("localhost").
		Port(8443).
		TLS().
		Timeout(5 * time.Second).
		Build()
	fmt.Println(s, err)

	// 组合不合理时，Build 拒绝创建
	_, err = builder.NewServerBuilder("localhost").TLS().Build()
	fmt.Println("出错:", err)
	// Output:
	// localhost:8443 timeout=5s maxConns=100 tls=true <nil>
	// 出错: 开启 TLS 时不能使用 80 端口
}

// --8<-- [end:example]
