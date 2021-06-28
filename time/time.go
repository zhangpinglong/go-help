package time

import (
	"time"
)

var (
	loc *time.Location
)

func SetLoc(locName string) error {
	var err error
	loc, err = time.LoadLocation(locName)
	if err != nil {
		return err
	}
	return nil
}

//Date 时间戳转日期
func Date(timestamps int64, layout string) string {
	if loc == nil {
		return time.Unix(timestamps, 0).Format(layout)
	}
	return time.Unix(timestamps, 0).In(loc).Format(layout)
}

//Timestamps 日期转时间戳
func Timestamps(date string, layout string) (int64, error) {
	timeIns, err := time.ParseInLocation(layout, date, loc)
	if err != nil {
		return 0, err
	}

	return timeIns.Unix(), nil
}

//UnixTime 时间戳
func UnixTime() int64 {
	return time.Now().Unix()
}
