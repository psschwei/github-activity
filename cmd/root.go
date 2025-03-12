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

var Domain string
var StartDate string
var EndDate string
var Username string
var Token string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Get your Github activity",
	Long:  `Get PRs, reviews, and issues created during a specific time interval.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return github.GetGithubActivity(Domain, StartDate, EndDate, Username, Token)
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
	rootCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "github.com", "Github domain")
	rootCmd.PersistentFlags().StringVarP(&StartDate, "start", "s", utils.GetDefaultStartDate(), "Collect activities starting on this date")
	rootCmd.PersistentFlags().StringVarP(&EndDate, "end", "e", utils.GetDefaultEndDate(), "Collect activities up to this date")
	rootCmd.PersistentFlags().StringVarP(&Username, "user", "u", utils.GetCurrentUsername(), "Username")
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Github Personal Access Token (default `$GITHUB_TOKEN`)")
}
