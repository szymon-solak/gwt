package cmd

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var ticketPattern = regexp.MustCompile(`^([A-Z]+-\d+)`)

func formatBranchName(branch string) string {
	if matches := ticketPattern.FindStringSubmatch(branch); matches != nil {
		return matches[1]
	}
	if len(branch) > 20 {
		return branch[:20]
	}
	return branch
}

func getWorktreePath(branch string) string {
	dirName := formatBranchName(branch)

	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return filepath.Join("branches", dirName)
	}

	// First line is always "worktree <path>" for the main/bare repo
	firstLine := strings.SplitN(strings.TrimSpace(string(output)), "\n", 2)[0]
	mainRoot := strings.TrimPrefix(firstLine, "worktree ")
	return filepath.Join(mainRoot, "branches", dirName)
}
