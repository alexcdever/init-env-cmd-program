package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/spf13/cobra"
)

var info, _ = host.Info()
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect the OS info",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("OS: %v", info.OS)
		fmt.Printf("platform: %v", info.Platform)
		fmt.Printf("family: %v", info.PlatformFamily)
		fmt.Printf("version: %v", info.PlatformVersion)
	},
}

func init() {
	rootCmd.AddCommand(collectCmd)
}
