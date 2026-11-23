package memento

import "fmt"

// --8<-- [start:bad]
// 反例：看似存了一份旧稿，其实行切片仍和编辑器共用底层数组。
func badUndo() string {
	e := &Editor{}
	e.Rename("周末计划")
	e.Append("去公园")

	oldTitle, oldLines := e.title, e.lines
	e.Rename("周日计划")
	e.ReplaceLine(0, "去书店")

	e.title, e.lines = oldTitle, oldLines
	return fmt.Sprintf("%s：%s", e.Title(), e.Text()) // 周末计划：去书店
}

// --8<-- [end:bad]
