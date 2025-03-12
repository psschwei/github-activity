package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func getGithubUrl(d string) string {

	url := ""
	switch d {
	case "github.com":
		url = "https://api.github.com/graphql"
	case "github.ibm.com":
		url = "https://github.ibm.com/api/graphql"
	}
	return url
}

func getGithubQuery(u, s, e string) string {
	query := fmt.Sprintf(`    query {
        user(login: "%s") {
            contributionsCollection(from: "%s", to: "%s") {
              issueContributions (first: 100) {
                edges {
                    node {
                        occurredAt
                        issue {
                            title
                            url
                            repository {
                                nameWithOwner
                            }
                        }
                    }
                }
              }
              pullRequestContributions (first: 100) {
                edges {
                    node {
                        occurredAt
                        pullRequest {
                            title
                            url
                            repository {
                                nameWithOwner
                            }
                        }
                    }
                }
              }
              pullRequestReviewContributions (first: 100) {
                edges {
                    node {
                        occurredAt
                        pullRequest {
                          title
                          url
                          repository {
                              nameWithOwner
                          }
                        }
                    }
                }
              }
          }
       }
    }`, u, s, e)

	return query
}

// Assisted by watsonx Code Assistant
type gitHubActivity struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				IssueContributions struct {
					Edges []struct {
						Node struct {
							OccurredAt string `json:"occurredAt"`
							Issue      struct {
								Title      string `json:"title"`
								Url        string `json:"url"`
								Repository struct {
									NameWithOwner string `json:"nameWithOwner"`
								} `json:"repository"`
							} `json:"issue"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"issueContributions"`
				PullRequestContributions struct {
					Edges []struct {
						Node struct {
							OccurredAt  string `json:"occurredAt"`
							PullRequest struct {
								Title      string `json:"title"`
								Url        string `json:"url"`
								Repository struct {
									NameWithOwner string `json:"nameWithOwner"`
								} `json:"repository"`
							} `json:"pullRequest"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"pullRequestContributions"`
				PullRequestReviewContributions struct {
					Edges []struct {
						Node struct {
							OccurredAt  string `json:"occurredAt"`
							PullRequest struct {
								Title      string `json:"title"`
								Url        string `json:"url"`
								Repository struct {
									NameWithOwner string `json:"nameWithOwner"`
								} `json:"repository"`
							} `json:"pullRequest"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"pullRequestReviewContributions"`
			}
		}
	}
}

func queryGithubApi(url, query, token string, userActivity *gitHubActivity) error {
	payload := map[string]string{"query": query}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	req.Header.Add("Authorization", "token "+token)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("unable to query github api")
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("error querying github api: " + res.Status)
	}

	resBody, _ := io.ReadAll(res.Body)

	if err = json.Unmarshal(resBody, &userActivity); err != nil {
		fmt.Println("cannot unmarshal json")
		return err
	}

	return nil

}

func printActivityOutput(userActivity *gitHubActivity) {
	data := userActivity.Data.User.ContributionsCollection

	if len(data.PullRequestContributions.Edges) > 0 {
		fmt.Println("Pull Requests")
	}
	for _, v := range data.PullRequestContributions.Edges {
		fmt.Println("repo:  " + v.Node.PullRequest.Repository.NameWithOwner)
		fmt.Println("title: " + v.Node.PullRequest.Title)
		fmt.Println("url:   " + v.Node.PullRequest.Url + "\n")
	}

	if len(data.PullRequestReviewContributions.Edges) > 0 {
		fmt.Println("Reviews")
	}
	for _, v := range data.PullRequestReviewContributions.Edges {
		fmt.Println("repo:  " + v.Node.PullRequest.Repository.NameWithOwner)
		fmt.Println("title: " + v.Node.PullRequest.Title)
		fmt.Println("url:   " + v.Node.PullRequest.Url + "\n")
	}

	if len(data.IssueContributions.Edges) > 0 {
		fmt.Println("Issues")
	}
	for _, v := range data.IssueContributions.Edges {
		fmt.Println("repo:  " + v.Node.Issue.Repository.NameWithOwner)
		fmt.Println("title: " + v.Node.Issue.Title)
		fmt.Println("url:   " + v.Node.Issue.Url + "\n")
	}

}
