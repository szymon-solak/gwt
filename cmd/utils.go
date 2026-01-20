package cmd

import "regexp"

var ticketPattern = regexp.MustCompile(`^([A-Z]+-\d+)`)

func sanitizeBranchName(branch string) string {
	if matches := ticketPattern.FindStringSubmatch(branch); matches != nil {
		return matches[1]
	}
	if len(branch) > 20 {
		return branch[:20]
	}
	return branch
}
