package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[1] = time.Unix(1747920000, 0)
	// fmt.Print(dates[1])

	// var dates[3]time.Time = [3]time.Time{time.Unix(1, 0), time.Unix(1747920000, 0), time.Unix(1447920001, 0)}
	// fmt.Print(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(1, 0),
		time.Unix(1747920000, 0),
		time.Unix(1447920001, 1), // 쉼표가 반드시 필요
	}
	fmt.Print(dates[0], dates[1], dates[2])
}
