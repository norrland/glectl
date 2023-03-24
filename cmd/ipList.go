/*
Copyright © 2023 Andreas Eriksson <norrland@nullbyte.se>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/glesys/glesys-go/v7"
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

		userid, token := helper.GetCredentials()
		agent := helper.UserAgent()

		client := glesys.NewClient(userid, token, agent)

		ipv6, _ := cmd.Flags().GetBool("ipv6")
		ipv4, _ := cmd.Flags().GetBool("ipv4")

		log.Printf("ipv4/ipv6? %t/%t", ipv4, ipv6)

		var ipParams glesys.ReservedIPsParams
		switch af := list_inetf(ipv4, ipv6); af {
		case "ipv4":
			ipParams = glesys.ReservedIPsParams{Version: 4}
		case "ipv6":
			ipParams = glesys.ReservedIPsParams{Version: 6}
		default:
			log.Println("No af set.")
		}

		ips, err := client.IPs.Reserved(context.Background(), ipParams)
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
	ipListCmd.Flags().BoolP("ipv6", "6", false, "list ipv6 addresses only.")
	ipListCmd.Flags().BoolP("ipv4", "4", false, "list ipv4 addresses only.")
	ipListCmd.MarkFlagsMutuallyExclusive("ipv4", "ipv6")
}

func do_listips() {}

func list_inetf(ipv4 bool, ipv6 bool) string {
	if ipv4 {
		return "ipv4"
	}
	if ipv6 {
		return "ipv6"
	}
	return "any"
}
