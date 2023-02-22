package cmd_dns_records

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zenion/infoblox-cli/pkg/infoblox"
)

var cmdAdd = &cobra.Command{
	Use:   "add [fqdn] [ip]",
	Short: "add a dns record to a zone",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ibClient := infoblox.New(
			viper.GetString("infoblox.host"),
			viper.GetString("infoblox.username"),
			viper.GetString("infoblox.password"),
		)
		_, err := ibClient.AddHostRecord(args[0], args[1], cmd.Flags().Lookup("view").Value.String())
		cobra.CheckErr(err)
		cmd.Println("Added record: ", args[0], args[1])
	},
}

func init() {
	cmdAdd.Flags().StringP("view", "v", "Internal", "View to add the record to")
}
