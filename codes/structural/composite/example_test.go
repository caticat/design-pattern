package composite_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/structural/composite"
)

// --8<-- [start:example]
// report 只认识 Node，文件和目录都能传进来。
func report(n composite.Node) {
	fmt.Printf("%s：%d KB\n", n.Name(), n.Size())
}

func Example() {
	photos := composite.NewDir("照片",
		composite.NewFile("猫.jpg", 300),
		composite.NewFile("狗.jpg", 500),
	)
	note := composite.NewFile("笔记.txt", 20)
	home := composite.NewDir("家", note, photos) // 目录里套目录

	report(note)
	report(photos)
	report(home)

	photos.Add(composite.NewFile("鸟.jpg", 100)) // 树改了，再问一次
	report(home)
	// Output:
	// 笔记.txt：20 KB
	// 照片：800 KB
	// 家：820 KB
	// 家：920 KB
}

// --8<-- [end:example]
