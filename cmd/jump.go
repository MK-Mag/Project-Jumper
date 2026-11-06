package cmd

import (
	"PJ/internal/Config"
	"PJ/internal/store"

	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// jumpCmd represents the jump command
var jumpCmd = &cobra.Command{
	Use:   "jump",
	Short: "jump to your saved project locations",

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := store.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		project, err := store.GetProject(db, args[0])
		if err != nil {
			return err
		}

		cfg, err := Config.ConfigLoad()
		if err != nil {
			return err
		}

		IDECMD := exec.Command(cfg.IDE, project.Path)
		if err := IDECMD.Start(); err != nil {
			return fmt.Errorf("failed to launch editor %q: %w", cfg.IDE, err)

		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(jumpCmd)
}
