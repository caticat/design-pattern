package composite

import (
	"fmt"
	"testing"
)

// 反例算出的结果一样，只是调用方要自己区分文件和目录。
func Example_bad() {
	photos := &badDir{files: []badFile{{300}, {500}}}
	home := &badDir{files: []badFile{{20}}, dirs: []*badDir{photos}}
	fmt.Println(totalBad(home))
	// Output:
	// 820
}

func TestEmptyAndDeep(t *testing.T) {
	if got := NewDir("空").Size(); got != 0 {
		t.Fatalf("空目录 = %d", got)
	}
	// 三层嵌套：a/b/c/x
	a := NewDir("a", NewDir("b", NewDir("c", NewFile("x", 7))), NewFile("y", 3))
	if got := a.Size(); got != 10 {
		t.Fatalf("a = %d, 期望 10", got)
	}
}

func TestErrorsTree(t *testing.T) {
	err := saveAll()
	if !isDiskFull(err) {
		t.Fatal("errors.Is 应在组合错误里找到 ErrDiskFull")
	}
	if isDiskFull(fmt.Errorf("外层：%w", fmt.Errorf("别的错误"))) {
		t.Fatal("不该匹配")
	}
}
