package bridge_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/structural/bridge"
)

// --8<-- [start:example]
func Example() {
	sms, email := bridge.SMS{}, bridge.Email{}

	notices := []bridge.Notice{
		bridge.NewNormal(sms),
		bridge.NewUrgent(sms),
		bridge.NewUrgent(email), // 同一种消息，换个渠道
	}
	for i, n := range notices {
		fmt.Printf("%d. ", i+1)
		n.Notify("服务器 CPU 超过 90%")
	}
	// Output:
	// 1. 短信：服务器 CPU 超过 90%
	// 2. 短信：【紧急】服务器 CPU 超过 90%，请立即处理
	// 3. 邮件：【紧急】服务器 CPU 超过 90%，请立即处理
}

// --8<-- [end:example]
