package lib

import (
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CallGit(args ...string) {
	dotfilesDir := viper.GetString("directory")

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	cmd := exec.Command("git", append([]string{"--git-dir", dotfilesDir, "--work-tree", home}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	err = cmd.Run()
	cobra.CheckErr(err)
}

func VerifyGitRepository() {
	dotfilesDir := viper.GetString("directory")

	_, err := git.PlainOpen(dotfilesDir)
	cobra.CheckErr(err)
}
