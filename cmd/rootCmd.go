package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig, func() {
		if viper.GetBool("experimental") {
			addExperimentalCommands()
		}
	})
}

var rootCmd = &cobra.Command{
	Use:   "cobra-1176",
	Short: "https://github.com/spf13/cobra/issues/1176",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("experimental: " + viper.GetString("experimental"))
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func addExperimentalCommands() {

	fmt.Println("enabling experimental commands")
	rootCmd.AddCommand(testCmd)
}

func initConfig() {
	fmt.Println("initializing config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(("."))
	viper.SetDefault("experimental", false)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
