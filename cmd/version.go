package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// description for user
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "v0.0.0",
	Long:  `All software has versions. This is IEXP's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init Env CMD Program v0.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
