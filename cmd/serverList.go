/*
Copyright © 2023 Andreas Eriksson <norrland@nullbyte.se>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	helper "github.com/norrland/glectl/helpers"
	"github.com/spf13/cobra"
)

// serverListCmd represents the serverList command
var serverListCmd = &cobra.Command{
	Use:   "list",
	Short: "List servers in your current project.",
	//Long: `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		dontPrintHeader, _ := cmd.Flags().GetBool("no-headers")

		client := helper.NewClient()

		servers, err := client.Servers.List(context.Background())

		if err != nil {
			cobra.CheckErr(err)
		}

		// TODO: fix nice tabs
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		if len(*servers) > 0 {
			if !dontPrintHeader {
				fmt.Fprintln(writer, "platform\tdatacenter\tid\thostname")
			}
			for _, srv := range *servers {
				_, err := fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", srv.Platform, srv.DataCenter, srv.ID, srv.Hostname)
				if err != nil {
					cobra.CheckErr(err)
				}
			}
			writer.Flush()
		}
	},
}

func init() {
	serverCmd.AddCommand(serverListCmd)

	serverListCmd.Flags().Bool("no-headers", false, "do not print header")
}
