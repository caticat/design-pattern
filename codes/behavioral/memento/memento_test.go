package memento_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/memento"
)

func TestHistoryUndoesInReverseOrder(t *testing.T) {
	e := &memento.Editor{}
	e.Rename("周末计划")
	e.Append("去公园")
	h := memento.NewHistory(e)

	h.Checkpoint()
	e.Rename("周日计划")
	e.ReplaceLine(0, "去书店")
	h.Checkpoint()
	e.Append("喝咖啡")

	for _, want := range []struct{ title, text string }{
		{"周日计划", "去书店"},
		{"周末计划", "去公园"},
	} {
		if !h.Undo() || e.Title() != want.title || e.Text() != want.text {
			t.Fatalf("撤销后得到 %q / %q，想要 %q / %q", e.Title(), e.Text(), want.title, want.text)
		}
	}
	if h.Undo() || e.Title() != "周末计划" || e.Text() != "去公园" {
		t.Fatal("空历史不应改动文稿")
	}
}

func TestSnapshotSurvivesEditsBeforeAndAfterRestore(t *testing.T) {
	e := &memento.Editor{}
	e.Append("去公园")
	saved := e.Save()

	for _, line := range []string{"去书店", "去博物馆"} {
		if !e.ReplaceLine(0, line) {
			t.Fatal("第一行应存在")
		}
		e.Restore(saved)
		if got := e.Text(); got != "去公园" {
			t.Fatalf("改为 %q 后，恢复得到 %q，想要 %q", line, got, "去公园")
		}
	}
}

func TestReplaceLineRejectsInvalidIndex(t *testing.T) {
	e := &memento.Editor{}
	e.Append("去公园")
	if e.ReplaceLine(-1, "改写") || e.ReplaceLine(1, "改写") || e.Text() != "去公园" {
		t.Fatal("非法行号不应改动文稿")
	}
}
