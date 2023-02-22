package cmd_config

import (
	"os"

	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "config",
	Short: "config management",
}

func Execute() {
	err := Root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Root.AddCommand(cmdGet)
	Root.AddCommand(cmdSet)
}
