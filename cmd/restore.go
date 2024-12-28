package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(restoreCommand)
}

var restoreCommand = &cobra.Command{
	Use:   "restore [repositoryUrl]",
	Args:  cobra.ExactArgs(1),
	Short: "Restores a dotfiles repository onto the current machine",
	Long:  `Initalizes dotbulter by creating a detached git repository in the user's home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		dotfilesDir := viper.GetString("directory")
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		command := exec.Command("git", "clone", "--separate-git-dir", dotfilesDir, args[0], home)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin
		command.Env = os.Environ()
	},
}
