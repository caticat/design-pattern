package proxy_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/structural/proxy"
)

// --8<-- [start:example]
func Example() {
	// 首页列出两部视频：只创建代理，什么都没加载
	sea := proxy.NewVideoProxy("海底世界", false)
	space := proxy.NewVideoProxy("星际旅行", true)
	list := []proxy.Video{sea, space} // 代理可以当作 Video 使用
	fmt.Println("首页已打开，共", len(list), "部")

	guest := proxy.User{Name: "小明"}
	vip := proxy.User{Name: "小红", VIP: true}

	_, err := space.Play(guest) // 被拒绝：不会加载
	fmt.Println("出错:", err)

	for range 2 { // 第一次加载，第二次直接复用
		msg, _ := space.Play(vip)
		fmt.Println(msg)
	}
	// Output:
	// 首页已打开，共 2 部
	// 出错: 小明 不是会员，不能看 星际旅行
	// 加载 星际旅行（很慢）
	// 小红 正在看 星际旅行
	// 小红 正在看 星际旅行
}

// --8<-- [end:example]
