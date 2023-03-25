/*
Copyright © 2023 Andreas Eriksson <norrland@nullbyte.se>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	helper "github.com/norrland/glectl/helpers"
	"github.com/spf13/cobra"
)

// ipListCmd represents the ipList command
var ipListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := helper.NewClient()

		ips, err := client.IPs.Reserved(context.Background())
		if err != nil {
			cobra.CheckErr(err)
		}

		var hasLBip bool
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		if len(*ips) > 0 {
			fmt.Fprintln(writer, "platform\tdatacenter\taddress\tserverid")
			for _, ip := range *ips {
				if strings.HasPrefix(ip.ServerID, "lb") {
					ip.Platform = "LoadBalancer*"
					hasLBip = true
				}
				fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", ip.Platform, ip.DataCenter, ip.Address, ip.ServerID)
			}
			if hasLBip {
				fmt.Fprintln(writer, "[INFO] LoadBalancer* - no platform returned for LoadBalancer IPs.")
			}
			writer.Flush()
		}

	},
}

func init() {
	ipCmd.AddCommand(ipListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
