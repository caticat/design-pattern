package flyweight

// --8<-- [start:bad]
// 反例：每棵树都带一份自己的贴图。
type badTree struct {
	x, y    int
	name    string
	texture []byte
}

func plantBad(n int, name string) []badTree {
	trees := make([]badTree, n)
	for i := range trees {
		// 每棵都重新加载一次，贴图内容其实完全一样
		trees[i] = badTree{x: i, y: i, name: name, texture: loadTexture(name)}
	}
	return trees
}

// 10 万棵树 × 每份 1 MB 贴图 ≈ 100 GB 内存。
// --8<-- [end:bad]
