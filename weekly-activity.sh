#! /bin/bash

# USAGE: To get enterprise activity:
#   GH_HOST=github.ibm.com ./weekly-activity.sh

set -e

which gdate > /dev/null && date_cmd="gdate" || date_cmd="date"

starting_date=$($date_cmd -dlast-monday +"%Y-%m-%d")
gh_search_flags="--updated >=$starting_date --json repository,title,url"

prs_cmd="gh search prs --author @me $gh_search_flags"
reviewed_cmd="gh search prs --reviewed-by @me $gh_search_flags"
issues_cmd="gh search issues --author @me $gh_search_flags"

prs_len=$($prs_cmd | jq 'length')
reviews_len=$($reviewed_cmd -- -author:@me | jq 'length')
issues_len=$($issues_cmd | jq 'length')

echo "${GH_HOST:=github.com} activity since $starting_date"
echo

if [[ $prs_len > 0 ]]; then
  echo "PRs ($prs_len)"
  $prs_cmd --template '{{ range . }}{{ printf "* [%s] %s: %s\n" .repository.nameWithOwner .title .url }}{{ end }}'
fi

if [[ $reviews_len > 0 ]]; then
echo "Reviews ($reviews_len)"
$reviewed_cmd --template '{{ range . }}{{ printf "* [%s] %s: %s\n" .repository.nameWithOwner .title .url }}{{ end }}' -- -author:@me
fi

if [[ $issues_len > 0 ]]; then
echo "Issues ($issues_len)"
$issues_cmd --template '{{ range . }}{{ printf "* [%s] %s: %s\n" .repository.nameWithOwner .title .url }}{{ end }}'
fi

echo
echo "Totals: PRs($prs_len) Reviews($reviews_len) Issues($issues_len)"
