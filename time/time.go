package time

import "time"

const timeFormat = "2006-01-02 15:04:05"

func SubDays(t1, t2 time.Time) (day int64) {
	day = int64(t1.Sub(t2).Hours() / 24)
	return
}

// s: string type value like "2020-09-08 13:47:59"
func ParseTime(s string) (time.Time, error) {
	res, err := time.Parse(timeFormat, s)
	return res, err
}
