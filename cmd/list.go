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

		rows, err := db.Query("SELECR name,path,created_at FROM projects ORDER BY name,")

		if err != nil {
			return fmt.Errorf("query projects: %w", err)                            // L30
		}
		defer rows.Close()

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 4, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tPATH\tCREATED AT")

		for rows.Next() {
			var name, path, createdAt string

			if err := rows.Scan(&name, &path, &createdAt); err != nil {
				return fmt.Errorf("read project: %w", err)
			}

			fmt.Fprintf(w, "%s\t%s\t%s\n", name, path, createdAt)

		}

		if err := rows.Err(); err != nil {
			return fmt.Errorf("read project: %w", err)
		}

		return w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
