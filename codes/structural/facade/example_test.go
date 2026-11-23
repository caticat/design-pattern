package facade_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/structural/facade"
)

// --8<-- [start:example]
func Example() {
	ht := facade.NewHomeTheater()

	ht.Watch("星际旅行") // 一个调用，背后 6 步
	fmt.Println("……两小时后……")
	ht.End()
	// Output:
	// 灯光：调到 10%
	// 投影仪：开机
	// 投影仪：信号源切到 播放器
	// 音响：开机
	// 音响：音量 30
	// 播放器：播放 星际旅行
	// ……两小时后……
	// 播放器：停止
	// 音响：关机
	// 投影仪：关机
	// 灯光：调到 100%
}

// --8<-- [end:example]
