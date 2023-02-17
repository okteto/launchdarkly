package cmd

import (
	"fmt"
)

type environmentSource struct {
	Key string `json:"key"`
}

type environment struct {
	ID        string            `json:"_id"`
	Key       string            `json:"key"`
	Name      string            `json:"name"`
	Color     string            `json:"color"`
	Source    environmentSource `json:"source"`
	ApiKey    string            `json:"apiKey"`
	MobileKey string            `json:"mobileKey"`
}

func getEnvironmentURL(project, name string) string {
	return fmt.Sprintf("https://app.launchdarkly.com/api/v2/projects/%s/environments/%s", project, name)
}

func getEnvironmentFeaturesURL(project, name string) string {
	return fmt.Sprintf("https://app.launchdarkly.com/api/v2/projects/%s/environments/%s/features", project, name)
}
