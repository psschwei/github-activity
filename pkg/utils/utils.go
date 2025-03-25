package utils

import (
	"os/user"
	"time"
)

func GetCurrentUsername() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}
	return user.Username

}

func GetDefaultStartDate() string {
	return time.Now().In(time.Local).AddDate(0, 0, -7).Format("2006-01-02")
}

func GetDefaultEndDate() string {
	return time.Now().In(time.Local).AddDate(0, 0, 1).Format("2006-01-02")
}

func GetTodayDates() (string, string) {
	now := time.Now()
	start := now.In(time.Local).Format("2006-01-02")
	end := now.In(time.Local).AddDate(0,now.Hour(), 0).Format("2006-01-02")
	return start, end
}

func GetThisWeekDates() (string, string) {
	now := time.Now()
	weekday := time.Duration(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	thisMonday := now.Add(-1 * (weekday - 1) * 24 * time.Hour)
	thisFriday := thisMonday.Add(5 * 24 * time.Hour)

    thisMonday = time.Date(thisMonday.Year(), thisMonday.Month(), thisMonday.Day(), 0, 0, 0, 0, thisMonday.Location())
    thisFriday = time.Date(thisFriday.Year(), thisFriday.Month(), thisFriday.Day(), 0, 0, 0, 0, thisFriday.Location())

	return thisMonday.Format("2006-01-02"), thisFriday.Format("2006-01-02")
}

func GetLastWeekDates() (string, string) {
	now := time.Now()
	weekday := now.Weekday()
	daysToSubtract := 0

	if weekday == time.Sunday {
		daysToSubtract = 6
	} else {
		daysToSubtract = int(weekday) - int(time.Monday) + 7
	}

	lastMonday := now.AddDate(0, 0, -daysToSubtract)
	lastFriday := lastMonday.Add(5 * 24 * time.Hour)

    lastMonday = time.Date(lastMonday.Year(), lastMonday.Month(), lastMonday.Day(), 0, 0, 0, 0, lastMonday.Location())
    lastFriday = time.Date(lastFriday.Year(), lastFriday.Month(), lastFriday.Day(), 0, 0, 0, 0, lastFriday.Location())

	return lastMonday.Format("2006-01-02"), lastFriday.Format("2006-01-02")
}

func FormatDate(input string) (string, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, input)
	if err != nil {
		return "", err
	}
	return t.Format("2006-01-02") + "T00:00:00Z", nil
}
