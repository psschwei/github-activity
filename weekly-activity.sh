#! /bin/bash
set -e

starting_date=$(date -d "7 days ago" +"%Y-%m-%d")

# PRs
gh search prs --author "@me" --created ">=$starting_date"

# Reviews
gh search prs --reviewed-by "@me" --created ">=$starting_date"

# Issues
gh search issues --author "@me" --created ">=$starting_date"
