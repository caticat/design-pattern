package abstractfactory

// --8<-- [start:bad]
// 反例：每个控件各自判断主题、各自创建。
func LoginFormBad(dark bool) string {
	var c Checkbox = lightCheckbox{"记住我"}
	if dark {
		c = darkCheckbox{"记住我"}
	}
	var b Button = lightButton{"登录"}
	if !dark { // 条件写反了：深色页面上冒出一个白底按钮
		b = darkButton{"登录"}
	}
	return c.Render() + " " + b.Render()
}

// --8<-- [end:bad]
