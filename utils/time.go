package dramautils

import "time"

// ITime - Time
type ITime interface {
	// Now - get now time
	Now() time.Time
}

// Time - default Time
type Time struct {
}

// Now - get now time
func (t Time) Now() time.Time {
	return time.Now()
}

var gTime ITime
var gUTCLocal *time.Location

// FormatNow - format time
func FormatNow(t ITime) string {
	return t.Now().In(gUTCLocal).Format("2006-01-02_15:04:05")
}

func GetCurTimestamp() int64 {
	return time.Now().Unix()
}

func init() {
	gTime = &Time{}

	l, err := time.LoadLocation("")
	if err == nil {
		gUTCLocal = l
	}
}
