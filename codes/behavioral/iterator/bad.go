package iterator

import "fmt"

// --8<-- [start:bad]
// 反例：为了让外面能遍历，把链表节点暴露了出去。
type songNode struct {
	Title string
	Next  *songNode
}

type badPlaylist struct{ Head *songNode }

// 调用方：每个要遍历的地方都这样手写指针移动。
func printTitles(p *badPlaylist) {
	for n := p.Head; n != nil; n = n.Next {
		fmt.Println(n.Title)
	}
}

// --8<-- [end:bad]
