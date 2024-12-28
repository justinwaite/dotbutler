package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize dotbutler",
	Long:  `Initalizes dotbulter by creating a detached git repository in the user's home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitDir := viper.GetString("directory")

		if _, err := os.Stat(gitDir); os.IsNotExist(err) {
			fmt.Println("Creating directory", gitDir)
			err := os.Mkdir(gitDir, 0755)
			cobra.CheckErr(err)

			fmt.Println("Initializing git repository")

			_, err = git.PlainInit(gitDir, true)
			cobra.CheckErr(err)

			r, err := git.PlainOpen(gitDir)
			cobra.CheckErr(err)

			cfg, err := r.Config()
			cobra.CheckErr(err)

			cfg.Raw.SetOption("status", "", "showUntrackedFiles", "no")
			err = r.SetConfig(cfg)
			cobra.CheckErr(err)
		} else {
			fmt.Println("Creating directory", gitDir, "[skipped]")
		}

		fmt.Println("dotbutler has been initialized")
	},
}
