// Package proxy 演示代理模式：用一个同接口的替身控制对真实对象的访问。
package proxy

import "fmt"

// --8<-- [start:subject]
// User 是看视频的用户。
type User struct {
	Name string
	VIP  bool
}

// Video 是公共接口：真实视频和代理都实现它，调用方分不出是哪一个。
type Video interface {
	Play(u User) (string, error)
}

// HDVideo 是真实对象：创建时要把整部影片读进内存，很慢。
type HDVideo struct {
	title string
}

// LoadHDVideo 加载影片（真实情况要读几百 MB，这里只打印一行表示）。
func LoadHDVideo(title string) *HDVideo {
	fmt.Printf("加载 %s（很慢）\n", title)
	return &HDVideo{title: title}
}

func (v *HDVideo) Play(u User) (string, error) {
	return fmt.Sprintf("%s 正在看 %s", u.Name, v.title), nil
}

// --8<-- [end:subject]
