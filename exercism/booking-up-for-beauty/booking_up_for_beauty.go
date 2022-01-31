package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	schedule, _ := time.Parse(layout, date)
	return schedule
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	schedule, _ := time.Parse(layout, date)
	currentTime := time.Now().UTC()
	return currentTime.After(schedule)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	schedule, _ := time.Parse(layout, date)
	return schedule.Hour() >= 12 && schedule.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	schedule, _ := time.Parse(layout, date)

	statement := "You have an appointment on "
	statement += schedule.Weekday().String() + ", " + schedule.Month().String() + " "
	statement += fmt.Sprint(schedule.Day())
	statement += ", " + fmt.Sprint(schedule.Year()) + ", at " + fmt.Sprint(schedule.Hour()) + ":" + fmt.Sprint(schedule.Minute()) + "."
	return statement
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	anniversary := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
