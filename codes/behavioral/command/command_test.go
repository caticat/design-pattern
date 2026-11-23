package command

import (
	"slices"
	"testing"
)

func TestQueue(t *testing.T) {
	var q Queue
	var got []int
	q.Add(Func(func() {
		got = append(got, 1)
		q.Add(Func(func() { got = append(got, 3) }))
	}))
	q.Add(Func(func() { got = append(got, 2) }))
	if len(got) != 0 {
		t.Fatal("入队不应立即执行")
	}
	q.Run()
	if !slices.Equal(got, []int{1, 2}) {
		t.Fatalf("首批命令乱序或多执行：%v", got)
	}
	q.Run()
	q.Run()
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("后续命令应只执行一次：%v", got)
	}
}

func TestLightOn(t *testing.T) {
	lamp := &Lamp{}
	LightOn{Lamp: lamp}.Execute()
	if !lamp.On {
		t.Fatal("台灯没有打开")
	}
}

func TestPress(t *testing.T) {
	lamp, fan := &Lamp{}, &Fan{}
	press("风扇", lamp, fan)
	if lamp.On || !fan.Running {
		t.Fatal("按风扇按钮应只打开风扇")
	}
}
