package time

import (
	"fmt"
	"testing"
)

func TestTimestamps(t *testing.T) {
	err := SetLoc("UTC")
	fmt.Println(err)
	a, err := Timestamps("2017-01-02", "2006-01-02")
	fmt.Println(a, err)
	fmt.Println(UnixTime())
}
