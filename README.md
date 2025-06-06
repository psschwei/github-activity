# github-activity

This utility retrieves the GitHub activity for a specified user and prints it to the console.

## Usage

Download the binary from the [Releases page](https://github.com/psschwei/github-activity/releases) and add it to your `PATH`.

To see a list of options, run `gha --help`:

```bash
$ gha --help
Get PRs, reviews, and issues created during a specific time interval.

Usage:
  gha [flags]
  gha [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  issues      Get GitHub Issues data
  prs         Get PR data

Flags:
  -d, --domain string         Github domain (default "github.com")
  -e, --end string            Collect activities up to this date (default "2025-03-11")
  -h, --help                  help for github-activity
  -l, --last-week             Collect activities for last week (last week Monday to last week Friday)
  -s, --start string          Collect activities starting on this date (default "2025-03-04")
  -w, --this-week             Collect activities for this week (Monday to Friday)
  -n, --today                 Collect activities for today
  -t, --token $GITHUB_TOKEN   Github Personal Access Token (default $GITHUB_TOKEN)
  -u, --user string           Username (default "paulschw")

Use "github-activity [command] --help" for more information about a command.
```

The utility will then retrieve the GitHub activity for the specified user and print it to the console.

By default, it will pull activity for the previous seven days. To choose a different date range, pass a start date and (optional) end date to the script.

```bash
gha -s 2025-01-01 -e 2025-02-28
```

## Example Output

Here is an example of the output that the script might produce:

```
$ ./gha -s 2025-03-01 -u psschwei
github.com activity for psschwei between 2025-03-01T00:00:00Z and 2025-03-11T00:00:00Z:

Pull Requests (1)
*  i-am-bee/beeai
    - feat(agents): add open deep research agent: https://github.com/i-am-bee/beeai/pull/253

Reviews (3)
* i-am-bee/beeai-labs
    - first version of Workflow::to_mermaid method: https://github.com/i-am-bee/beeai-labs/pull/296

* i-am-bee/beeai
    - feat(agents): add open deep research agent: https://github.com/i-am-bee/beeai/pull/253

* i-am-bee/beeai-labs
    - adding contributing guidelines for demo: https://github.com/i-am-bee/beeai-labs/pull/283
```

## Getting IBM Github contributions

To get contributions from IBM Github (github.ibm.com), you will need to use your [IBM Github personal access token](https://github.ibm.com/settings/tokens?type=beta) as the token and pass `github.ibm.com` as the domain:

```bash
gha -d github.ibm.com -t <IBM_GITHUB_TOKEN>
```

## Getting PR Data

The utility can also gather PR data for a given repo and label using the following command:

```bash
gha prs -r i-am-bee/beeai-framework --label python
```

## Getting Issue Data

The utility can also gather Issue data for a given repo and label using the following command:

```bash
gha issues -r i-am-bee/beeai-framework --label python
```


## Contributing

To build the utility locally run the following command:

```bash
go build -o gha
```

## License
This script is licensed under the Apache License, Version 2.0. You can find a copy of the license in the `LICENSE` file in the root directory of this repository.
