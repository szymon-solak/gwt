package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func getBranches() []string {
	cmd := exec.Command("git", "branch", "-a", "--format=%(refname:short)")
	output, err := cmd.Output()
	if err != nil {
		return []string{}
	}

	branches := strings.Split(strings.TrimSpace(string(output)), "\n")
	var result []string
	for _, b := range branches {
		if b != "" && !strings.HasPrefix(b, "origin/HEAD") {
			result = append(result, b)
		}
	}
	return result
}

func getWorktreeDirs() []string {
	entries, err := os.ReadDir("branches")
	if err != nil {
		return []string{}
	}

	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e.Name())
		}
	}
	return dirs
}

func branchCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return getBranches(), cobra.ShellCompDirectiveNoFileComp
}

// worktreeDirCompletion provides completion for worktree directories
func worktreeDirCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return getWorktreeDirs(), cobra.ShellCompDirectiveNoFileComp
}

// createCompletion provides completion for create command (only base-branch argument)
func createCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		// First argument is new branch name - no completion
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	// Second argument is base branch
	return getBranches(), cobra.ShellCompDirectiveNoFileComp
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `Generate shell completion script for gwt.

To load completions:

Bash:
  $ source <(gwt completion bash)
  # To load for each session, add to ~/.bashrc:
  $ gwt completion bash > /etc/bash_completion.d/gwt

Zsh:
  # If shell completion is not enabled, enable it:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  $ source <(gwt completion zsh)
  # To load for each session, add to ~/.zshrc:
  $ gwt completion zsh > "${fpath[1]}/_gwt"

Fish:
  $ gwt completion fish | source
  # To load for each session:
  $ gwt completion fish > ~/.config/fish/completions/gwt.fish

PowerShell:
  PS> gwt completion powershell | Out-String | Invoke-Expression
  # To load for each session, add to profile:
  PS> gwt completion powershell > gwt.ps1
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
