#! /bin/bash
set -e

which gdate > /dev/null && date_cmd="gdate" || date_cmd="date"

starting_date=$($date_cmd -dlast-monday +"%Y-%m-%d")
gh_search_flags="--updated >=$starting_date --json repository,title,url"

echo "PRs ($(gh search prs --author @me $gh_search_flags | jq 'length'))"
gh search prs --author @me $gh_search_flags --template '{{ range . }}{{ printf "* [%s] %s: %s\n" .repository.nameWithOwner .title .url }}{{ end }}'

echo "Reviews ($(gh search prs --reviewed-by @me $gh_search_flags | jq 'length'))"
gh search prs --reviewed-by @me $gh_search_flags --template '{{ range . }}{{ printf "* [%s] %s: %s\n" .repository.nameWithOwner .title .url }}{{ end }}'

echo "Issues ($(gh search issues --author @me $gh_search_flags | jq 'length'))"
gh search issues --author @me $gh_search_flags --template '{{ range . }}{{ printf "* [%s] %s: %s\n" .repository.nameWithOwner .title .url }}{{ end }}'

