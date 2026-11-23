package proxy

import "fmt"

// --8<-- [start:proxy]
// VideoProxy 是代理：只记下片名，真实视频等到第一次播放才加载。
type VideoProxy struct {
	title   string
	vipOnly bool
	real    *HDVideo // 第一次播放前为 nil
}

// NewVideoProxy 创建代理：很便宜，此时什么都不加载。
func NewVideoProxy(title string, vipOnly bool) *VideoProxy {
	return &VideoProxy{title: title, vipOnly: vipOnly}
}

// --8<-- [end:proxy]

// --8<-- [start:play]
func (p *VideoProxy) Play(u User) (string, error) {
	// 1. 保护：不满足条件，直接拒绝，真实对象根本不会被碰到
	if p.vipOnly && !u.VIP {
		return "", fmt.Errorf("%s 不是会员，不能看 %s", u.Name, p.title)
	}
	// 2. 延迟加载：第一次真正需要时才创建，之后一直复用
	if p.real == nil {
		p.real = LoadHDVideo(p.title)
	}
	// 3. 转交给真实对象
	return p.real.Play(u)
}

// --8<-- [end:play]
