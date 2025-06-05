package github

import (
	"errors"
	"fmt"
	"os"

	"github-activity/pkg/utils"
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

	fmt.Println(domain + " activity for " + username + " between " + startDate + " and " + endDate + ":\n")
	printActivityOutput(&userActivity)

	return nil
}

func GetPRData(domain, token, repo, output, enddate, startdate string, labels []string) error {
	var prData gitHubPrs
	var nextPage gitHubPrs

	cursor := ""

	url := getGithubUrl(domain)

	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	if token == "" {
		return errors.New("no Github token specified")
	}

	prQuery := getGithubPrsQuery(repo, labels, enddate, startdate, cursor)

	if err := queryGithubApiPrs(url, prQuery, token, &nextPage); err != nil {
		return err
	}

	prData = nextPage

	for nextPage.Data.Search.PageInfo.HasNextPage {
		prQuery := getGithubPrsQuery(repo, labels, enddate, startdate, nextPage.Data.Search.PageInfo.EndCursor)

		if err := queryGithubApiPrs(url, prQuery, token, &nextPage); err != nil {
			return err
		}
		prData.Data.Search.Edges = append(prData.Data.Search.Edges, nextPage.Data.Search.Edges...)
		prData.Data.Search.PageInfo = nextPage.Data.Search.PageInfo
	}

	printPrDataOutput(&prData, output)

	return nil

}

func GetIssuesData(domain, token, repo, output, enddate, startdate string, labels []string) error {
	var issueData gitHubIssues
	var nextPage gitHubIssues

	cursor := ""

	url := getGithubUrl(domain)

	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	if token == "" {
		return errors.New("no Github token specified")
	}

	issueQuery := getGithubIssuesQuery(repo, labels, enddate, startdate, cursor)

	if err := queryGithubApiIssues(url, issueQuery, token, &nextPage); err != nil {
		return err
	}

	issueData = nextPage

	for nextPage.Data.Search.PageInfo.HasNextPage {
		issueQuery := getGithubIssuesQuery(repo, labels, enddate, startdate, nextPage.Data.Search.PageInfo.EndCursor)

		if err := queryGithubApiIssues(url, issueQuery, token, &nextPage); err != nil {
			return err
		}
		issueData.Data.Search.Edges = append(issueData.Data.Search.Edges, nextPage.Data.Search.Edges...)
		issueData.Data.Search.PageInfo = nextPage.Data.Search.PageInfo
	}

	printIssueDataOutput(&issueData, output)

	return nil

}