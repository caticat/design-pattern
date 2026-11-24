package iterator_test

import (
	"iter"
	"slices"
	"strings"
	"testing"

	"github.com/caticat/design-pattern/codes/behavioral/iterator"
)

// Take 只放行前 n 个元素。
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n <= 0 {
			return
		}
		count := 0
		for v := range seq {
			if !yield(v) {
				return // 调用方 break 了
			}
			count++
			if count == n {
				return // 数够了
			}
		}
	}
}

// 练习 1：再写一个组合器 Take
func TestExercise1_Take(t *testing.T) {
	p := &iterator.Playlist{}
	for _, title := range []string{"晨光", "雨夜", "海风", "雨巷", "雨后"} {
		p.Add(title)
	}
	rainy := func(s string) bool { return strings.HasPrefix(s, "雨") }

	got := slices.Collect(Take(iterator.Filter(p.All(), rainy), 2))
	if want := []string{"雨夜", "雨巷"}; !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
