/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"PJ/internal/store"
	"fmt"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved projects",

	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := store.Open()
		if err != nil {
			return fmt.Errorf("open database: %w", err)
		}
		defer db.Close()

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 4, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tPATH\tCREATED AT")

		projects, err := store.ListProjects(db)
		if err != nil {
			return err
		}

		for _, project := range projects {
			fmt.Fprintf(
				w,
				"%s\t%s\t%s\n",
				project.Name,
				project.Path,
				project.CreatedAt,
			)
		}

		return w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
