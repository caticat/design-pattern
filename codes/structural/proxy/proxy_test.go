package proxy

import "testing"

var (
	_ Video = (*HDVideo)(nil)
	_ Video = (*VideoProxy)(nil)
)

func TestNotLoadedUntilPlay(t *testing.T) {
	p := NewVideoProxy("海底世界", false)
	if p.real != nil {
		t.Fatal("创建代理时不应加载")
	}
	if _, err := p.Play(User{Name: "小明"}); err != nil {
		t.Fatal(err)
	}
	if p.real == nil {
		t.Fatal("播放后应已加载")
	}
}

func TestDeniedDoesNotLoad(t *testing.T) {
	p := NewVideoProxy("星际旅行", true)
	if _, err := p.Play(User{Name: "小明"}); err == nil {
		t.Fatal("非会员应被拒绝")
	}
	if p.real != nil {
		t.Fatal("被拒绝时不应加载")
	}
}

func TestLoadsOnce(t *testing.T) {
	p := NewVideoProxy("星际旅行", true)
	vip := User{Name: "小红", VIP: true}
	p.Play(vip)
	first := p.real
	p.Play(vip)
	if p.real != first {
		t.Fatal("第二次播放应复用已加载的视频")
	}
}

func TestBad(t *testing.T) {
	if got := openHomeBad([]string{"a", "b", "c"}); len(got) != 3 {
		t.Fatalf("反例应一次加载全部，got %d", len(got))
	}
	if _, err := playBad(LoadHDVideo("a"), User{Name: "小明"}, true); err == nil {
		t.Fatal("非会员应被拒绝")
	}
}
