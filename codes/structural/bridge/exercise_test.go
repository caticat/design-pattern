package bridge_test

import (
	"io"
	"os"
	"testing"

	"github.com/caticat/design-pattern/codes/structural/bridge"
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

// Scheduled 是抽象一边新增的「定时」消息。
type Scheduled struct {
	At     string
	Sender bridge.Sender
}

func (s Scheduled) Notify(msg string) {
	s.Sender.Send("【" + s.At + " 定时】" + msg)
}

var _ bridge.Notice = Scheduled{}

// 练习 1：在抽象一边加一种「定时」消息
func TestExercise1_Scheduled(t *testing.T) {
	out := captureOutput(t, func() {
		Scheduled{At: "09:00", Sender: bridge.Email{}}.Notify("周报已生成")
		Scheduled{At: "09:00", Sender: bridge.SMS{}}.Notify("周报已生成")
	})
	want := "邮件：【09:00 定时】周报已生成\n" +
		"短信：【09:00 定时】周报已生成\n"
	if out != want {
		t.Fatalf("输出 %q，期望 %q", out, want)
	}
}

// Switch 是可切换的渠道：Send 转给当前的 S。
type Switch struct{ S bridge.Sender }

func (s *Switch) Send(text string) { s.S.Send(text) }

// 练习 2：不改包，也能运行时换渠道
func TestExercise2_Switch(t *testing.T) {
	out := captureOutput(t, func() {
		sw := &Switch{S: bridge.SMS{}}
		n := bridge.NewNormal(sw)
		n.Notify("值班交接")
		sw.S = bridge.Email{} // 运行时换渠道
		n.Notify("值班交接")
	})
	if want := "短信：值班交接\n邮件：值班交接\n"; out != want {
		t.Fatalf("输出 %q，期望 %q", out, want)
	}
}
