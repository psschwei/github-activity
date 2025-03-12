package github

import (
	"errors"
	"fmt"
	"os"

	"github.com/psschwei/github-activity/pkg/utils"
)

func GetGithubActivity(domain, startDate, endDate, username, token string) error {
	var userActivity gitHubActivity

	startDate, err := utils.FormatDate(startDate)
	if err != nil {
		return err
	}

	endDate, err = utils.FormatDate(endDate)
	if err != nil {
		return err
	}

	url := getGithubUrl(domain)
	query := getGithubQuery(username, startDate, endDate)

	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	if token == "" {
		return errors.New("No Github token specified")
	}

	if err := queryGithubApi(url, query, token, &userActivity); err != nil {
		return err
	}

	fmt.Println("Github Activity for " + username + " on " + domain + " between " + startDate + " and " + endDate + ":\n")
	printActivityOutput(&userActivity)

	return nil
}
