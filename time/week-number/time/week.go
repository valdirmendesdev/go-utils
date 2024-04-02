package time

import (
	"time"
)

const (
	daysInAWeek = 7
)

func GetWeekDays(year, weekNumber int) (result []time.Time) {
	
  //Defines the year first date
  startsDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

	//Calculates the first day of the week
	//Here, you can define the first day of the week you want by default,
	//I am considering Sunday
  startsDate = startsDate.AddDate(0, 0, -int(startsDate.Weekday())+int(time.Sunday))

	//Jumps to the first day of the correct week
  startsDate = startsDate.AddDate(0, 0, (weekNumber-1)*daysInAWeek)

  // Generates all the days of the week
  for i := 0; i < 7; i++ {
		result = append(result, startsDate.AddDate(0,0,i))
  }
	return
}