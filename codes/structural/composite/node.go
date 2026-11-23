// Package composite 演示组合模式：单个对象和由对象组成的树，调用方一视同仁。
package composite

// --8<-- [start:node]
// Node 是文件和目录的共同接口。
type Node interface {
	Name() string
	Size() int64 // 单位 KB
}

// --8<-- [end:node]

// --8<-- [start:leaf]
// File 是叶子：没有子节点，大小就是它自己的大小。
type File struct {
	name string
	size int64
}

func NewFile(name string, size int64) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string { return f.name }
func (f *File) Size() int64  { return f.size }

// --8<-- [end:leaf]
