package github

import (
	"encoding/json"
	"fmt"
	"io"

	"github-activity/pkg/utils"
)

func getGithubPrsQuery(repo string, labels []string, endDate string, startDate string, cursor string) string {
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

	prQuery := fmt.Sprintf(`
	query {
		search(query: "repo:%s is:pr %s %s ", type: ISSUE, first: 100 , after: "%s") {
			issueCount
			edges {
				node {
					... on PullRequest {
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
						mergedAt
						baseRefName
						changedFiles
						additions
						deletions
						isDraft
						commits(first: 100) {
							nodes {
								commit {
									committedDate
								}
							}
						}
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
						closingIssuesReferences(first: 5){
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

	return prQuery
}

type gitHubPrs struct {
	Data struct {
		Search struct {
			IssueCount int `json:"issueCount"`
			Edges      []struct {
				Node Node `json:"node"`
			} `json:"edges"`
			PageInfo struct {
				HasNextPage bool   `json:"hasNextPage"`
				EndCursor   string `json:"endCursor"`
			} `json:"pageInfo"`
		} `json:"search"`
	} `json:"data"`
}

type Node struct {
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
	MergedAt     string `json:"mergedAt"`
	BaseRefName  string `json:"baseRefName"`
	ChangedFiles int    `json:"changedFiles"`
	Additions    int    `json:"additions"`
	Deletions    int    `json:"deletions"`
	IsDraft      bool   `json:"isDraft"`
	Commits      struct {
		Nodes []struct {
			Commit struct {
				CommittedDate string `json:"committedDate"`
			} `json:"commit"`
		} `json:"nodes"`
	} `json:"commits"`
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
	ClosingIssuesReferences struct {
		Edges []interface{} `json:"edges"`
	} `json:"closingIssuesReferences"`
	Participants struct {
		Edges []struct {
			Node Participant `json:"node"`
		} `json:"edges"`
	} `json:"participants"`
}

type Label struct {
	Name string `json:"name"`
}

type Comment struct {
	Name string `json:"body"`
}

type Participant struct {
	Login string `json:"login"`
	URL   string `json:"url"`
}

func queryGithubApiPrs(url, query, token string, prData *gitHubPrs) error {
	res, err := callGithubApi(url, query, token)
	if err != nil {
		return err
	}

	resBody, _ := io.ReadAll(res.Body)

	if err = json.Unmarshal(resBody, &prData); err != nil {
		fmt.Println("cannot unmarshal json")
		return err
	}

	return nil

}

func printPrDataOutput(prData *gitHubPrs, output string) {
	jsonPretty, err := json.MarshalIndent(prData, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}

	if err = printToFile(string(jsonPretty), output); err != nil {
		fmt.Println(string(jsonPretty))
	}
}
