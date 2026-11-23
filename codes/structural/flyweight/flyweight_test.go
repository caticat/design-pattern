package flyweight

import "testing"

// 反例：3 棵树加载 3 次贴图。
func Example_bad() {
	plantBad(3, "松树")
	// Output:
	// 加载贴图：松树
	// 加载贴图：松树
	// 加载贴图：松树
}

func TestShared(t *testing.T) {
	f := NewTreeFactory()
	a, b := f.Get("松树"), f.Get("松树")
	if a != b {
		t.Fatal("同名 TreeType 应是同一个对象")
	}
	if f.Get("橡树") == a {
		t.Fatal("不同名不应共用")
	}
}

func TestDraw(t *testing.T) {
	forest := NewForest()
	forest.Plant(1, 2, "松树")
	forest.Plant(3, 4, "松树")
	if got := forest.trees[0].Draw(); got != "松树@(1,2)" {
		t.Fatalf("got %q", got)
	}
	if got := forest.trees[1].Draw(); got != "松树@(3,4)" {
		t.Fatalf("got %q", got)
	}
}

func Example_unique() {
	internDemo()
	// Output:
	// true
	// 上海
}
