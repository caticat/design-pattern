package templatemethod

import "testing"

var rows = []Record{{Name: "张三", Score: 95}, {Name: "李四", Score: 82}}

// 模板方法与反例的输出必须完全一致：重构不改变行为。
func TestSameAsBad(t *testing.T) {
	for _, in := range [][]Record{rows, nil} {
		if got, want := Export(CSV{}, in), exportCSV(in); got != want {
			t.Errorf("CSV: got %q, want %q", got, want)
		}
		if got, want := Export(Markdown{}, in), exportMarkdown(in); got != want {
			t.Errorf("Markdown: got %q, want %q", got, want)
		}
	}
}

// 没实现 Footer 的格式不会多出结尾行。
func TestHookOptional(t *testing.T) {
	if got, want := Export(CSV{}, rows), "name,score\n张三,95\n李四,82\n"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// 嵌入不是继承：Run 调用的仍是 baseExporter.Header。
func TestEmbedTrap(t *testing.T) {
	if got := (csvExporter{}).Run(); got != "（默认表头）" {
		t.Errorf("got %q", got)
	}
}
