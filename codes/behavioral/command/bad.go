package command

// --8<-- [start:bad]
// 反例：面板既决定何时执行，也得知道每种设备怎么启动。
func press(button string, lamp *Lamp, fan *Fan) {
	switch button {
	case "灯":
		lamp.TurnOn()
	case "风扇":
		fan.Start()
	}
}

// --8<-- [end:bad]
