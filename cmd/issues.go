/*
Copyright © 2025 psschwei (paul@paulschweigert.com)
*/
package cmd

import (
	"github-activity/pkg/github"
	"github.com/spf13/cobra"
)

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Get GitHub Issues data",
	Long:  `Get GitHub Issues data for a given repo and labels`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return github.GetIssuesData(domain, token, repo, output, enddate, startdate, labels)
	},
}

func init() {
	issuesCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "Github org/repo")
	issuesCmd.PersistentFlags().StringArrayVarP(&labels, "label", "l", []string{}, "Issue/PR label")
	issuesCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output Filename (JSON)")
	issuesCmd.MarkPersistentFlagRequired("repo")
	rootCmd.AddCommand(issuesCmd)
}
