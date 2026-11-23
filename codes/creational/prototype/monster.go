// Package prototype 演示原型模式：复制一个配好的对象来创建新对象，而不是从零构造。
package prototype

import (
	"fmt"
	"maps"
	"slices"
)

// --8<-- [start:monster]
// Position 是怪物在地图上的位置。
type Position struct{ X, Y int }

// Monster 是原型：字段多，从零配置很繁琐。
type Monster struct {
	Name   string
	HP     int
	Skills []string       // 切片
	Drops  map[string]int // map：掉落物 → 数量
	Pos    *Position      // 指针
}

// --8<-- [end:monster]

func (m *Monster) String() string {
	return fmt.Sprintf("%s HP=%d 技能=%v 掉落=%v 位置=(%d,%d)",
		m.Name, m.HP, m.Skills, m.Drops, m.Pos.X, m.Pos.Y)
}

// --8<-- [start:clone]
// Clone 返回深拷贝：切片、map、指针指向的数据也各复制一份。
func (m *Monster) Clone() *Monster {
	c := *m // 先整体复制：Name、HP 这类普通字段已经独立
	c.Skills = slices.Clone(m.Skills)
	c.Drops = maps.Clone(m.Drops)
	if m.Pos != nil {
		p := *m.Pos
		c.Pos = &p
	}
	return &c
}

// --8<-- [end:clone]
