package proxy

import "fmt"

// --8<-- [start:bad]
// 反例：首页一打开就加载全部视频，用户可能一部都不看。
func openHomeBad(titles []string) []*HDVideo {
	videos := make([]*HDVideo, 0, len(titles))
	for _, t := range titles {
		videos = append(videos, LoadHDVideo(t))
	}
	return videos
}

// 反例：会员检查写在调用处，每个播放入口都要记得抄一遍。
func playBad(v *HDVideo, u User, vipOnly bool) (string, error) {
	if vipOnly && !u.VIP {
		return "", fmt.Errorf("%s 不是会员", u.Name)
	}
	return v.Play(u)
}

// --8<-- [end:bad]
