/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jumpCmd represents the jump command
var jumpCmd = &cobra.Command{
	Use:   "jump",
	Short: "jump to your saved project locations",
	Long:  "jump to your saved project locations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Requested Project: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(jumpCmd)

}
