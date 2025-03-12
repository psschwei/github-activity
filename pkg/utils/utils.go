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

func FormatDate(input string) (string, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, input)
	if err != nil {
		return "", err
	}
	return t.Format("2006-01-02") + "T00:00:00Z", nil
}
