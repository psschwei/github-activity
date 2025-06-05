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
	issuesCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "i-am-bee/beeai-framework", "Github org/repo")
	issuesCmd.PersistentFlags().StringArrayVarP(&labels, "label", "l", []string{"python"}, "Issue/PR label")
	issuesCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output Filename (JSON)")
	rootCmd.AddCommand(issuesCmd)
}
