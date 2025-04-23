package util

import (
	"time"
)

func FormatJSTTime(t time.Time) string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return t.In(jst).Format("2006-01-02 15:04:05")
}
