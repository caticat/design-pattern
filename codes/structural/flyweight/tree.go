// Package flyweight 演示享元模式：大量对象共用不变的那部分数据。
package flyweight

import "fmt"

// --8<-- [start:type]
// TreeType 是享元：同一种树只有一份，创建后不再修改。
type TreeType struct {
	name    string
	texture []byte // 贴图，很大
}

// Draw 画一棵树。位置每棵树都不同，由调用方传进来，不存在享元里。
func (t *TreeType) Draw(x, y int) string {
	return fmt.Sprintf("%s@(%d,%d)", t.name, x, y)
}

// --8<-- [end:type]

func loadTexture(name string) []byte {
	fmt.Println("加载贴图：" + name)
	return make([]byte, 64<<10) // 示意：真实贴图可能有几 MB
}
