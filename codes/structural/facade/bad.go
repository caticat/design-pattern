package facade

// --8<-- [start:bad]
// 反例：调用方自己按顺序操作每个设备。
// 手机 App、语音助手……每个想看电影的入口都要抄一遍这 6 步。
func watchBad(movie string) {
	var (
		lights    Lights
		projector Projector
		speaker   Speaker
		player    Player
	)
	lights.Dim(10)
	projector.On()
	projector.Input("播放器") // 忘了这一步：画面是黑的
	speaker.On()
	speaker.Volume(30)
	player.Play(movie)
}

// --8<-- [end:bad]
