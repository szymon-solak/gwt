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

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return filepath.Join("branches", dirName)
	}

	gitRoot := strings.TrimSpace(string(output))
	return filepath.Join(gitRoot, "branches", dirName)
}
