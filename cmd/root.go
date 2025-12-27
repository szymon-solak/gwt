package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gwt",
	Short: "Simplified git worktree management",
	Long:  "gwt simplifies git worktree workflows - initialize bare repos, create and manage worktrees across branches",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'gwt --help' for usage")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
