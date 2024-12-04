package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/arichardsmith/aoc-2024-go/internal"
	"github.com/spf13/cobra"
)

func registerCreateCommand(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new day's challenge",
		Run:   runCreate,
	}

	cmd.PersistentFlags().IntP("day", "d", 0, "Day of advent (1-25)")
	cmd.MarkPersistentFlagRequired("day")

	root.AddCommand(cmd)
}

func runCreate(cmd *cobra.Command, args []string) {
	day, _ := cmd.Flags().GetInt("day")
	templateData := NewTemplateData(day)

	srcDir, err := internal.HereN(1, "../challenge/template")
	if err != nil {
		fmt.Printf("Error determining template directory: %v\n", err)
		return
	}

	destDir, err := internal.HereN(1, fmt.Sprintf("../challenge/day%02d", day))
	if err != nil {
		fmt.Printf("Error determining destination directory: %v\n", err)
		return
	}

	if _, err := os.Stat(destDir); err == nil {
		fmt.Printf("Code already exists for day %d\n", day)
		return
	}

	// Create destination directory if it doesn't exist
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Walk through template directory
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Skip the source directory itself
		if path == srcDir {
			return nil
		}

		// Calculate relative and destination paths
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destDir, strings.TrimSuffix(relPath, ".template"))

		// Read template file
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading template file: %v", err)
		}

		// Process template
		tmpl, err := template.New(relPath).Parse(string(content))
		if err != nil {
			return fmt.Errorf("error parsing template: %v", err)
		}

		// Create destination file
		destFile, err := os.Create(destPath)
		if err != nil {
			return fmt.Errorf("error creating destination file: %v", err)
		}
		defer destFile.Close()

		// Execute template and write to destination
		if err := tmpl.Execute(destFile, templateData); err != nil {
			return fmt.Errorf("error executing template: %v", err)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error processing templates: %v\n", err)
		return
	}

	fmt.Printf("Created boilerplate for day %d\n", day)
}

type TemplateData struct {
	Day    string
	DayPad string
}

func NewTemplateData(day int) TemplateData {
	return TemplateData{
		Day:    fmt.Sprintf("%d", day),
		DayPad: fmt.Sprintf("%02d", day),
	}
}
