package cmd_dns

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zenion/infoblox-cli/pkg/infoblox"
)

var cmdGet = &cobra.Command{
	Use:   "get [record]",
	Short: "get dns record",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		host := viper.GetString("infoblox.host")
		username := viper.GetString("infoblox.username")
		password := viper.GetString("infoblox.password")
		ibclient, err := infoblox.NewClient(host, username, password)
		cobra.CheckErr(err)

		res, err := ibclient.GetARecordByRef(args[0])
		cobra.CheckErr(err)
		infoblox.MarshalAndPrint(res)
	},
}

func init() {
}
