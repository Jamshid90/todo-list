package time

import "time"

const (
	TimeFormatDate     = "2006-01-02"
	TimeFormatDateTime = "2006-01-02 15:04:05"
)

// date
func ToDateFormat(t time.Time) string {
	return t.Format(TimeFormatDate)
}

func ParseToDate(s string) (time.Time, error) {
	return time.Parse(TimeFormatDate, s)
}

func ToDateTimeFormat(t time.Time) string {
	return t.Format(TimeFormatDateTime)
}

func ParseToDateTime(s string) (time.Time, error) {
	return time.Parse(TimeFormatDateTime, s)
}

func ToDateTimeFormatRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ParseToDateTimeRFC3339(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}
