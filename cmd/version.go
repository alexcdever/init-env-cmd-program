package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
// description for user
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "v0.0.0",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
