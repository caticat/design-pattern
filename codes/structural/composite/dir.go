package composite

// --8<-- [start:composite]
// Dir 是组合：子节点可以是文件，也可以是另一个目录。
type Dir struct {
	name     string
	children []Node
}

func NewDir(name string, children ...Node) *Dir {
	return &Dir{name: name, children: children}
}

// Add 添加子节点。只有目录有这个方法，文件没有。
func (d *Dir) Add(n Node) { d.children = append(d.children, n) }

func (d *Dir) Name() string { return d.name }

// Size 把每个子节点的大小加起来；子节点是目录时，它会继续往下加。
func (d *Dir) Size() int64 {
	var total int64
	for _, c := range d.children {
		total += c.Size() // 不用管 c 是文件还是目录
	}
	return total
}

// --8<-- [end:composite]
