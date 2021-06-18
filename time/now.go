package time

import "time"

//NowLocalYear 本地时间的年份
func NowLocalYear() int {
	return time.Now().Year()
}

//NowLocalMonth 本地时间的月份
func NowLocalMonth() time.Month {
	return time.Now().Month()
}

//NowLocalDay 本地时间的日期
func NowLocalDay() int {
	return time.Now().Day()
}

//NowLocalTime 本地时间的时间戳
func NowLocalTime() int64 {
	return time.Now().Unix()
}

//NowLocalTime 本地时间的增加年月日 返回layout格式的日期字符串
func NowLocalAdd(years int, months int, days int, layout string) string {
	return time.Now().AddDate(years, months, days).Format(layout)
}

//NowAddLocalDay 当前时间增加天数 返回layout格式的日期字符串
func NowAddLocalDay(days int, layout string) string {
	return NowLocalAdd(0, 0, days, layout)
}

//NowAddLocalMonth 当前时间增加月份 返回layout格式的日期字符串
func NowAddLocalMonth(months int, layout string) string {
	return NowLocalAdd(0, months, 0, layout)
}

//NowAddLocalYear  当前时间增加年数 返回layout格式的日期字符串
func NowAddLocalYear(years int, layout string) string {
	return NowLocalAdd(years, 0, 0, layout)
}

//NowTimeByZone 根据时区返回时间戳
func NowTimeByZone(loc *time.Location) int64 {
	return time.Now().In(loc).Unix()
}
