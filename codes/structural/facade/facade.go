package facade

// --8<-- [start:facade]
// HomeTheater 是外观：持有全部设备，对外只给两个按钮。
type HomeTheater struct {
	lights    Lights
	projector Projector
	speaker   Speaker
	player    Player
}

// NewHomeTheater 创建家庭影院（真实项目中在这里创建或接收各个设备）。
func NewHomeTheater() *HomeTheater {
	return &HomeTheater{}
}

// --8<-- [end:facade]

// --8<-- [start:methods]
// Watch 一键看电影：按正确的顺序操作每个设备。
func (h *HomeTheater) Watch(movie string) {
	h.lights.Dim(10)
	h.projector.On()
	h.projector.Input("播放器")
	h.speaker.On()
	h.speaker.Volume(30)
	h.player.Play(movie)
}

// End 一键结束：顺序和开启时相反。
func (h *HomeTheater) End() {
	h.player.Stop()
	h.speaker.Off()
	h.projector.Off()
	h.lights.Dim(100)
}

// --8<-- [end:methods]
