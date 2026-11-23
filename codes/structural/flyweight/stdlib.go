package flyweight

import (
	"fmt"
	"strings"
	"unique"
)

// --8<-- [start:unique]
func internDemo() {
	a := unique.Make("上海")
	b := unique.Make(strings.Clone("上海")) // 另一份内容相同的字符串

	fmt.Println(a == b)    // true：同一个句柄，比较只比一个指针
	fmt.Println(a.Value()) // 上海：取回值
}

// --8<-- [end:unique]
