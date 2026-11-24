package visitor_test

import (
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/visitor"
)

var gallery = []visitor.Exhibit{
	visitor.Sculpture{Title: "海风", WeightKG: 12},
	visitor.Painting{Title: "晨光", WidthCM: 30, HeightCM: 40},
	visitor.Sculpture{Title: "远山", WeightKG: 30},
}

// WeightVisitor 只累加雕塑的重量。
type WeightVisitor struct{ TotalKG int }

func (v *WeightVisitor) VisitPainting(visitor.Painting) {} // 画作不计重
func (v *WeightVisitor) VisitSculpture(s visitor.Sculpture) {
	v.TotalKG += s.WeightKG
}

// 练习 1：新增「称重」操作
func TestExercise1_Weight(t *testing.T) {
	w := &WeightVisitor{}
	for _, e := range gallery {
		e.Accept(w)
	}
	if w.TotalKG != 42 {
		t.Errorf("got %d, want 42", w.TotalKG)
	}
}

// Counter 用值接收者：每次改的都是副本。
type Counter struct{ N int }

func (c Counter) VisitPainting(visitor.Painting)   { c.N++ }
func (c Counter) VisitSculpture(visitor.Sculpture) { c.N++ }

// PtrCounter 用指针接收者：改的是同一个计数器。
type PtrCounter struct{ N int }

func (c *PtrCounter) VisitPainting(visitor.Painting)   { c.N++ }
func (c *PtrCounter) VisitSculpture(visitor.Sculpture) { c.N++ }

// 练习 2：值接收者的访问者
func TestExercise2_Receiver(t *testing.T) {
	c := Counter{}
	for _, e := range gallery {
		e.Accept(c)
	}
	if c.N != 0 {
		t.Errorf("值接收者：got %d, want 0", c.N)
	}

	p := &PtrCounter{}
	for _, e := range gallery {
		e.Accept(p)
	}
	if p.N != 3 {
		t.Errorf("指针接收者：got %d, want 3", p.N)
	}
}
