package cmd

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed notes.md.tmpl
var notes embed.FS

const notesPath = "notes.md"

func generateNotesFile(url, clientID, apiKey, mobileKey string) error {
	tmpl, err := template.ParseFS(notes, "*.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse notes template: %w", err)
	}

	f, err := os.Create(notesPath)
	if err != nil {
		return fmt.Errorf("failed to create notes file: %w", err)
	}

	defer f.Close()

	config := map[string]string{
		"Url":       url,
		"SDKKey":    apiKey,
		"MobileKey": mobileKey,
		"ClientID":  clientID,
	}

	if err := tmpl.Execute(f, config); err != nil {
		return fmt.Errorf("failed to write notes file: %w", err)
	}

	return nil
}
