/*
Copyright © 2023 Okteto Inc.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/cobra"
)

const colorOktetoGreen = "00d1ca"

var errAlreadyExists = errors.New("environment already exists")

var ldProjectSource string
var ldEnvironmentColor string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a LaunchDarkly environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		ldEnvironmentsURL := fmt.Sprintf("https://app.launchdarkly.com/api/v2/projects/%s/environments", ldProjectKey)

		var env = environment{
			Key:   ldEnvironmentName,
			Name:  ldEnvironmentName,
			Color: ldEnvironmentColor,
		}

		if ldProjectSource != "" {
			env.Source = environmentSource{Key: ldProjectSource}
		}

		e, err := createEnviroment(env, ldEnvironmentsURL)

		if err == nil {
			return publishResults(e)
		}

		if errors.Is(err, errAlreadyExists) {
			e, err := getExistingEnvironment()
			if err != nil {
				return err
			}

			return publishResults(e)
		}

		return fmt.Errorf("failed to create the LaunchDarkly environment: %w", err)
	},
}

func createEnviroment(e environment, ldEnvironmentsURL string) (environment, error) {
	marshaled, err := json.Marshal(e)
	if err != nil {
		return environment{}, fmt.Errorf("failed to parse LaunchDarkly environment request")
	}

	request, err := retryablehttp.NewRequest("POST", ldEnvironmentsURL, bytes.NewBuffer(marshaled))
	if err != nil {
		return environment{}, fmt.Errorf("failed to start the request to create LaunchDarkly environment: %w", err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", ldAccessToken)

	client := getRetryableClient()
	response, err := client.Do(request)
	if err != nil {
		return environment{}, fmt.Errorf("failed to clone the LaunchDarkly environment: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode == 409 {
		return environment{}, errAlreadyExists
	}

	if response.StatusCode > 300 {
		return environment{}, fmt.Errorf(response.Status)
	}

	var newEnv environment
	if err := json.NewDecoder(response.Body).Decode(&newEnv); err != nil {
		return environment{}, fmt.Errorf("failed to decode the response of the LaunchDarkly API: %w", err)
	}

	return newEnv, nil

}

func getExistingEnvironment() (environment, error) {
	ldEnvironmentURL := getEnvironmentURL(ldProjectKey, ldEnvironmentName)
	request, err := retryablehttp.NewRequest("GET", ldEnvironmentURL, nil)
	if err != nil {
		return environment{}, fmt.Errorf("failed to start the request to get a LaunchDarkly environment: %w", err)
	}

	request.Header.Set("Authorization", ldAccessToken)
	client := getRetryableClient()
	response, err := client.Do(request)
	if err != nil {
		return environment{}, fmt.Errorf("failed to get the LaunchDarkly environment: %w", err)
	}

	defer response.Body.Close()

	var existingEnvironment = environment{}
	if err := json.NewDecoder(response.Body).Decode(&existingEnvironment); err != nil {
		return environment{}, fmt.Errorf("failed to decode environment from response: %w", err)
	}

	return existingEnvironment, nil
}

func publishResults(e environment) error {
	url := getEnvironmentFeaturesURL(ldProjectKey, ldEnvironmentName)
	if err := generateNotes(url, e.ID, e.ApiKey, e.MobileKey); err != nil {
		return fmt.Errorf("failed to create notes: %w", err)
	}

	return generateResults(url, e.ID, e.ApiKey, e.MobileKey)
}

func writeEnvFile(e environment) error {

}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&ldProjectSource, "source", "", "LaunchDarkly Project Source")
	createCmd.Flags().StringVar(&ldEnvironmentColor, "color", colorOktetoGreen, "The color of your LaunchDarkly environment")

}
