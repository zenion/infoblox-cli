package cmd_dns

import (
	"os"

	"github.com/spf13/cobra"
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
	Root.AddCommand(cmdGet)
}
