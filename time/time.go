package time

import (
	"log"
	"time"
)

var (
	defaultLoc        *time.Location
)

func init() {
	var err error
	defaultLoc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
}

//Date 时间戳转日期
func Date(timestamps int64, layout string) string {
	return time.Unix(timestamps, 0).Format(layout)
}

//Timestamps 日期转时间戳
func Timestamps(date string, layout string) (int64, error) {
	timeIns, err := time.Parse(layout, date)
	if err != nil {
		return 0, err
	}
	return timeIns.Unix(), nil
}
