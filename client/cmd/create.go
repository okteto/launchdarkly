/*
Copyright © 2023 Okteto Inc.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/cobra"
)

const colorOktetoGreen = "00d1ca"

var ldProjectSource string
var ldEnvironmentColor string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a LaunchDarkly environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		ldEnvironmentsURL := fmt.Sprintf("https://app.launchdarkly.com/api/v2/projects/%s/environments", ldProjectKey)

		var environment = environment{
			Key:   ldEnvironmentName,
			Name:  ldEnvironmentName,
			Color: ldEnvironmentColor,
		}

		if ldProjectSource != "" {
			environment.Source = environmentSource{Key: ldProjectSource}
		}

		marshaled, err := json.Marshal(environment)
		if err != nil {
			return fmt.Errorf("failed to parse LaunchDarkly environment request")
		}

		request, err := retryablehttp.NewRequest("POST", ldEnvironmentsURL, bytes.NewBuffer(marshaled))
		if err != nil {
			return fmt.Errorf("failed to start the request to create LaunchDarkly environment: %w", err)
		}

		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", ldAccessToken)

		client := getRetryableClient()
		response, err := client.Do(request)
		if err != nil {
			return fmt.Errorf("failed to clone the LaunchDarkly environment: %w", err)
		}

		defer response.Body.Close()

		if response.StatusCode < 300 || response.StatusCode == 409 {
			url := fmt.Sprintf("https://app.launchdarkly.com/%s/%s/features", ldProjectKey, ldEnvironmentName)
			if err := generateNotes(url); err != nil {
				return fmt.Errorf("failed to create notes: %w", err)
			}
			fmt.Print(url)
			return nil
		}

		return fmt.Errorf("failed to create the LaunchDarkly environment: %s", response.Status)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&ldProjectSource, "source", "", "LaunchDarkly Project Source")
	createCmd.Flags().StringVar(&ldEnvironmentColor, "color", colorOktetoGreen, "The color of your LaunchDarkly environment")

}
