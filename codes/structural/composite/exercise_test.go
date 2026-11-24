package composite_test

import (
	"io"
	"os"
	"testing"

	"github.com/caticat/design-pattern/codes/structural/composite"
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

// Shortcut 是新增的一种节点：快捷方式，自身只占 1 KB。
type Shortcut struct{ Target string }

func (s Shortcut) Name() string { return s.Target + " 的快捷方式" }
func (s Shortcut) Size() int64  { return 1 }

// 练习 1：加一种节点：快捷方式
func TestExercise1_Shortcut(t *testing.T) {
	photos := composite.NewDir("照片",
		composite.NewFile("猫.jpg", 300),
		composite.NewFile("狗.jpg", 500),
	)
	home := composite.NewDir("家", composite.NewFile("笔记.txt", 20), photos)
	home.Add(Shortcut{Target: "照片"})

	out := captureOutput(t, func() { report(home) }) // report 定义在 example_test.go
	if want := "家：821 KB\n"; out != want {
		t.Fatalf("输出 %q，期望 %q", out, want)
	}
}

// 练习 2「亲手造一个环」会以 fatal error: stack overflow 崩溃，无法写成测试。
