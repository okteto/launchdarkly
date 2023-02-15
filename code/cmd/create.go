/*
Copyright © 2023 Okteto Inc.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var ldProjectSource string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a LaunchDarkly environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		color := "00d1ca"

		ldEnvironmentsURL := fmt.Sprintf("https://app.launchdarkly.com/api/v2/projects/%s/environments", ldProjectKey)

		var environment = environment{
			Key:    ldName,
			Name:   ldName,
			Color:  color,
			Source: environmentSource{Key: ldProjectSource},
		}

		marshaled, err := json.Marshal(environment)
		if err != nil {
			return fmt.Errorf("failed to parse LaunchDarkly environment request")
		}

		request, err := http.NewRequest("POST", ldEnvironmentsURL, bytes.NewBuffer(marshaled))
		if err != nil {
			return fmt.Errorf("failed to start the request to create LaunchDarkly environment: %w", err)
		}

		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", ldToken)

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return fmt.Errorf("failed to clone the LaunchDarkly environment: %w", err)
		}

		defer response.Body.Close()

		if response.StatusCode == 200 || response.StatusCode == 409 {
			fmt.Printf("https://app.launchdarkly.com/%s/%s/features", ldProjectKey, ldName)
			return nil
		}

		return fmt.Errorf("failed to create the LaunchDarkly environment: %s", response.Status)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&ldProjectSource, "source", "production", "LaunchDarkly Project Source")

}
