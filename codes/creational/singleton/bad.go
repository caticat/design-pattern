package singleton

// --8<-- [start:bad]
var naive *Config

// NaiveGetInstance 是反例：多个 goroutine 同时调用时，
// 都可能看到 naive == nil，于是配置被加载多次，还存在数据竞争。
func NaiveGetInstance() *Config {
	if naive == nil {
		naive = loadConfig()
	}
	return naive
}

// --8<-- [end:bad]
