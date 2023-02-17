/*
Copyright © 2023 Okteto Inc.
*/
package cmd

import (
	"fmt"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your LaunchDarkly environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		ldEnvironmentURL := fmt.Sprintf("https://app.launchdarkly.com/api/v2/projects/%s/environments/%s", ldProjectKey, ldEnvironmentName)

		request, err := retryablehttp.NewRequest("DELETE", ldEnvironmentURL, nil)
		if err != nil {
			return fmt.Errorf("failed to start the request to delete LaunchDarkly environment: %w", err)
		}

		request.Header.Set("Authorization", ldAccessToken)

		client := getRetryableClient()
		response, err := client.Do(request)
		if err != nil {
			return fmt.Errorf("failed to delete the LaunchDarkly environment: %w", err)
		}

		defer response.Body.Close()

		if response.StatusCode < 300 || response.StatusCode == 404 {
			return nil
		}

		return fmt.Errorf("failed to delete the LaunchDarkly environment: %s", response.Status)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
