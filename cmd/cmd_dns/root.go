package cmd_dns

import (
	"os"

	"github.com/spf13/cobra"
	cmd_dns_records "github.com/zenion/infoblox-cli/cmd/cmd_dns/records"
	cmd_dns_zones "github.com/zenion/infoblox-cli/cmd/cmd_dns/zones"
)

var Root = &cobra.Command{
	Use:   "dns",
	Short: "dns record management",
}

func Execute() {
	err := Root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Root.AddCommand(cmd_dns_zones.Root)
	Root.AddCommand(cmd_dns_records.Root)
}
