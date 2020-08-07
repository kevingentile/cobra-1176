package cmd

import "github.com/spf13/cobra"

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "experimental test command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
