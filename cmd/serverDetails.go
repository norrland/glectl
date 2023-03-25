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
	Use:   "details",
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
		"%s\nHostname:\t%s\nIPs:\t%s\nDescription:\t%s\nDatacenter:\t%s\nCPU:\t%d\nRAM:\t%d\nStorage:\t%d\n",
		server.ID,
		server.Hostname,
		ips2,
		server.Description,
		server.DataCenter,
		server.CPU,
		server.Memory,
		server.Storage,
	)
	writer.Flush()
}

func init() {
	serverCmd.AddCommand(serverDetailsCmd)
}
