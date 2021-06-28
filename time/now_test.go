package time

import (
	"fmt"
	"testing"
	"time"
)

func TestNowLocal(t *testing.T) {
	fmt.Println(NowYear())
	fmt.Println(NowMonth())
	fmt.Println(NowDay())
	fmt.Println(NowClock())

	err := SetLoc("America/Dawson")
	fmt.Println(err)
	fmt.Println(NowYear())
	fmt.Println(NowMonth())
	fmt.Println(NowDay())
	fmt.Println(NowClock())
}

func TestLocalAdd(t *testing.T) {
	fmt.Println(NowAddDayLayout(1, time.RFC3339))
	fmt.Println(NowAddYearLayout(1, time.RFC3339))
	fmt.Println(NowAddMonthLayout(1, time.RFC3339))
}
