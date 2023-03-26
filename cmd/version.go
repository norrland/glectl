/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	helper "github.com/norrland/glectl/helpers"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	//Long: `A longer description that spans multiple lines and likely contains examples
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s\n", rootCmd.Root().Name(), helper.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
