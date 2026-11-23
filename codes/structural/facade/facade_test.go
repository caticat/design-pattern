package facade

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// 反例和外观做的事完全一样，区别只在谁来记住这些步骤。
func Example_bad() {
	watchBad("星际旅行")
	// Output:
	// 灯光：调到 10%
	// 投影仪：开机
	// 投影仪：信号源切到 播放器
	// 音响：开机
	// 音响：音量 30
	// 播放器：播放 星际旅行
}

// 外观不是墙：需要精细控制时，仍可直接使用子系统。
func Example_subsystem() {
	var s Speaker
	s.Volume(50)
	// Output:
	// 音响：音量 50
}

func TestReadSame(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.txt")
	want := []byte("hello")
	if err := os.WriteFile(path, want, 0o644); err != nil {
		t.Fatal(err)
	}
	a, err := readManual(path)
	if err != nil {
		t.Fatal(err)
	}
	b, err := readFacade(path)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(a, want) || !bytes.Equal(b, want) {
		t.Fatalf("got %q %q", a, b)
	}
}
