package time

import "time"

//NowYear 当前时间的年
func NowYear() int {
	if loc == nil {
		return time.Now().Year()
	}
	return time.Now().In(loc).Year()
}

//NowMonth 当前时间的月
func NowMonth() time.Month {
	if loc == nil {
		return time.Now().Month()
	}
	return time.Now().In(loc).Month()
}

//NowDay 当前时间的日期
func NowDay() int {
	if loc == nil {
		return time.Now().Day()
	}
	return time.Now().In(loc).Day()
}

//NowClock 当前时间时分秒
func NowClock() (hour int, min int, sec int) {
	if loc == nil {
		return time.Now().Clock()
	}
	return time.Now().In(loc).Clock()
}

// NowAddLayout  当前时间的增加年月日 返回layout格式的日期字符串
func NowAddLayout(years int, months int, days int, layout string) string {
	if loc == nil {
		return time.Now().AddDate(years, months, days).Format(layout)
	}
	return time.Now().In(loc).AddDate(years, months, days).Format(layout)
}

//NowAddDayLayout 当前时间增加天数 返回layout格式的日期字符串
func NowAddDayLayout(days int, layout string) string {
	return NowAddLayout(0, 0, days, layout)
}

//NowAddMonthLayout 当前时间增加月份 返回layout格式的日期字符串
func NowAddMonthLayout(months int, layout string) string {
	return NowAddLayout(0, months, 0, layout)
}

//NowAddYearLayout  当前时间增加年数 返回layout格式的日期字符串
func NowAddYearLayout(years int, layout string) string {
	return NowAddLayout(years, 0, 0, layout)
}

//NowAddUnix 当前时间加上年月日 返回时间戳
func NowAddUnix(years int, months int, days int) int64 {
	return time.Now().In(loc).AddDate(years, months, days).Unix()
}

//NowAddDayUnix 当前时间增加天数 返回时间戳
func NowAddDayUnix(days int) int64 {
	return NowAddUnix(0, 0, days)
}

//NowAddMonthUnix 当前时间增加月份 返回时间戳
func NowAddMonthUnix(months int) int64 {
	return NowAddUnix(0, months, 0)
}

//NowAddYearUnix  当前时间增加年数 返回时间戳
func NowAddYearUnix(years int) int64 {
	return NowAddUnix(years, 0, 0)
}

//NowAddDurationUnix 添加
func NowAddDurationUnix(s string) (int64, error) {
	dur, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}
	return time.Now().Add(dur).Unix(), err
}
