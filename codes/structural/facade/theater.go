// Package facade 演示外观模式：给复杂的子系统一个简单入口。
package facade

import "fmt"

// --8<-- [start:subsystem]
// Lights 是灯光。
type Lights struct{}

func (Lights) Dim(percent int) { fmt.Printf("灯光：调到 %d%%\n", percent) }

// Projector 是投影仪：要先开机，再选信号源。
type Projector struct{}

func (Projector) On()                 { fmt.Println("投影仪：开机") }
func (Projector) Input(source string) { fmt.Println("投影仪：信号源切到", source) }
func (Projector) Off()                { fmt.Println("投影仪：关机") }

// Speaker 是音响。
type Speaker struct{}

func (Speaker) On()          { fmt.Println("音响：开机") }
func (Speaker) Volume(v int) { fmt.Println("音响：音量", v) }
func (Speaker) Off()         { fmt.Println("音响：关机") }

// Player 是播放器。
type Player struct{}

func (Player) Play(movie string) { fmt.Println("播放器：播放", movie) }
func (Player) Stop()             { fmt.Println("播放器：停止") }

// --8<-- [end:subsystem]
