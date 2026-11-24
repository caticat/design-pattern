package prototype_test

import (
	"slices"
	"testing"

	"github.com/caticat/design-pattern/codes/creational/prototype"
)

// newGoblinRegistry 返回登记好「哥布林」原型的注册表，与本章示例相同。
func newGoblinRegistry() *prototype.Registry {
	r := prototype.NewRegistry()
	r.Register("goblin", &prototype.Monster{
		Name:   "哥布林",
		HP:     30,
		Skills: []string{"挥砍"},
		Drops:  map[string]int{"金币": 5},
		Pos:    &prototype.Position{},
	})
	return r
}

// 练习 1：运行时派生一个新原型
func TestExercise1_DeriveElite(t *testing.T) {
	r := newGoblinRegistry()

	// 不写新构造函数：取一只、改一改、登记成新原型
	m, _ := r.Spawn("goblin")
	m.Name, m.HP = "精英哥布林", 60
	m.Skills = append(m.Skills, "怒吼")
	r.Register("elite", m)

	e1, _ := r.Spawn("elite")
	e2, _ := r.Spawn("elite")
	e1.Skills[0] = "重击"
	if e2.Skills[0] != "挥砍" {
		t.Fatalf("e2.Skills = %v，不应受 e1 影响", e2.Skills)
	}
	if got, want := e2.String(), "精英哥布林 HP=60 技能=[挥砍 怒吼] 掉落=map[金币:5] 位置=(0,0)"; got != want {
		t.Fatalf("elite = %q，期望 %q", got, want)
	}

	g, _ := r.Spawn("goblin")
	if got, want := g.String(), "哥布林 HP=30 技能=[挥砍] 掉落=map[金币:5] 位置=(0,0)"; got != want {
		t.Fatalf("goblin = %q，期望 %q", got, want)
	}
}

// 练习 2：复制一支怪物小队
func TestExercise2_CloneTeam(t *testing.T) {
	r := newGoblinRegistry()
	newTeam := func() []*prototype.Monster {
		a, _ := r.Spawn("goblin")
		b, _ := r.Spawn("goblin")
		return []*prototype.Monster{a, b}
	}

	// slices.Clone 只复制切片，两边的元素指向同一批怪物
	team := newTeam()
	shallow := slices.Clone(team)
	shallow[0].HP = 99
	if team[0].HP != 99 {
		t.Fatalf("slices.Clone 后 team[0].HP = %d，期望跟着变成 99", team[0].HP)
	}

	// 逐个 Clone：每只怪物都是新的
	team = newTeam()
	deep := make([]*prototype.Monster, len(team))
	for i, m := range team {
		deep[i] = m.Clone()
	}
	deep[0].HP = 99
	if team[0].HP != 30 {
		t.Fatalf("逐个 Clone 后 team[0].HP = %d，期望保持 30", team[0].HP)
	}
}
