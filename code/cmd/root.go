/*
Copyright © 2023 Okteto Inc.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "launchdarkly",
	Short: "Manage your LaunchDarkly environment as part of your Okteto environment",
}

var ldToken string
var ldProjectKey string
var ldName string

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&ldToken, "token", "", "LaunchDarkly API Token")
	rootCmd.PersistentFlags().StringVar(&ldProjectKey, "project", "okteto", "LaunchDarkly Project Key")
	rootCmd.PersistentFlags().StringVar(&ldName, "name", "okteto", "Name of the LaunchDarkly environment")

}
