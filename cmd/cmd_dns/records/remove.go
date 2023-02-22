package cmd_dns_records

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zenion/infoblox-cli/pkg/infoblox"
)

var cmdRemove = &cobra.Command{
	Use:   "remove [fqdn]",
	Short: "remove a dns record from a zone",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ibClient := infoblox.New(
			viper.GetString("infoblox.host"),
			viper.GetString("infoblox.username"),
			viper.GetString("infoblox.password"),
		)
		_, err := ibClient.RemoveHostRecord(args[0], cmd.Flags().Lookup("view").Value.String())
		cobra.CheckErr(err)
		cmd.Println("Removed record: ", args[0])
	},
}

func init() {
	cmdRemove.Flags().StringP("view", "v", "Internal", "View to add the record to")
}
