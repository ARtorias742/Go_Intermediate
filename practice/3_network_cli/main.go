package main

import (
	"fmt"
	"network-cli/cmd"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "netcli",
		Short: "A simple network CLI tool",
		Long:  "A command-Line tool for basic network operations",
	}

	rootCmd.AddCommand(cmd.PingCmd)
	rootCmd.AddCommand(cmd.ScanCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
