package parse

import "strings"

type GitLog struct {
	RawString string
}

func (g *GitLog) LastCommitHash() string {
	logLines := strings.Split(g.RawString, "\n")
	commitHash := strings.Split(logLines[0], " ")
	return commitHash[1]
}