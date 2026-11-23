package composite

// --8<-- [start:bad]
// 反例：文件和目录是两种不相干的类型，调用方得分开处理。
type badFile struct{ size int64 }

type badDir struct {
	files []badFile
	dirs  []*badDir
}

func totalBad(d *badDir) int64 {
	var total int64
	for _, f := range d.files { // 文件一种循环
		total += f.size
	}
	for _, sub := range d.dirs { // 目录另一种循环，还得自己递归
		total += totalBad(sub)
	}
	return total
}

// 再加一种「快捷方式」？badDir 多一个切片，totalBad 多一个循环，
// 数文件个数、找最大文件……每个统计函数都要跟着改。
// --8<-- [end:bad]
