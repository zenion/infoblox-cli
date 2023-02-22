package cmd_dns_records

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zenion/infoblox-cli/pkg/infoblox"
)

var cmdList = &cobra.Command{
	Use:   "list [zone]",
	Short: "list all dns records for a zone",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ibClient := infoblox.New(
			viper.GetString("infoblox.host"),
			viper.GetString("infoblox.username"),
			viper.GetString("infoblox.password"),
		)
		records, err := ibClient.GetZoneRecords(args[0])
		cobra.CheckErr(err)

		w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
		fmt.Fprintln(w, "NAME\tIP\n---------\t---------")
		for _, record := range records {
			fmt.Fprintf(w, "%s\t%s\n", record.Name, record.Ipv4Addrs[0].Ipv4Addr)
		}
		w.Flush()
	},
}

func init() {
}
