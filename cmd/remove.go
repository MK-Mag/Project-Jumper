package cmd

import (
	"PJ/internal/store"
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <project_name>",
	Short: "Remove a saved project",
	
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		
		db, err := store.Open()
		if err != nil{
			return fmt.Errorf("failed to open database: %w", err)
		}
		defer db.Close()

		if err := store.RemoveProject(db,args[0]); err != nil {
			return err
		}

		cmd.Printf("Removed project %q\n", args[0])

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
