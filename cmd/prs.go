/*
Copyright © 2025 psschwei (paul@paulschweigert.com)
*/
package cmd

import (
	"github-activity/pkg/github"
	"github.com/spf13/cobra"
)

var prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "Get PR data",
	Long:  `Get PR data for a given repo and labels`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return github.GetPRData(domain, token, repo, output, enddate, startdate, labels,)
	},
}

func init() {
	prsCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "Github org/repo")
	prsCmd.PersistentFlags().StringArrayVarP(&labels, "label", "l", []string{}, "Issue/PR label")
	prsCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output Filename (JSON)")
	prsCmd.MarkPersistentFlagRequired("repo")
	rootCmd.AddCommand(prsCmd)
}
