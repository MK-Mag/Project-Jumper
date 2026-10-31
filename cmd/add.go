/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"PJ/internal/store"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <project_name> <project_path>",
	Short: "Save a project directory",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		name := args[0]
		projpath := args[1]

		info, err := os.Stat(projpath)

		if err != nil {
			return fmt.Errorf("os.Stat: %w", err)
		}

		if !info.IsDir() {
			return fmt.Errorf("project path is not a directory")
		}

		abspath, err := filepath.Abs(projpath)
		if err != nil {
			return fmt.Errorf("resolve absolute path: %w", err)
		}

		db, err := store.Open()
		if err != nil {
			return fmt.Errorf("open database: %w", err)
		}
		defer db.Close()

		if err := store.AddProject(db, name, abspath); err != nil {
			return err
		}

		fmt.Printf("Saved project %q\nPath: %s\n", name, abspath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
