// Package abstractfactory 演示抽象工厂模式：一个工厂负责创建一整族配套的对象，
// 换工厂就整族一起换，不会混搭。
package abstractfactory

// --8<-- [start:product]
// Button 和 Checkbox 是两种产品，每套主题都要提供。
type Button interface {
	Render() string
}

type Checkbox interface {
	Render() string
}

// --8<-- [end:product]

// --8<-- [start:factory]
// Theme 是抽象工厂：一套主题能创建的所有控件。
type Theme interface {
	NewButton(label string) Button
	NewCheckbox(label string) Checkbox
}

// --8<-- [end:factory]

// --8<-- [start:client]
// LoginForm 只认 Theme，不知道是哪套主题，控件自然配套。
func LoginForm(t Theme) string {
	return t.NewCheckbox("记住我").Render() + " " + t.NewButton("登录").Render()
}

// --8<-- [end:client]
