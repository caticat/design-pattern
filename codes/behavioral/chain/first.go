package chain

// --8<-- [start:first]
// TryFunc 是一个环节：能处理就返回结果和 true，否则返回 false。
type TryFunc func(r Request) (string, bool)

// First 把若干环节按顺序串成一条链，第一个处理成功的说了算。
func First(handlers ...TryFunc) func(Request) string {
	return func(r Request) string {
		for _, h := range handlers {
			if result, ok := h(r); ok {
				return result
			}
		}
		return "无人处理"
	}
}

// --8<-- [end:first]
