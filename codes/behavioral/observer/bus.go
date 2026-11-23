package observer

import "slices"

// --8<-- [start:bus]
// Bus 用函数当观察者；On 返回一个取消订阅的函数。
type Bus struct {
	subs []*sub
}

type sub struct{ fn func(Order) }

func (b *Bus) On(fn func(Order)) (cancel func()) {
	s := &sub{fn}
	b.subs = append(b.subs, s)
	return func() {
		b.subs = slices.DeleteFunc(b.subs, func(x *sub) bool { return x == s })
	}
}

func (b *Bus) Emit(o Order) {
	for _, s := range b.subs {
		s.fn(o)
	}
}

// --8<-- [end:bus]
