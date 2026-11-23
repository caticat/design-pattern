package flyweight_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/structural/flyweight"
)

// --8<-- [start:example]
func Example() {
	forest := flyweight.NewForest()
	for i := range 1000 {
		forest.Plant(i, i*2, "松树")
		forest.Plant(i, i*3, "橡树")
	}
	fmt.Println("树：", forest.Trees(), "棵")
	fmt.Println("贴图：", forest.Kinds(), "份")
	// Output:
	// 加载贴图：松树
	// 加载贴图：橡树
	// 树： 2000 棵
	// 贴图： 2 份
}

// --8<-- [end:example]
