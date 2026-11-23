package templatemethod_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/templatemethod"
)

// --8<-- [start:example]
// Text 是新加的格式：写在包外，只实现两个步骤，流程由 Export 保证。
type Text struct{}

func (Text) Header() string { return "== 成绩单 ==" }
func (Text) Row(r templatemethod.Record) string {
	return fmt.Sprintf("%s：%d 分", r.Name, r.Score)
}

func Example() {
	rows := []templatemethod.Record{{Name: "张三", Score: 95}, {Name: "李四", Score: 82}}

	fmt.Print(templatemethod.Export(templatemethod.Markdown{}, rows))
	fmt.Print(templatemethod.Export(Text{}, rows))
	// Output:
	// | 姓名 | 分数 |
	// |---|---|
	// | 张三 | 95 |
	// | 李四 | 82 |
	// 共 2 行
	// == 成绩单 ==
	// 张三：95 分
	// 李四：82 分
}

// --8<-- [end:example]
