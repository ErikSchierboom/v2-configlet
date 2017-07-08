package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "3.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the current version of the tool.",
	Long:  "Output the current version of the tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("configlet v%s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
