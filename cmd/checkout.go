package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:               "checkout <branch>",
	Aliases:           []string{"co"},
	Short:             "Checkout an existing branch to a worktree",
	Long:              "Create a new worktree in branches/<branch> directory for the specified branch",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: branchCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		path := getWorktreePath(branch)

		gitCmd := exec.Command("git", "worktree", "add", path, branch)
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
	rootCmd.AddCommand(checkoutCmd)
}
