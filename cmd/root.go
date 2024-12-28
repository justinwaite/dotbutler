/*
Copyright © 2024 Justin Waite <justindwaite@gmail.com>
*/
package cmd

import (
	"fmt"
	"justinwaite/dotbutler/lib"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgDir string

	rootCmd = &cobra.Command{
		Use:   "dotbutler",
		Args:  cobra.ArbitraryArgs,
		Short: "Backup and restore your dotfiles",
		Long: `DotButler is a CLI tool to backup and restore your dotfiles.
It is designed to be simple and easy to use.`,
		DisableFlagParsing: true,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			lib.VerifyGitRepository()
			lib.CallGit(args...)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	rootCmd.PersistentFlags().StringVar(&cfgDir, "directory", filepath.Join(home, ".butler"), "directory of the repository (default is $HOME/.butler/)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// if cfgDir != "" {
	// 	viper.AddConfigPath(cfgDir)
	// } else {
	// 	home, err := os.UserHomeDir()
	// 	cobra.CheckErr(err)
	//
	// 	viper.AddConfigPath(filepath.Join(home, ".butler"))
	// }

	// viper.SetConfigType("yaml")
	// viper.SetConfigName(".butler")

	viper.Set("directory", cfgDir)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
