package github

import (
	"encoding/json"
	"fmt"
	"io"

	"github-activity/pkg/utils"
)

func getGithubIssuesQuery(repo string, labels []string, endDate string, startDate string, cursor string) string {
	labelStr := ""
	dateStr := ""
	for _, label := range labels {
		if label != "" {
			labelStr += fmt.Sprintf("label:\\\"%s\\\" ", label)
		}
	}

	if endDate != utils.GetDefaultEndDate() {
		dateStr += fmt.Sprintf(` updated:<%s `, endDate)
	}

	if startDate != utils.GetDefaultStartDate() {
		dateStr += fmt.Sprintf(` created:>%s `, startDate)
	}

	issueQuery := fmt.Sprintf(`
	query {
		search(query: "repo:%s is:issue %s %s ", type: ISSUE, first: 100 , after: "%s") {
			issueCount
			edges {
				node {
					... on Issue {
						id
						number
						url
						title
						state
						labels(first: 10){
							edges {
								node {
									name
								}
							}
						}
						comments(first: 25) {
							totalCount
							edges {
								node {
									body
								}
							}
						}
						createdAt
						updatedAt
						closedAt
						repository{
							id
						}
						author {
							login
							url
						}
						assignees(first: 5){
							edges {
								node {
									login
									url
								}
							}
						}
						closedByPullRequestsReferences(first: 5){
							edges {
								node {
									number
									url
								}
							}
						}
						participants(first: 10){
							edges {
								node {
									login
									url
								}
							}
						}
					}
				}
			}
			pageInfo {
					hasNextPage
					endCursor
			}
		}
	}`, repo, labelStr, dateStr, cursor)

	return issueQuery
}

type gitHubIssues struct {
	Data struct {
		Search struct {
			IssueCount int `json:"issueCount"`
			Edges      []struct {
				Node IssueNode `json:"node"`
			} `json:"edges"`
			PageInfo struct {
				HasNextPage bool   `json:"hasNextPage"`
				EndCursor   string `json:"endCursor"`
			} `json:"pageInfo"`
		} `json:"search"`
	} `json:"data"`
}

type IssueNode struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
	URL    string `json:"url"`
	Title  string `json:"title"`
	State  string `json:"state"`
	Labels struct {
		Edges []struct {
			Node Label `json:"node"`
		} `json:"edges"`
	} `json:"labels"`
	Comments struct {
		TotalCount int `json:"totalCount"`
		Edges []struct {
			Node Comment `json:"node"`
		} `json:"edges"`
	} `json:"comments"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	ClosedAt     string `json:"closedAt"`
	Repository struct {
		ID string `json:"id"`
	} `json:"repository"`
	Author struct {
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"author"`
	Assignees struct {
		Edges []interface{} `json:"edges"`
	} `json:"assignees"`
	ClosedByPullRequestsReferences struct {
		Edges []interface{} `json:"edges"`
	} `json:"closedByPullRequestsReferences"`
	Participants struct {
		Edges []struct {
			Node Participant `json:"node"`
		} `json:"edges"`
	} `json:"participants"`
}

func queryGithubApiIssues(url, query, token string, issueData *gitHubIssues) error {
	res, err := callGithubApi(url, query, token)
	if err != nil {
		return err
	}

	resBody, _ := io.ReadAll(res.Body)

	if err = json.Unmarshal(resBody, &issueData); err != nil {
		fmt.Println("cannot unmarshal json")
		return err
	}

	return nil

}

func printIssueDataOutput(issueData *gitHubIssues, output string) {
	jsonPretty, err := json.MarshalIndent(issueData, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}

	if err = printToFile(string(jsonPretty), output); err != nil {
		fmt.Println(string(jsonPretty))
	}
}
