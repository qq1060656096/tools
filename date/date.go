package date

import (
	"strings"
	"time"
)

const (
	DefaultDateTime = "Y-m-d h:i:s"
	DefaultDate     = "Y-m-d"
)

type formatReplace struct {
	from string
	to   string
}

var formatFromTo = [...]formatReplace{
	{"d", "02"},                        // Day:    Day of the month, 2 digits with leading zeros. Eg: 01 to 31.
	{"D", "Mon"},                       // Day:    A textual representation of a day, three letters. Eg: Mon through Sun.
	{"w", "Monday"},                    // Day:    Numeric representation of the day of the week. Eg: 0 (for Sunday) through 6 (for Saturday).
	{"N", "Monday"},                    // Day:    ISO-8601 numeric representation of the day of the week. Eg: 1 (for Monday) through 7 (for Sunday).
	{"S", "02"},                        // Day:    English ordinal suffix for the day of the month, 2 characters. Eg: st, nd, rd or th. Works well with j.
	{"l", "Monday"},                    // Day:    A full textual representation of the day of the week. Eg: Sunday through Saturday.
	{"z", ""},                          // Day:    The day of the year (starting from 0). Eg: 0 through 365.
	{"W", ""},                          // Week:   ISO-8601 week number of year, weeks starting on Monday. Eg: 42 (the 42nd week in the year).
	{"F", "January"},                   // Month:  A full textual representation of a month, such as January or March. Eg: January through December.
	{"m", "01"},                        // Month:  Numeric representation of a month, with leading zeros. Eg: 01 through 12.
	{"M", "Jan"},                       // Month:  A short textual representation of a month, three letters. Eg: Jan through Dec.
	{"n", "1"},                         // Month:  Numeric representation of a month, without leading zeros. Eg: 1 through 12.
	{"t", ""},                          // Month:  Number of days in the given month. Eg: 28 through 31.
	{"Y", "2006"},                      // Year:   A full numeric representation of a year, 4 digits. Eg: 1999 or 2003.
	{"y", "06"},                        // Year:   A two digit representation of a year. Eg: 99 or 03.
	{"a", "pm"},                        // Time:   Lowercase Ante meridiem and Post meridiem. Eg: am or pm.
	{"A", "PM"},                        // Time:   Uppercase Ante meridiem and Post meridiem. Eg: AM or PM.
	{"g", "3"},                         // Time:   12-hour format of an hour without leading zeros. Eg: 1 through 12.
	{"h", "03"},                        // Time:   12-hour format of an hour with leading zeros. Eg: 01 through 12.
	{"H", "15"},                        // Time:   24-hour format of an hour with leading zeros. Eg: 00 through 23.
	{"i", "04"},                        // Time:   Minutes with leading zeros. Eg: 00 to 59.
	{"s", "05"},                        // Time:   Seconds with leading zeros. Eg: 00 through 59.
	{"U", ""},                          // Time:   Seconds since the Unix Epoch (January 1 1970 00:00:00 GMT).
	{"O", "-0700"},                     // Zone:   Difference to Greenwich time (GMT) in hours. Eg: +0200.
	{"P", "-07:00"},                    // Zone:   Difference to Greenwich time (GMT) with colon between hours and minutes. Eg: +02:00.
	{"T", "MST"},                       // Zone:   Timezone abbreviation. Eg: UTC, EST, MDT ...
	{"c", "2006-01-02T15:04:05-07:00"}, // Format: ISO 8601 date. Eg: 2004-02-12T15:19:21+00:00.
	{"r", "Mon, 02 Jan 06 15:04 MST"},  // Format: RFC 2822 formatted date. Eg: Thu, 21 Dec 2000 16:01:07 +0200.
}

func Replace(format string) string {
	for _, v := range formatFromTo {
		format = strings.Replace(format, v.from, v.to, -1)
	}
	return format
}

func Format(format string, t time.Time) string {
	toFormat := Replace(format)
	return t.Format(toFormat)
}

func Parse(format string, date string) (t time.Time, err error) {
	toFormat := Replace(format)
	return time.Parse(toFormat, date)
}

func Now(format string) string {
	return Format(format, time.Now())
}


// IsLeapYear 检测日期是否是闰年
func IsLeapYear(t time.Time) bool {
	y := t.Year()
	switch {
	case y%400 == 0:
		return true
	case y%100 == 0:
		return true
	case y%4 == 0:
		return true
	}
	return false
}

func Days(t time.Time) int {
	t = t.AddDate(0, 1, -t.Day())
	return t.Day()
}