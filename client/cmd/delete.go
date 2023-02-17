/*
Copyright © 2023 Okteto Inc.
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your LaunchDarkly environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		ldProjectKey := os.Getenv("LAUNCHDARKLY_PROJECT_KEY")
		ldAccessToken := os.Getenv("LAUNCHDARKLY_ACCESS_TOKEN")
		oktetoNamespace := os.Getenv("OKTETO_NAMESPACE")

		ldEnvironmentURL := fmt.Sprintf("https://app.launchdarkly.com/api/v2/projects/%s/environments/%s", ldProjectKey, oktetoNamespace)

		request, err := http.NewRequest("DELETE", ldEnvironmentURL, nil)
		if err != nil {
			return fmt.Errorf("failed to start the request to delete LaunchDarkly environment: %w", err)
		}

		request.Header.Set("Authorization", ldAccessToken)

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return fmt.Errorf("failed to delete the LaunchDarkly environment: %w", err)
		}

		defer response.Body.Close()

		if response.StatusCode == 204 || response.StatusCode == 404 {
			return nil
		}

		return fmt.Errorf("failed to delete the LaunchDarkly environment: %s", response.Status)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
