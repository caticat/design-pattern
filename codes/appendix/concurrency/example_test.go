package concurrency

import (
	"fmt"
	"sync"
)

// --8<-- [start:race]
// 1000 个 goroutine 同时给计数器加 1，没有任何同步。
// 结果每次可能不同，所以不写 Output 注释（go test 只编译、不运行它）。
func Example_race() {
	var wg sync.WaitGroup
	count := 0
	for range 1000 {
		wg.Add(1)
		go func() { // go：这个函数在新的 goroutine 里并发执行
			defer wg.Done()
			count++ // 数据竞争：多个 goroutine 同时读写 count
		}()
	}
	wg.Wait()          // 等 1000 个 goroutine 全部结束
	fmt.Println(count) // 常常小于 1000
}

// --8<-- [end:race]

// --8<-- [start:mutex]
// 同样的计数，用 Mutex 保护 count++，结果一定是 1000。
func Example_mutex() {
	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		count int
	)
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // 同一时刻只有一个 goroutine 能进来
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
	// Output: 1000
}

// --8<-- [end:mutex]
