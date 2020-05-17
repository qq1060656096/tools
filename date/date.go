package date

import (
	"strings"
	"time"
)

const (
	DefaultDateTime = "YYYY-MM-DD hh:mm:ss"
	DefaultDate = "YYYY-MM-DD"
)

type formatReplace struct {
	from string
	to string
}
var formatFromTo = [...]formatReplace{
	{"hh", "15"},
	{"mm", "04"},
	{"ss", "05"},
	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},
	{"M", "1"},
	{"pm", "PM"},
	{"ZZZZ", "-0700"},
	{"ZZZ", "MST"},
	{"ZZ", "Z07:00"},
	{"YYYY", "2006"},
	{"YY", "06"},
	{"DDDD", "Monday"},
	{"DDD", "Mon"},
	{"DD", "02"},
	{"D", "2"},
}


func replace(format string) string {
	for _, v := range formatFromTo {
		format = strings.Replace(format, v.from, v.to, -1)
	}
	return format
}

func Format(format string, t time.Time) string {
	toFormat := replace(format)
	return t.Format(toFormat)
}

func Parse(format string, date string) (t time.Time, err error) {
	toFormat := replace(format)
	return time.Parse(toFormat, date)
}

func Now(format string) string {
	return Format(format, time.Now())
}