package main

import (
	"fmt"
	"time"

	. "github.com/valdirmendesdev/time-utils/time"
)

func main() {
	
	days := GetWeekDays(2024, 14)
	for _, day := range days {
		fmt.Println(day.Format(time.DateOnly))
	}
}