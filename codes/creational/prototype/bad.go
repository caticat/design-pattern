package prototype

// --8<-- [start:bad]
// 每种怪物都从零写一遍；精英版只改了两项，其余照抄。
func newGoblin() *Monster {
	return &Monster{
		Name:   "哥布林",
		HP:     30,
		Skills: []string{"挥砍", "逃跑"},
		Drops:  map[string]int{"金币": 5, "木棒": 1},
		Pos:    &Position{},
	}
}

func newEliteGoblin() *Monster {
	return &Monster{
		Name:   "精英哥布林",
		HP:     60,
		Skills: []string{"挥砍", "逃跑"},
		Drops:  map[string]int{"金币": 5, "木棒": 1},
		Pos:    &Position{},
	}
}

// --8<-- [end:bad]

// --8<-- [start:shallow]
// shallowCopy 直接解引用赋值：只复制了一层。
func shallowCopy(m *Monster) *Monster {
	c := *m
	return &c
}

// --8<-- [end:shallow]
