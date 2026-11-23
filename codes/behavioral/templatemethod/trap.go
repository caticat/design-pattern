package templatemethod

// --8<-- [start:trap]
// 反例：照搬 Java 的「父类定流程，子类覆盖步骤」。
type baseExporter struct{}

func (baseExporter) Header() string { return "（默认表头）" }
func (b baseExporter) Run() string  { return b.Header() } // b 永远是 baseExporter

type csvExporter struct{ baseExporter } // 嵌入，想「继承」Run

func (csvExporter) Header() string { return "name,score" } // 以为覆盖了父类

// csvExporter{}.Run() 返回「（默认表头）」，而不是 "name,score"。
// --8<-- [end:trap]
