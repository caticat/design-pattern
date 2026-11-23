package prototype_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/creational/prototype"
)

// --8<-- [start:example]
func Example() {
	r := prototype.NewRegistry()
	r.Register("goblin", &prototype.Monster{
		Name:   "哥布林",
		HP:     30,
		Skills: []string{"挥砍"},
		Drops:  map[string]int{"金币": 5},
		Pos:    &prototype.Position{},
	})

	// 复制一只，只改不同的部分
	elite, _ := r.Spawn("goblin")
	elite.Name, elite.HP = "精英哥布林", 60
	elite.Skills[0] = "重击"
	elite.Drops["金币"] = 20
	elite.Pos.X = 10

	// 再复制一只：原型完全没被改动
	normal, _ := r.Spawn("goblin")
	fmt.Println(elite)
	fmt.Println(normal)

	_, err := r.Spawn("dragon")
	fmt.Println("出错:", err)
	// Output:
	// 精英哥布林 HP=60 技能=[重击] 掉落=map[金币:20] 位置=(10,0)
	// 哥布林 HP=30 技能=[挥砍] 掉落=map[金币:5] 位置=(0,0)
	// 出错: 没有名为 "dragon" 的原型
}

// --8<-- [end:example]
