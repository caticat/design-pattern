package visitor

import "fmt"

// --8<-- [start:bad]
// 反例：每加一种展品，都要找到并更新各项操作里的类型分支。
func badLabels(exhibits []Exhibit) []string {
	var labels []string
	for _, e := range exhibits {
		switch x := e.(type) {
		case Painting:
			labels = append(labels, fmt.Sprintf("画作：%s（%d×%d 厘米）", x.Title, x.WidthCM, x.HeightCM))
		case Sculpture:
			labels = append(labels, fmt.Sprintf("雕塑：%s（%d 千克）", x.Title, x.WeightKG))
		}
	}
	return labels
}

func badPacking(exhibits []Exhibit) (flatBoxes, crates int) {
	for _, e := range exhibits {
		switch e.(type) {
		case Painting:
			flatBoxes++
		case Sculpture:
			crates++
		}
	}
	return
}

// --8<-- [end:bad]
