package cmd_dns_zones

import (
	"os"

	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "zones",
	Short: "dns zones management",
}

func Execute() {
	err := Root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Root.AddCommand(cmdList)
}
