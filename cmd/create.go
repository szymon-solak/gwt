package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:               "create <branch> [base-branch]",
	Aliases:           []string{"add"},
	Short:             "Create a new branch with a worktree",
	Long:              "Create a new branch and worktree in branches/<branch> directory",
	Args:              cobra.RangeArgs(1, 2),
	ValidArgsFunction: createCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		dirName := sanitizeBranchName(branch)
		path := filepath.Join("branches", dirName)

		gitArgs := []string{"worktree", "add", "-b", branch, path}
		if len(args) == 2 {
			gitArgs = append(gitArgs, args[1])
		}

		gitCmd := exec.Command("git", gitArgs...)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr

		if err := gitCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		absPath, _ := filepath.Abs(path)
		fmt.Printf("\nWorktree created at: %s\n", absPath)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
