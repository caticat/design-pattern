package singleton_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/creational/singleton"
)

func Example() {
	a := singleton.GetInstance()
	b := singleton.GetInstance()

	fmt.Println("同一个实例:", a == b)
	fmt.Println("监听地址:", a.Addr)
	// Output:
	// 同一个实例: true
	// 监听地址: :8080
}
