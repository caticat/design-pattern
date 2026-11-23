package visitor

import (
	"slices"
	"testing"
)

func TestLabelVisitorPreservesVisitOrder(t *testing.T) {
	exhibits := []Exhibit{
		Sculpture{Title: "海风", WeightKG: 12},
		Painting{Title: "晨光", WidthCM: 30, HeightCM: 40},
		Painting{Title: "雨夜", WidthCM: 20, HeightCM: 25},
	}
	v := &LabelVisitor{}
	for _, exhibit := range exhibits {
		exhibit.Accept(v)
	}
	want := []string{"雕塑：海风（12 千克）", "画作：晨光（30×40 厘米）", "画作：雨夜（20×25 厘米）"}
	if !slices.Equal(v.Labels, want) {
		t.Errorf("展签 = %q，期望 %q", v.Labels, want)
	}
}

func TestPackingVisitorCountsSeparately(t *testing.T) {
	var painting Exhibit = Painting{}
	var sculpture Exhibit = Sculpture{}
	v := &PackingVisitor{}
	painting.Accept(v)
	sculpture.Accept(v)
	painting.Accept(v)
	if v.FlatBoxes != 2 || v.Crates != 1 {
		t.Errorf("画箱 = %d，木箱 = %d；期望 2、1", v.FlatBoxes, v.Crates)
	}
}
