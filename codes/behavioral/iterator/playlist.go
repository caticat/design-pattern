// Package iterator 演示迭代器模式：播放列表内部用链表保存，外部只能逐首取歌。
package iterator

import "iter"

// --8<-- [start:collection]
// node 不导出：外部看不到播放列表用的是链表。
type node struct {
	title string
	next  *node
}

type Playlist struct {
	head, tail *node
}

func (p *Playlist) Add(title string) {
	n := &node{title: title}
	if p.tail == nil {
		p.head = n
	} else {
		p.tail.next = n
	}
	p.tail = n
}

// --8<-- [end:collection]

// --8<-- [start:classic]
// Iterator 是经典的迭代器对象：记住当前走到哪一首。
type Iterator struct{ cur *node }

func (p *Playlist) Iterator() *Iterator { return &Iterator{cur: p.head} }

func (it *Iterator) HasNext() bool { return it.cur != nil }

func (it *Iterator) Next() string {
	title := it.cur.title
	it.cur = it.cur.next
	return title
}

// --8<-- [end:classic]

// --8<-- [start:seq]
// All 返回 Go 1.23 的迭代器，调用方直接 for range。
func (p *Playlist) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		for n := p.head; n != nil; n = n.next {
			if !yield(n.title) {
				return // 调用方 break 了，马上停
			}
		}
	}
}

// --8<-- [end:seq]
