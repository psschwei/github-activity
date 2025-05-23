/*
Copyright © 2025 psschwei (paul@paulschweigert.com)
*/
package cmd

import (
	"os"

	"github.com/psschwei/github-activity/pkg/github"
	"github.com/psschwei/github-activity/pkg/utils"
	"github.com/spf13/cobra"
)

var domain string
var startdate string
var enddate string
var username string
var token string
var lastweek bool
var thisweek bool
var today bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Get your Github activity",
	Long:  `Get PRs, reviews, and issues created during a specific time interval.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if lastweek == true {
			startdate, enddate = utils.GetLastWeekDates()
		} else if thisweek == true {
			startdate, enddate = utils.GetThisWeekDates()
		} else if today == true {
			startdate, enddate = utils.GetTodayDates()
		}

		return github.GetGithubActivity(domain, startdate, enddate, username, token)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "github.com", "Github domain")
	rootCmd.PersistentFlags().StringVarP(&startdate, "start", "s", utils.GetDefaultStartDate(), "Collect activities starting on this date")
	rootCmd.PersistentFlags().StringVarP(&enddate, "end", "e", utils.GetDefaultEndDate(), "Collect activities up to this date")
	rootCmd.PersistentFlags().BoolVarP(&lastweek, "last-week", "l", false, "Collect activities for last week (last week Monday to last week Friday")
	rootCmd.PersistentFlags().BoolVarP(&thisweek, "this-week", "w", false, "Collect activities for this week (Monday to Friday")
	rootCmd.PersistentFlags().BoolVarP(&today, "today", "n", false, "Collect activities for today")
	rootCmd.PersistentFlags().StringVarP(&username, "user", "u", utils.GetCurrentUsername(), "Username")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Github Personal Access Token (default `$GITHUB_TOKEN`)")
}
