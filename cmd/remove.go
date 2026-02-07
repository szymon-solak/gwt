package cmd

import (
	"fmt"
	"os"
	"os/exec"

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
		path := getWorktreePath(branch)

		cmdArgs := []string{"worktree", "remove"}

		force, _ := cmd.Flags().GetBool("force")
		if force {
			cmdArgs = append(cmdArgs, "--force")
		}

		cmdArgs = append(cmdArgs, path)

		gitCmd := exec.Command("git", cmdArgs...)
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
	removeCmd.Flags().BoolP("force", "f", false, "Force removal even if worktree is dirty or locked")
}
