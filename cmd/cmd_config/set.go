package cmd_config

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdSet = &cobra.Command{
	Use:   "set",
	Short: "set config",
	Run: func(cmd *cobra.Command, args []string) {

		prompt := &promptui.Prompt{
			Label: "Infoblox Host",
			Validate: func(input string) error {
				if input == "" {
					return errors.New("hostname can't be empty")
				}
				return nil
			},
		}
		promptStr, err := prompt.Run()
		cobra.CheckErr(err)
		viper.Set("infoblox.host", promptStr)

		prompt = &promptui.Prompt{
			Label: "Infoblox Username",
			Validate: func(input string) error {
				if input == "" {
					return errors.New("username can't be empty")
				}
				return nil
			},
		}
		promptStr, err = prompt.Run()
		cobra.CheckErr(err)
		viper.Set("infoblox.username", promptStr)

		prompt = &promptui.Prompt{
			Label: "Infoblox Password",
			Mask:  '*',
			Validate: func(input string) error {
				if input == "" {
					return errors.New("password can't be empty")
				}
				return nil
			},
		}
		promptStr, err = prompt.Run()
		cobra.CheckErr(err)
		viper.Set("infoblox.password", promptStr)

		err = viper.WriteConfig()
		cobra.CheckErr(err)

		cmd.Println("Config saved successfully")
	},
}

func init() {
}
