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

// serverDetailsCmd represents the serverDetails command
var serverDetailsCmd = &cobra.Command{
	Use:   "details SERVERID",
	Short: "Detailed server information.",
	Long: `Example:

glectl server details kvm12345

kvm12345
Hostname:       web.example.com
IPs:            192.168.1.199, fe80::199
Description:	Webserver-01
Datacenter:     Falkenberg
CPU:            2
RAM:            4096
Storage:        40`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return ServerList(cmd, args, toComplete)
	},
	Run: details,
}

func details(ccmd *cobra.Command, args []string) {
	client := helper.NewClient()

	id := args[0]

	server, err := client.Servers.Details(context.Background(), id)
	if err != nil {
		cobra.CheckErr(err)
	}

	var ips []string
	for _, ip := range server.IPList {
		ips = append(ips, ip.Address)
	}
	ips2 := strings.Join(ips, ", ")

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintf(
		writer,
		"%s\nHostname:\t%s\nIPs:\t%s\nDescription:\t%s\nDatacenter:\t%s\nCPU:\t%d\nRAM:\t%d\nStorage:\t%d\nState:\t%s\n",
		server.ID,
		server.Hostname,
		ips2,
		server.Description,
		server.DataCenter,
		server.CPU,
		server.Memory,
		server.Storage,
		server.State,
	)
	writer.Flush()
}

func ServerList(ccmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// log.Printf("ServerList called")
	client := helper.NewClient()

	var srvList []string
	servers, err := client.Servers.List(context.Background())
	if err != nil {
		cobra.CheckErr(err)
	}

	toComplete = strings.ToLower(toComplete)
	for _, srv := range *servers {
		if strings.HasPrefix(srv.ID, toComplete) {
			// log.Printf("floff %s\n", srv.ID)
			// TODO: see if we can output the hostname in some way without adding it to the autocomplete suggestion.
			// srvList = append(srvList, fmt.Sprintf("%s:%s", srv.ID, srv.Hostname))
			srvList = append(srvList, srv.ID)
		}
	}

	return srvList, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	serverCmd.AddCommand(serverDetailsCmd)
}
