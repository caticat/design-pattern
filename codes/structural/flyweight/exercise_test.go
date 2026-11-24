package flyweight_test

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/caticat/design-pattern/codes/structural/flyweight"
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

// Sapling 是新的上下文：外部状态（位置）放在这里，享元原样复用。
type Sapling struct {
	x, y int
	kind *flyweight.TreeType
}

func (s Sapling) Draw() string { return s.kind.Draw(s.x, s.y) + " 树苗" }

// 练习 1：加一种上下文：树苗
func TestExercise1_Sapling(t *testing.T) {
	var a, b Sapling
	out := captureOutput(t, func() {
		f := flyweight.NewTreeFactory()
		a = Sapling{x: 1, y: 2, kind: f.Get("松树")}
		b = Sapling{x: 3, y: 4, kind: f.Get("松树")}
	})

	if out != "加载贴图：松树\n" {
		t.Fatalf("输出 %q，期望只加载一次", out)
	}
	if a.kind != b.kind {
		t.Fatal("两棵树苗应共用同一个享元")
	}
	if got := a.Draw(); got != "松树@(1,2) 树苗" {
		t.Fatalf("Draw = %q", got)
	}
}

// SafeFactory 在包外给工厂加锁，让多个 goroutine 可以同时 Get。
type SafeFactory struct {
	mu sync.Mutex
	f  *flyweight.TreeFactory
}

func (s *SafeFactory) Get(name string) *flyweight.TreeType {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.f.Get(name)
}

// 练习 2：让工厂撞上并发，再从外面加锁
// 只验证加锁后的 SafeFactory；直接用工厂的竞态要在 -race 下才能看到，这里不演示。
func TestExercise2_SafeFactory(t *testing.T) {
	sf := &SafeFactory{f: flyweight.NewTreeFactory()}

	out := captureOutput(t, func() {
		var wg sync.WaitGroup
		for i := range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if i%2 == 0 {
					sf.Get("松树")
				} else {
					sf.Get("橡树")
				}
			}()
		}
		wg.Wait()
	})

	if n := strings.Count(out, "加载贴图"); n != 2 {
		t.Fatalf("加载贴图 %d 次，期望恰好 2 次", n)
	}
}
