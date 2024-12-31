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

var sysDate *time.Time

func init() {
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
		return time.Date(sysDate.Year(), sysDate.Month(), sysDate.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	}
	return now
}
