package parse

import "strings"

type GitLog struct {
	RawString string
}

// returns the short hash of the last commit
func (g *GitLog) LastCommitHash() string {
	logLines := strings.Split(g.RawString, "\n")
	commitHash := strings.Split(logLines[0], " ")
	return commitHash[1][:6]
}