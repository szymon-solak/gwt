package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <repository-url> [destination]",
	Short: "Initialize a bare git repository from a URL",
	Long:  "Clone a git repository as a bare repo, suitable for use with git worktrees",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		var dest string
		if len(args) == 2 {
			dest = args[1]
		} else {
			// Extract repo name from URL
			parts := strings.Split(url, "/")
			repoName := parts[len(parts)-1]
			repoName = strings.TrimSuffix(repoName, ".git")
			dest = repoName + ".git"
		}

		gitCmd := exec.Command("git", "clone", "--bare", url, dest)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr

		if err := gitCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		absPath, _ := filepath.Abs(dest)
		fmt.Printf("\nBare repository created at: %s\n", absPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
