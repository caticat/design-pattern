package abstractfactory

// --8<-- [start:light]
// Light 是浅色主题工厂，只生产白底控件。
type Light struct{}

func (Light) NewButton(label string) Button     { return lightButton{label} }
func (Light) NewCheckbox(label string) Checkbox { return lightCheckbox{label} }

// 具体产品不导出，原因见正文提示。
type lightButton struct{ label string }
type lightCheckbox struct{ label string }

func (b lightButton) Render() string   { return "[白底按钮 " + b.label + "]" }
func (c lightCheckbox) Render() string { return "[白底勾选框 " + c.label + "]" }

// --8<-- [end:light]
