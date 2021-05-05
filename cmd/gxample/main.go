package main

import (
	"fmt"
	"os"

	"github.com/changyoungkwon/gxample/cmd/gxample/cli"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gxample",
	Short: "Gxample",
}

func main() {
	RootCmd.AddCommand(cli.ServeCmd)
	RootCmd.AddCommand(cli.GendocCmd)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
