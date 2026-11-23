package iterator

import "iter"

// --8<-- [start:filter]
// Filter 包装一个迭代器，只放行满足条件的元素；不管它背后是链表、切片还是文件。
func Filter[T any](seq iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if keep(v) && !yield(v) {
				return
			}
		}
	}
}

// --8<-- [end:filter]
