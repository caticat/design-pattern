package proxy_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/caticat/design-pattern/codes/structural/proxy"
)

// captureOutput 运行 f，返回它打印到标准输出的内容。
func captureOutput(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = old }()

	f()
	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

// LogProxy 先记日志，再把请求转交给里面的 Video。
type LogProxy struct{ V proxy.Video }

func (p LogProxy) Play(u proxy.User) (string, error) {
	fmt.Printf("[日志] %s 请求播放\n", u.Name)
	return p.V.Play(u)
}

// 练习 1：加一个日志代理
func TestExercise1_LogProxy(t *testing.T) {
	var v proxy.Video = LogProxy{V: proxy.NewVideoProxy("星际旅行", true)} // 调用方只拿着 Video

	out := captureOutput(t, func() {
		msg, err := v.Play(proxy.User{Name: "小红", VIP: true})
		if err != nil {
			t.Error(err)
		}
		fmt.Println(msg)
	})
	want := "[日志] 小红 请求播放\n" +
		"加载 星际旅行（很慢）\n" +
		"小红 正在看 星际旅行\n"
	if out != want {
		t.Fatalf("输出 %q，期望 %q", out, want)
	}
}

// 练习 2「让延迟加载撞上并发」需改动包内 VideoProxy 并用 -race 观察，不在此给出。
