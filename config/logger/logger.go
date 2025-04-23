package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type ctxKeyLogger struct{}

var (
	Log *logrus.Entry
)

func Init() {
	base := logrus.New()
	base.SetOutput(os.Stdout)
	base.SetLevel(logrus.InfoLevel)
	base.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	Log = logrus.NewEntry(base)
}

// Context に logger を埋め込む
func ToContext(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, ctxKeyLogger{}, entry)
}

// Context から logger を取り出す（なければ default）
func FromContext(ctx context.Context) *logrus.Entry {
	if logger, ok := ctx.Value(ctxKeyLogger{}).(*logrus.Entry); ok {
		return logger
	}
	return Log
}
