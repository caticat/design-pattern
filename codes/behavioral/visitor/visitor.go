package visitor

import "fmt"

// --8<-- [start:contracts]
// Exhibit 是允许访问者处理的展品。
type Exhibit interface {
	Accept(Visitor)
}

// Visitor 为每种展品声明一个处理方法。
type Visitor interface {
	VisitPainting(Painting)
	VisitSculpture(Sculpture)
}

// --8<-- [end:contracts]

// --8<-- [start:exhibits]
// Painting 是用厘米记录尺寸的画作。
type Painting struct {
	Title             string
	WidthCM, HeightCM int
}

func (p Painting) Accept(v Visitor) { v.VisitPainting(p) }

// Sculpture 是用千克记录重量的雕塑。
type Sculpture struct {
	Title    string
	WeightKG int
}

func (s Sculpture) Accept(v Visitor) { v.VisitSculpture(s) }

// --8<-- [end:exhibits]

// --8<-- [start:operations]
// LabelVisitor 按访问顺序收集展签。
type LabelVisitor struct{ Labels []string }

func (v *LabelVisitor) VisitPainting(p Painting) {
	v.Labels = append(v.Labels, fmt.Sprintf("画作：%s（%d×%d 厘米）", p.Title, p.WidthCM, p.HeightCM))
}

func (v *LabelVisitor) VisitSculpture(s Sculpture) {
	v.Labels = append(v.Labels, fmt.Sprintf("雕塑：%s（%d 千克）", s.Title, s.WeightKG))
}

// PackingVisitor 统计运输所需的画箱和木箱。
type PackingVisitor struct{ FlatBoxes, Crates int }

func (v *PackingVisitor) VisitPainting(Painting)   { v.FlatBoxes++ }
func (v *PackingVisitor) VisitSculpture(Sculpture) { v.Crates++ }

// --8<-- [end:operations]
