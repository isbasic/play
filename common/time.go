package common

import (
	"fmt"
	"time"
)

func FormatTime(t time.Time) string {
	resTmpl := "%04d%02d%02d%02d%02d%02d%09d" // yyyyMMddhhmmssnanosecond
	y, m, d := t.Date()
	h, mn, s, nano := t.Hour(), t.Minute(), t.Second(), t.Nanosecond()
	return fmt.Sprintf(resTmpl, y, m, d, h, mn, s, nano)
}
