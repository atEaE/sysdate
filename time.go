package systime

import (
	"os"
	"time"
)

const (
	// SysDateKey is the key for the system date in the context.
	SysDateKey = "GO_SYS_DATE"
	// SysTimeKey is the key for the system time in the context.
	SysDateFormat = "2006-01-02"
)

// IsRepeat
// If IsRepeat is true, the date is not updated after 24:00 and the date set in GO_SYS_DATE is repeated
var Repeat bool

var (
	sysDate *time.Time

	initDateUTC time.Time
	elapsedDay  int
)

func init() {
	now := time.Now()
	initDateUTC = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	elapsedDay = 0

	v := os.Getenv(SysDateKey)
	if v != "" {
		t, err := time.Parse(SysDateFormat, v)
		if err == nil {
			sysDate = &t
		}
	}
}

// Now returns the current local time.
func Now() time.Time {
	now := time.Now()
	if sysDate != nil {
		sysnow := time.Date(sysDate.Year(), sysDate.Month(), sysDate.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)
		if !Repeat {
			elapsedDay = calcElapsedDay(initDateUTC, now.UTC())
			return sysnow.Add(time.Duration(elapsedDay*24) * time.Hour)
		}
		return sysnow
	}
	return now
}

// calcElapsedDay calculates the elapsed days from the initial date.
func calcElapsedDay(initUTC, nowUTC time.Time) int {
	elapsedDay = int(nowUTC.Sub(initUTC).Hours() / 24)
	return elapsedDay
}
