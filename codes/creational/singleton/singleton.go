// Package singleton 演示单例模式：保证一个类型只有一个实例，并提供全局访问点。
package singleton

import (
	"sync"
	"sync/atomic"
)

// Config 是应用的全局配置，整个程序只应加载一次。
type Config struct {
	Addr  string
	Debug bool
}

// loads 记录配置被真正加载的次数，用于验证「只加载一次」。
var loads atomic.Int32

// loadConfig 模拟一次昂贵的加载：读文件、解析、校验……
func loadConfig() *Config {
	loads.Add(1)
	return &Config{Addr: ":8080", Debug: true}
}

// Loads 返回配置被加载的次数。
func Loads() int32 { return loads.Load() }

// --8<-- [start:classic]
var (
	instance *Config
	once     sync.Once
)

// GetInstance 返回全局唯一的 Config。
// sync.Once 保证即使被多个 goroutine 同时调用，loadConfig 也只执行一次。
func GetInstance() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

// --8<-- [end:classic]

// --8<-- [start:idiom]
// GetConfig 是 Go 1.21+ 的惯用写法：sync.OnceValue 把
// 「Once + 包级变量 + 访问函数」三件套合成一行。
var GetConfig = sync.OnceValue(loadConfig)

// --8<-- [end:idiom]
