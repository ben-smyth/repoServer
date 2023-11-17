package cmd

import (
	"ben-smyth/go-git-fun/git"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	RunE:  Run,
}

func Run(cmd *cobra.Command, args []string) error {

	_, fs, err := git.CloneRepo("https://github.com/kyma-project/cli.git")
	if err != nil {
		return err
	}

	err = git.ListFiles(fs, "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}