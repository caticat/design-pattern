package iterator

import (
	"slices"
	"strings"
	"testing"
)

func newPlaylist(titles ...string) *Playlist {
	p := &Playlist{}
	for _, t := range titles {
		p.Add(t)
	}
	return p
}

func TestClassicAndSeqAgree(t *testing.T) {
	p := newPlaylist("a", "b", "c")
	var classic []string
	for it := p.Iterator(); it.HasNext(); {
		classic = append(classic, it.Next())
	}
	seq := slices.Collect(p.All())
	want := []string{"a", "b", "c"}
	if !slices.Equal(classic, want) || !slices.Equal(seq, want) {
		t.Fatalf("classic=%v seq=%v, want %v", classic, seq, want)
	}
}

func TestEmpty(t *testing.T) {
	p := &Playlist{}
	if p.Iterator().HasNext() {
		t.Fatal("空列表不应有下一首")
	}
	if got := slices.Collect(p.All()); len(got) != 0 {
		t.Fatalf("got %v", got)
	}
}

func TestBreakStops(t *testing.T) {
	p := newPlaylist("a", "b", "c")
	n := 0
	for range p.All() {
		n++
		break
	}
	if n != 1 {
		t.Fatalf("n=%d", n)
	}
}

func TestFilter(t *testing.T) {
	p := newPlaylist("晨光", "雨夜", "海风", "雨巷")
	got := slices.Collect(Filter(p.All(), func(s string) bool { return strings.HasPrefix(s, "雨") }))
	if !slices.Equal(got, []string{"雨夜", "雨巷"}) {
		t.Fatalf("got %v", got)
	}
}

func TestBadRuns(t *testing.T) {
	bp := &badPlaylist{Head: &songNode{Title: "a", Next: &songNode{Title: "b"}}}
	printTitles(bp) // 反例只确认能跑
}
