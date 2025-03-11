# github-activity

This script retrieves the GitHub activity for a specified user and prints it to the console.

## Usage

To use this script, you will need to install the `requests`  library. You can do this by running the following command in your terminal:

```bash
pip install requests
```

You will also need to store your [Github personal access token](https://github.com/settings/personal-access-tokens) in an environmental variable:

```bash
export GITHUB_TOKEN=<token>
```

Once you have installed the `requests` library, you can run the script by providing the GitHub username as a command-line argument. For example, to retrieve the GitHub activity for the user `example-user`, you would run the following command:

```bash
python github-activity.py example-user
```

The script will then retrieve the GitHub activity for the specified user and print it to the console.

By default, the script will pull activity for the previous seven days. To choose a different date range, pass a start date and (optional) end date to the script.

```bash
python github-activity.py example-user 2024-01-01 2024-01-31
```

## Example Output

Here is an example of the output that the script might produce:

```
Getting all activiting between 2024-01-01T00:00:00Z and 2024-01-31T00:00:00Z for example-user
PRs:
* awesome/coder
    - feat(coder): create new coder: https://github.com/awesome/coder/pull/789
Reviews:
* sample/repo
    - adding contributing guidelines for demo: https://github.com/sample/repo/pull/283
Issues:
* some-org/some-repo
    - some issue: https://github.com/some-org/some-repo/issue/123
```

## Getting IBM Github contributions

To get contributions from IBM Github (github.ibm.com), you will need to store your [IBM Github personal access token](https://github.ibm.com/settings/tokens?type=beta) in an environmental variable:

```bash
export GHE_TOKEN=<token>
```

You will also need to pass `IBMGITHUB=yes` when running the command:

```bash
IBMGITHUB=yes python github-activity.py 
```

## License
This script is licensed under the Apache License, Version 2.0. You can find a copy of the license in the `LICENSE` file in the root directory of this repository.
