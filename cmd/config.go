package cmd

import (
	"PJ/internal/Config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Project Jumper settings",
}
var ideCmd = &cobra.Command{
	Use:  "ide [name]",
	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := Config.ConfigLoad()
		if err != nil {
			return err
		}

		if len(args) == 0 {
			cmd.Printf("Current editor: %s\n", cfg.IDE)
			return nil
		}

		cfg.IDE = args[0]

		if err := Config.SaveConfig(cfg); err != nil {
			return err
		}

		cmd.Printf("IDE set to:%s\n",cfg.IDE)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(ideCmd)
}
