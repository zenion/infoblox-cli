package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zenion/infoblox-cli/cmd/cmd_config"
	"github.com/zenion/infoblox-cli/cmd/cmd_dns"
)

var rootCmd = &cobra.Command{
	Use:   "infoblox",
	Short: "Infoblox cli tool for managing things and stuff",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configDir := filepath.Join(home, ".config", "infoblox-cli")
	err = os.MkdirAll(configDir, os.ModePerm)
	cobra.CheckErr(err)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("IBLOX")

	if err := viper.ReadInConfig(); err != nil {
		// create config file with empty values if no config file found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.Set("infoblox.host", "")
			viper.Set("infoblox.username", "")
			viper.Set("infoblox.password", "")

			err = viper.WriteConfigAs(filepath.Join(configDir, "config.yaml"))
			cobra.CheckErr(err)
		} else {
			// Config file was found but another error was produced
			cobra.CheckErr(err)
		}
	}

	rootCmd.AddCommand(cmd_dns.Root)
	rootCmd.AddCommand(cmd_config.Root)
}
