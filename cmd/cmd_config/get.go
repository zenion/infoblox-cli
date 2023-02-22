package cmd_config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "get config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nRun 'infoblox config set' to update your configuration")
		fmt.Println("\nInfoblox Host:\t\t", viper.GetString("infoblox.host"))
		fmt.Println("Infoblox Username:\t", viper.GetString("infoblox.username"))
	},
}

func init() {
}
