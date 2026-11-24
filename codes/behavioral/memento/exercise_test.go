package memento_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/memento"
)

// Limited 是只保留最近 max 份快照的管理者。
type Limited struct {
	*memento.Editor
	max    int
	points []memento.Memento
}

func (l *Limited) Checkpoint() {
	l.points = append(l.points, l.Save())
	if len(l.points) > l.max {
		l.points = l.points[1:] // 丢掉最早一份
	}
}

func (l *Limited) Undo() bool {
	if len(l.points) == 0 {
		return false
	}
	last := len(l.points) - 1
	l.Restore(l.points[last])
	l.points = l.points[:last]
	return true
}

// 练习 1：只保留最近 N 步的历史
func TestExercise1_Limited(t *testing.T) {
	l := &Limited{Editor: &memento.Editor{}, max: 2}
	for _, title := range []string{"A", "B", "C", "D"} {
		l.Checkpoint()
		l.Rename(title)
	}

	wants := []struct {
		ok    bool
		title string
	}{{true, "C"}, {true, "B"}, {false, "B"}}
	for i, w := range wants {
		ok := l.Undo()
		if ok != w.ok || l.Title() != w.title {
			t.Errorf("第 %d 次撤销：got %v %s, want %v %s", i+1, ok, l.Title(), w.ok, w.title)
		}
	}
}

// 练习 2：撤销边界划错了会怎样
func TestExercise2_Boundary(t *testing.T) {
	e := &memento.Editor{}
	e.Rename("周末计划")
	e.Append("去公园")
	h := memento.NewHistory(e)

	h.Checkpoint()
	e.Rename("周日计划")
	h.Checkpoint() // 多留了一次档
	e.ReplaceLine(0, "去书店")

	for i, want := range []string{"周日计划：去公园", "周末计划：去公园"} {
		h.Undo()
		if got := e.Title() + "：" + e.Text(); got != want {
			t.Errorf("第 %d 次撤销：got %q, want %q", i+1, got, want)
		}
	}
}
