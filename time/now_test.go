package time

import (
	"fmt"
	"testing"
	"time"
)

func TestNowLocal(t *testing.T) {
	//fmt.Println(NowYear())
	//fmt.Println(NowMonth())
	//fmt.Println(NowDay())
	//fmt.Println(NowClock())
	//
	//err := SetLoc("America/Dawson")
	//fmt.Println(err)
	//fmt.Println(NowYear())
	//fmt.Println(NowMonth())
	//fmt.Println(NowDay())
	//fmt.Println(NowClock())
}

func TestLocalAdd(t *testing.T) {
	//fmt.Println(NowAddDayLayout(1, time.RFC3339))
	//fmt.Println(NowAddYearLayout(1, time.RFC3339))
	//fmt.Println(NowAddMonthLayout(1, time.RFC3339))
	d, err := time.Parse("2006-01-05 15:04:05", "2017-01-12 15:03:50")
	fmt.Println(err)
	p := d
	ts := d.Truncate(10 * time.Second)
	tp := p.Round(10 * time.Second)
	fmt.Println(ts, tp)
}
