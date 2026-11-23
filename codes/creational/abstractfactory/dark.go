package abstractfactory

// --8<-- [start:dark]
// Dark 是深色主题工厂，只生产黑底控件。
type Dark struct{}

func (Dark) NewButton(label string) Button     { return darkButton{label} }
func (Dark) NewCheckbox(label string) Checkbox { return darkCheckbox{label} }

type darkButton struct{ label string }
type darkCheckbox struct{ label string }

func (b darkButton) Render() string   { return "[黑底按钮 " + b.label + "]" }
func (c darkCheckbox) Render() string { return "[黑底勾选框 " + c.label + "]" }

// --8<-- [end:dark]
