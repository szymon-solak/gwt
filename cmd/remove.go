package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:               "remove <branch>",
	Aliases:           []string{"rm"},
	Short:             "Remove a worktree",
	Long:              "Remove the worktree for the specified branch from branches/<branch> directory",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: worktreeDirCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		path := filepath.Join("branches", branch)

		gitCmd := exec.Command("git", "worktree", "remove", path)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr

		if err := gitCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nWorktree removed: %s\n", path)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
