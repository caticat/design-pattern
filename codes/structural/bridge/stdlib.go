package bridge

import (
	"io"
	"log/slog"
)

// --8<-- [start:slog]
// logWith 只用 Logger 的方法，不关心日志最终是什么格式、写到哪里。
func logWith(h slog.Handler) {
	logger := slog.New(h) // 抽象：Info、Warn、With……
	logger.Info("用户登录", "user", "tom")
}

// logBoth 同一段日志代码，接上两种 Handler（实现）。
func logBoth(w io.Writer) {
	opts := &slog.HandlerOptions{ReplaceAttr: dropTime} // 去掉时间，方便对比
	logWith(slog.NewTextHandler(w, opts))
	logWith(slog.NewJSONHandler(w, opts))
}

// 输出：
//   level=INFO msg=用户登录 user=tom
//   {"level":"INFO","msg":"用户登录","user":"tom"}
// --8<-- [end:slog]

func dropTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}
