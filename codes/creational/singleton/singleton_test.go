package singleton

import (
	"sync"
	"testing"
)

// callConcurrently 让 n 个 goroutine 同时调用 get，返回所有结果。
func callConcurrently(n int, get func() *Config) []*Config {
	var wg sync.WaitGroup
	results := make([]*Config, n)
	for i := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = get()
		}()
	}
	wg.Wait()
	return results
}

func assertSame(t *testing.T, results []*Config) {
	t.Helper()
	for i, c := range results {
		if c == nil || c != results[0] {
			t.Fatalf("第 %d 次调用拿到了不同的实例", i)
		}
	}
}

func TestSingleInstance(t *testing.T) {
	before := Loads()
	assertSame(t, callConcurrently(100, GetInstance))
	assertSame(t, callConcurrently(100, GetConfig))

	// 两种写法各自只加载一次（可能已被其它测试触发过，故用 <=）。
	if got := Loads() - before; got > 2 {
		t.Fatalf("配置被加载了 %d 次，期望最多 2 次", got)
	}
}
