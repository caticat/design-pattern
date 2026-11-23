package visitor_test

import (
	"fmt"

	"github.com/caticat/design-pattern/codes/behavioral/visitor"
)

// --8<-- [start:example]
func Example() {
	exhibits := []visitor.Exhibit{
		visitor.Painting{Title: "晨光", WidthCM: 30, HeightCM: 40},
		visitor.Sculpture{Title: "海风", WeightKG: 12},
	}
	labels := &visitor.LabelVisitor{}
	packing := &visitor.PackingVisitor{}
	for _, e := range exhibits {
		e.Accept(labels)
		e.Accept(packing)
	}
	for _, label := range labels.Labels {
		fmt.Println(label)
	}
	fmt.Printf("运输：%d 个画箱，%d 个木箱\n", packing.FlatBoxes, packing.Crates)
	// Output:
	// 画作：晨光（30×40 厘米）
	// 雕塑：海风（12 千克）
	// 运输：1 个画箱，1 个木箱
}

// --8<-- [end:example]
