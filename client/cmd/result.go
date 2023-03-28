package cmd

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed result.env.tmpl
var result embed.FS

const resultsPath = "ld_results.env"

func generateResultsFile(url, clientID, apiKey, mobileKey string) error {
	tmpl, err := template.ParseFS(result, "*.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse result template: %w", err)
	}

	f, err := os.Create(resultsPath)
	if err != nil {
		return fmt.Errorf("failed to create result file: %w", err)
	}

	defer f.Close()

	config := map[string]string{
		"Url":       url,
		"SDKKey":    apiKey,
		"MobileKey": mobileKey,
		"ClientID":  clientID,
	}

	if err := tmpl.Execute(f, config); err != nil {
		return fmt.Errorf("failed to write results file: %w", err)
	}

	return nil
}
