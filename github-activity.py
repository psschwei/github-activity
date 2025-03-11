#! /usr/bin/env python3
# Assisted by watsonx Code Assistant
from datetime import datetime, timedelta
import json
import os
import requests
import sys


def run_github_query(username, start_dt, end_dt):
    ibmgithub = os.getenv("IBMGITHUB", "")
    url = "https://api.github.com/graphql"
    if ibmgithub:
        url = "https://github.ibm.com/api/graphql"

    token = os.getenv("GITHUB_TOKEN", "")
    if ibmgithub:
        token = os.getenv("GHE_TOKEN", "")

    if not token:
        exit(1)

    query = f"""
    query {{
        user(login: "{username}") {{
            contributionsCollection(from: "{start_dt}", to: "{end_dt}") {{
              issueContributions (first: 100) {{ 
                edges {{ 
                    node {{ 
                        occurredAt
                        issue {{
                            title
                            url
                            repository {{
                                nameWithOwner
                            }}
                        }}
                    }}
                }}
              }}
              pullRequestContributions (first: 100) {{
                edges {{ 
                    node {{ 
                        occurredAt
                        pullRequest {{ 
                            title 
                            url
                            repository {{
                                nameWithOwner
                            }}
                        }}
                    }}
                }}
              }}
              pullRequestReviewContributions (first: 100) {{ 
                edges {{
                    node {{
                        occurredAt
                        pullRequest {{
                          title
                          url
                          repository {{
                              nameWithOwner
                          }}
                        }}
                    }}
                }}
              }}
          }}
       }}
    }}
    """

    headers = {
        "Authorization": f"token {token}",
        "Content-Type": "application/json",
    }

    response = requests.post(url, json={"query": query}, headers=headers)
    return response.json()


def output_contributions(data):
    issues = data["data"]["user"]["contributionsCollection"]["issueContributions"][
        "edges"
    ]
    prs = data["data"]["user"]["contributionsCollection"]["pullRequestContributions"][
        "edges"
    ]
    reviews = data["data"]["user"]["contributionsCollection"][
        "pullRequestReviewContributions"
    ]["edges"]

    # Return PRs
    if prs:
        print("Pull Requests:")
        output = {}
        for pr in prs:
            repo = pr["node"]["pullRequest"]["repository"]["nameWithOwner"]
            title = pr["node"]["pullRequest"]["title"]
            url = pr["node"]["pullRequest"]["url"]
            if repo in output.keys():
                output[repo].append({"title": title, "url": url})
            else:
                output[repo] = [{"title": title, "url": url}]

        for key in output.keys():
            print(f"* {key}")
            for i in output[key]:
                print(f"    - {i["title"]}: {i["url"]}")

    # Return Reviews
    if reviews:
        print("Reviews:")
        output = {}
        for review in reviews:
            repo = review["node"]["pullRequest"]["repository"]["nameWithOwner"]
            title = review["node"]["pullRequest"]["title"]
            url = review["node"]["pullRequest"]["url"]
            if repo in output.keys():
                output[repo].append({"title": title, "url": url})
            else:
                output[repo] = [{"title": title, "url": url}]

        for key in output.keys():
            print(f"* {key}")
            for i in output[key]:
                print(f"    - {i["title"]}: {i["url"]}")

    # Return Issues
    if issues:
        print("Issues:")
        output = {}
        for issue in issues:
            repo = issue["node"]["issue"]["repository"]["nameWithOwner"]
            title = issue["node"]["issue"]["title"]
            url = issue["node"]["issue"]["url"]
            if repo in output.keys():
                output[repo].append({"title": title, "url": url})
            else:
                output[repo] = [{"title": title, "url": url}]

        for key in output.keys():
            print(f"* {key}")
            for i in output[key]:
                print(f"    - {i["title"]}: {i["url"]}")


if __name__ == "__main__":

    if len(sys.argv) > 1:
        username = sys.argv[1]
    else:
        print(f"INFO: Assuming github user ID is {os.getlogin}")
        username = os.getlogin()

    if len(sys.argv) > 2:
        start_dt = sys.argv[2] + "T00:00:00Z"
    else:
        print("INFO: No start date provided, will look back seven days from today")
        start_dt = (
            datetime.today().date() - timedelta(days=7)
        ).isoformat() + "T00:00:00Z"

    if len(sys.argv) > 3:
        end_dt = sys.argv[3] + "T00:00:00Z"
    else:
        print("INFO: No end date provided, so will including everying up to... NOW")
        end_dt = (
            datetime.today().date() + timedelta(days=1)
        ).isoformat() + "T00:00:00Z"

    print(f"Getting all activiting between {start_dt} and {end_dt} for {username}\n\n")

    data = run_github_query(username, start_dt, end_dt)

    output_contributions(data)
