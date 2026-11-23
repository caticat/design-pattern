package facade

import (
	"io"
	"os"
)

// --8<-- [start:manual]
// readManual 不用外观：自己打开、读取、关闭。
func readManual(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close() // 忘了关：文件句柄泄漏
	return io.ReadAll(f)
}

// readFacade 用标准库提供的外观：一行搞定，细节由 os 包负责。
func readFacade(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// --8<-- [end:manual]
