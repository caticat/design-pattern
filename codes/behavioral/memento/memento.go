// Package memento 用文本编辑器的撤销演示备忘录模式。
package memento

import (
	"slices"
	"strings"
)

// --8<-- [start:editor]
// Editor 保存一份文稿的标题和正文。
type Editor struct {
	title string
	lines []string
}

// Rename 修改标题。
func (e *Editor) Rename(title string) { e.title = title }

// Title 返回当前标题。
func (e *Editor) Title() string { return e.title }

// Append 在文末添加一行。
func (e *Editor) Append(line string) { e.lines = append(e.lines, line) }

// ReplaceLine 改写一行，从 0 开始计数；行号无效时返回 false。
func (e *Editor) ReplaceLine(i int, line string) bool {
	if i < 0 || i >= len(e.lines) {
		return false
	}
	e.lines[i] = line
	return true
}

// Text 返回正文，各行之间用换行分隔。
func (e *Editor) Text() string { return strings.Join(e.lines, "\n") }

// --8<-- [end:editor]

// --8<-- [start:snapshot]
// Memento 是一份文稿快照，内容不对包外公开；零值表示空稿。
type Memento struct {
	title string
	lines []string
}

// Save 保存当前标题和正文，不与编辑器共享行切片。
func (e *Editor) Save() Memento {
	return Memento{title: e.title, lines: slices.Clone(e.lines)}
}

// Restore 恢复标题和正文，不修改传入的快照。
func (e *Editor) Restore(m Memento) {
	e.title = m.title
	e.lines = slices.Clone(m.lines)
}

// --8<-- [end:snapshot]

// --8<-- [start:history]
// History 管理一个编辑器的撤销记录。请用 NewHistory 创建。
type History struct {
	editor *Editor
	points []Memento
}

// NewHistory 为指定编辑器创建一份空的撤销记录。
func NewHistory(editor *Editor) *History { return &History{editor: editor} }

// Checkpoint 在编辑前把当前文稿压入历史栈。
func (h *History) Checkpoint() {
	h.points = append(h.points, h.editor.Save())
}

// Undo 取出最近的快照并恢复；没有记录时返回 false。
func (h *History) Undo() bool {
	if len(h.points) == 0 {
		return false
	}
	last := len(h.points) - 1
	h.editor.Restore(h.points[last])
	h.points[last] = Memento{} // 释放历史栈对这份快照的引用
	h.points = h.points[:last]
	return true
}

// --8<-- [end:history]
