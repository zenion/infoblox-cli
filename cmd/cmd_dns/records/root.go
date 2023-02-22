package cmd_dns_records

import (
	"os"

	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "records",
	Short: "dns records management",
}

func Execute() {
	err := Root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Root.AddCommand(cmdList)
	Root.AddCommand(cmdAdd)
	Root.AddCommand(cmdRemove)
}
