package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadv/oh-my-pg-linter/cmd/commands"
)

var (
	cfgFile string
	// AppVersion of cli.
	AppVersion = `devel`
)

func main() {
	cobra.OnInitialize(initConfig)
	rootCmd := &cobra.Command{
		Use:     "oh-my-pg-linter",
		Version: AppVersion,
	}
	rootCmd.PersistentFlags().StringP("rules", "r", ".", "Path to directory with rules.")
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(commands.Check())
	rootCmd.AddCommand(commands.Test())
	rootCmd.AddCommand(commands.Run())
	rootCmd.AddCommand(commands.TestAll())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".storage-inventory" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".oh-my-pg-linter")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}
