package prototype

import (
	"fmt"
	"testing"
)

// --8<-- [start:shallow-demo]
func Example_shallowCopy() {
	a := newGoblin()
	b := shallowCopy(a)
	b.Name = "精英哥布林"   // 普通字段：互不影响
	b.Skills[0] = "重击" // 切片：改到同一个底层数组
	b.Drops["金币"] = 20 // map：本来就是同一个 map
	b.Pos.X = 10       // 指针：指向同一个 Position
	fmt.Println(a)
	// Output: 哥布林 HP=30 技能=[重击 逃跑] 掉落=map[木棒:1 金币:20] 位置=(10,0)
}

// --8<-- [end:shallow-demo]

func TestCloneIsDeep(t *testing.T) {
	a := newGoblin()
	b := a.Clone()
	b.Name = "x"
	b.Skills[0] = "x"
	b.Drops["金币"] = 99
	b.Pos.X = 99
	if got, want := a.String(), newGoblin().String(); got != want {
		t.Errorf("原对象被改动：got %q, want %q", got, want)
	}
}

func TestCloneNilFields(t *testing.T) {
	c := (&Monster{Name: "空"}).Clone()
	if c.Skills != nil || c.Drops != nil || c.Pos != nil {
		t.Errorf("nil 字段应保持 nil：%+v", c)
	}
	if got, want := c.String(), "空 HP=0 技能=[] 掉落=map[] 位置=无"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestRegisterCopiesInput(t *testing.T) {
	r := NewRegistry()
	m := newGoblin()
	r.Register("goblin", m)
	m.HP = 1 // 登记后再改原始对象
	got, err := r.Spawn("goblin")
	if err != nil {
		t.Fatal(err)
	}
	if got.HP != 30 {
		t.Errorf("HP = %d, want 30", got.HP)
	}
}

func TestSpawnUnknown(t *testing.T) {
	if _, err := NewRegistry().Spawn("dragon"); err == nil {
		t.Error("未登记的名字应当返回错误")
	}
}

func TestBadDuplicates(t *testing.T) {
	g, e := newGoblin(), newEliteGoblin()
	if fmt.Sprint(g.Skills, g.Drops) != fmt.Sprint(e.Skills, e.Drops) {
		t.Error("反例中两种怪物的技能和掉落应当相同（照抄）")
	}
}
