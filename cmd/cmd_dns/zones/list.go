package cmd_dns_zones

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zenion/infoblox-cli/pkg/infoblox"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "list all dns zones",
	Run: func(cmd *cobra.Command, args []string) {
		ibClient := infoblox.New(
			viper.GetString("infoblox.host"),
			viper.GetString("infoblox.username"),
			viper.GetString("infoblox.password"),
		)
		zones, err := ibClient.GetZones()
		cobra.CheckErr(err)

		w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
		fmt.Fprintln(w, "DOMAIN\tVIEW\n---------\t---------")
		for _, zone := range zones {
			fmt.Fprintf(w, "%s\t%s\n", zone.Fqdn, zone.View)
		}
		w.Flush()
	},
}

func init() {
}
