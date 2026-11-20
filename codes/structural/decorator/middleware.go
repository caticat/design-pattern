package decorator

import (
	"fmt"
	"io"
	"net/http"
)

// --8<-- [start:middleware]
// Middleware 是 HTTP 装饰器：输入一个 Handler，返回包装后的 Handler。
type Middleware func(http.Handler) http.Handler

// Logging 在请求前后打印日志。
func Logging(out io.Writer) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(out, "-> %s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
			fmt.Fprintf(out, "<- %s\n", r.URL.Path)
		})
	}
}

// RequireToken 校验请求头，缺少 token 时直接拒绝，不再调用 next。
func RequireToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Token") == "" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// --8<-- [end:middleware]

// --8<-- [start:chain]
// Chain 按顺序套上中间件：ms[0] 在最外层，最先执行。
func Chain(h http.Handler, ms ...Middleware) http.Handler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i](h)
	}
	return h
}

// --8<-- [end:chain]
