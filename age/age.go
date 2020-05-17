package age

import (
	"github.com/qq1060656096/tools/date"
	"time"
)

// At 获取特定时间之间的年龄
func At(birthDate time.Time, nowDate time.Time) int {
	age := nowDate.Year() - birthDate.Year()
	if nowDate.Month() <= birthDate.Month() && nowDate.Day() < birthDate.Day() {
		age -= 1
	}
	return age
}

// Get 根据出生日期获取年龄
func Get(birthdayTime time.Time) int {
	return At(birthdayTime, time.Now())
}

// IsLeapYear 检测日期是否是闰年
func IsLeapYear(t time.Time) bool {
	return date.IsLeapYear(t)
}
