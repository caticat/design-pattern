package abstractfactory_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/creational/abstractfactory"
)

// --8<-- [start:example]
func Example() {
	// 主题通常来自用户设置；换主题只换这一个值
	for _, t := range []abstractfactory.Theme{abstractfactory.Light{}, abstractfactory.Dark{}} {
		fmt.Println(abstractfactory.LoginForm(t))
	}
	// Output:
	// [白底勾选框 记住我] [白底按钮 登录]
	// [黑底勾选框 记住我] [黑底按钮 登录]
}

// --8<-- [end:example]
