package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func callGithubApi(url, query, token string) (*http.Response, error) {
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
		return res, fmt.Errorf("unable to query github api")
	}

	if res.StatusCode != 200 {
		return res, errors.New("error querying github api: " + res.Status)
	}

	return res, err

}

func queryGithubApi(url, query, token string, userActivity *gitHubActivity) error {
	res, err := callGithubApi(url, query, token)
	if err != nil {
		return err
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
		fmt.Println(fmt.Sprintf("Pull Requests (%d)", len(data.PullRequestContributions.Edges)))

		// Assisted by watsonx Code Assistant
		prOutput := make(map[string][]map[string]interface{})
		for _, review := range data.PullRequestContributions.Edges {
			repo := review.Node.PullRequest.Repository.NameWithOwner
			title := review.Node.PullRequest.Title
			url := review.Node.PullRequest.Url

			if _, exists := prOutput[repo]; !exists {
				prOutput[repo] = make([]map[string]interface{}, 0)
			}

			prOutput[repo] = append(prOutput[repo], map[string]interface{}{
				"title": title,
				"url":   url,
			})
		}

		for key, repos := range prOutput {
			fmt.Printf("* %s\n", key)
			for _, repo := range repos {
				fmt.Printf("    - %s: %s\n", repo["title"], repo["url"])
			}
		}
	}

	if len(data.PullRequestReviewContributions.Edges) > 0 {
		fmt.Println(fmt.Sprintf("Reviews (%d)", len(data.PullRequestReviewContributions.Edges)))

		// Assisted by watsonx Code Assistant
		reviewOutput := make(map[string][]map[string]interface{})
		for _, review := range data.PullRequestReviewContributions.Edges {
			repo := review.Node.PullRequest.Repository.NameWithOwner
			title := review.Node.PullRequest.Title
			url := review.Node.PullRequest.Url

			if _, exists := reviewOutput[repo]; !exists {
				reviewOutput[repo] = make([]map[string]interface{}, 0)
			}

			reviewOutput[repo] = append(reviewOutput[repo], map[string]interface{}{
				"title": title,
				"url":   url,
			})
		}

		for key, repos := range reviewOutput {
			fmt.Printf("* %s\n", key)
			for _, repo := range repos {
				fmt.Printf("    - %s: %s\n", repo["title"], repo["url"])
			}
		}
	}

	if len(data.IssueContributions.Edges) > 0 {
		fmt.Println(fmt.Sprintf("Issues (%d)", len(data.IssueContributions.Edges)))

		// Assisted by watsonx Code Assistant
		issueOutput := make(map[string][]map[string]interface{})
		for _, review := range data.IssueContributions.Edges {
			repo := review.Node.Issue.Repository.NameWithOwner
			title := review.Node.Issue.Title
			url := review.Node.Issue.Url

			if _, exists := issueOutput[repo]; !exists {
				issueOutput[repo] = make([]map[string]interface{}, 0)
			}

			issueOutput[repo] = append(issueOutput[repo], map[string]interface{}{
				"title": title,
				"url":   url,
			})
		}

		for key, repos := range issueOutput {
			fmt.Printf("* %s\n", key)
			for _, repo := range repos {
				fmt.Printf("    - %s: %s\n", repo["title"], repo["url"])
			}
		}

	}

	fmt.Printf(fmt.Sprintf("\nTotals: PRs(%d) Reviews(%d) Issues(%d)", len(data.PullRequestContributions.Edges), len(data.PullRequestReviewContributions.Edges), len(data.IssueContributions.Edges)))
}

func printToFile(str, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(str); err != nil {
		return err
	}

	return nil
}
