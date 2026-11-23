package flyweight

// --8<-- [start:factory]
// TreeFactory 按名字缓存 TreeType：同名只创建一次，之后都返回同一个。
type TreeFactory struct {
	types map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{types: map[string]*TreeType{}}
}

func (f *TreeFactory) Get(name string) *TreeType {
	if t, ok := f.types[name]; ok {
		return t // 已有，直接共用
	}
	t := &TreeType{name: name, texture: loadTexture(name)}
	f.types[name] = t
	return t
}

// --8<-- [end:factory]

// --8<-- [start:forest]
// Tree 只存每棵树自己的位置，种类指向共享的 TreeType。
type Tree struct {
	x, y int
	kind *TreeType
}

func (t Tree) Draw() string { return t.kind.Draw(t.x, t.y) }

// Forest 种树时通过工厂拿 TreeType。
type Forest struct {
	factory *TreeFactory
	trees   []Tree
}

func NewForest() *Forest { return &Forest{factory: NewTreeFactory()} }

func (f *Forest) Plant(x, y int, name string) {
	f.trees = append(f.trees, Tree{x: x, y: y, kind: f.factory.Get(name)})
}

func (f *Forest) Trees() int { return len(f.trees) }
func (f *Forest) Kinds() int { return len(f.factory.types) }

// --8<-- [end:forest]
