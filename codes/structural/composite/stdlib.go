package composite

import (
	"errors"
	"fmt"
)

// --8<-- [start:errors]
var ErrDiskFull = errors.New("磁盘已满")

// saveAll 保存多个文件，把所有失败组合成一个 error 返回。
func saveAll() error {
	e1 := fmt.Errorf("保存 a.txt：%w", ErrDiskFull)
	e2 := errors.New("保存 b.txt：权限不足")
	return errors.Join(e1, e2) // 两个错误组合成一个，它仍然是 error
}

// 调用方不管拿到的是一个错误还是一组错误，都这样检查：
// errors.Is 会沿着错误树一层层往下找。
func isDiskFull(err error) bool {
	return errors.Is(err, ErrDiskFull)
}

// --8<-- [end:errors]
