package iterator_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/iterator"
)

// --8<-- [start:example]
func Example() {
	p := &iterator.Playlist{}
	p.Add("晨光")
	p.Add("雨夜")
	p.Add("海风")

	// 经典迭代器：HasNext + Next
	for it := p.Iterator(); it.HasNext(); {
		fmt.Println("经典：" + it.Next())
	}

	// Go 1.23：直接 range，break 也能用
	for title := range p.All() {
		if title == "海风" {
			break
		}
		fmt.Println("range：" + title)
	}
	// Output:
	// 经典：晨光
	// 经典：雨夜
	// 经典：海风
	// range：晨光
	// range：雨夜
}

// --8<-- [end:example]
