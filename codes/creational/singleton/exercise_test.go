package singleton_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/caticat/design-pattern/codes/creational/singleton"
)

// Logger 是练习里的日志器。
type Logger struct{}

var loggerCreates atomic.Int32 // 记录 Logger 被创建了几次

// GetLogger 仿照 GetConfig：第一次调用时创建，之后都返回同一个。
var GetLogger = sync.OnceValue(func() *Logger {
	loggerCreates.Add(1)
	return &Logger{}
})

// 练习 1：给日志器也做一个单例
func TestExercise1_GetLogger(t *testing.T) {
	results := make([]*Logger, 100)
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = GetLogger()
		}()
	}
	wg.Wait()

	for _, l := range results {
		if l != results[0] {
			t.Fatal("拿到了不同的 Logger")
		}
	}
	if n := loggerCreates.Load(); n != 1 {
		t.Fatalf("创建了 %d 次，期望 1 次", n)
	}
	if addr := singleton.GetInstance().Addr; addr != ":8080" {
		t.Fatalf("GetInstance().Addr = %q，不应受影响", addr)
	}
}

// try 运行 f，返回它 panic 出来的值；没有 panic 时返回 nil。
func try(f func()) (p any) {
	defer func() { p = recover() }()
	f()
	return nil
}

// 练习 2：观察初始化 panic 后的两种结局
func TestExercise2_PanicInInit(t *testing.T) {
	// sync.OnceValue：记住了 panic，每次调用都再 panic 一次
	var valueRuns int
	load := sync.OnceValue(func() int {
		valueRuns++
		panic("加载失败")
	})
	p1 := try(func() { load() })
	p2 := try(func() { load() })
	if p1 != "加载失败" || p2 != "加载失败" {
		t.Fatalf("OnceValue 两次得到 %v、%v，期望都是「加载失败」", p1, p2)
	}

	// sync.Once：第一次 panic 后就算「做过了」，第二次什么也不做
	var onceRuns int
	var once sync.Once
	initFn := func() {
		onceRuns++
		panic("加载失败")
	}
	q1 := try(func() { once.Do(initFn) })
	q2 := try(func() { once.Do(initFn) })
	if q1 != "加载失败" || q2 != nil {
		t.Fatalf("Once 两次得到 %v、%v，期望「加载失败」、nil", q1, q2)
	}

	if valueRuns != 1 || onceRuns != 1 {
		t.Fatalf("初始化函数执行次数 OnceValue=%d Once=%d，期望都是 1", valueRuns, onceRuns)
	}
}
