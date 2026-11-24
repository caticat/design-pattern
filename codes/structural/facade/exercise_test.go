package facade_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/caticat/design-pattern/codes/structural/facade"
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

// Karaoke 是另一个外观，复用同样的子系统。
type Karaoke struct {
	lights  facade.Lights
	speaker facade.Speaker
	player  facade.Player
}

func (k *Karaoke) Sing(song string) {
	k.lights.Dim(50)
	k.speaker.On()
	k.speaker.Volume(60)
	k.player.Play(song)
}

// 练习 1：再做一个「K 歌」外观
func TestExercise1_Karaoke(t *testing.T) {
	out := captureOutput(t, func() {
		var k Karaoke // 子系统都是零值结构体，不用初始化
		k.Sing("小幸运")
	})
	want := "灯光：调到 50%\n" +
		"音响：开机\n" +
		"音响：音量 60\n" +
		"播放器：播放 小幸运\n"
	if out != want {
		t.Fatalf("输出 %q，期望 %q", out, want)
	}
}

// Theater 是调用方自己定义的接口，只含它用到的两个方法。
type Theater interface {
	Watch(movie string)
	End()
}

// MovieNight 只依赖 Theater 接口，先看再关。
func MovieNight(t Theater, movie string) {
	t.Watch(movie)
	t.End()
}

// fakeTheater 只记录调用，不碰任何设备。
type fakeTheater struct{ calls []string }

func (f *fakeTheater) Watch(movie string) { f.calls = append(f.calls, "Watch "+movie) }
func (f *fakeTheater) End()               { f.calls = append(f.calls, "End") }

// 练习 2：测试时替换掉整个外观
func TestExercise2_FakeTheater(t *testing.T) {
	fake := &fakeTheater{}
	out := captureOutput(t, func() { MovieNight(fake, "星际旅行") })
	if out != "" {
		t.Fatalf("fakeTheater 不应打印设备输出，得到 %q", out)
	}
	if got := fmt.Sprint(fake.calls); got != "[Watch 星际旅行 End]" {
		t.Fatalf("calls = %s", got)
	}

	// 换成真的外观：输出与本章示例相同（少了中间那句旁白）
	out = captureOutput(t, func() { MovieNight(facade.NewHomeTheater(), "星际旅行") })
	want := "灯光：调到 10%\n" +
		"投影仪：开机\n" +
		"投影仪：信号源切到 播放器\n" +
		"音响：开机\n" +
		"音响：音量 30\n" +
		"播放器：播放 星际旅行\n" +
		"播放器：停止\n" +
		"音响：关机\n" +
		"投影仪：关机\n" +
		"灯光：调到 100%\n"
	if out != want {
		t.Fatalf("输出 %q，期望 %q", out, want)
	}
}
