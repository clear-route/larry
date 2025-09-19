package parse

import (
	"os/exec"
	"strings"
)

// struct for the git logs
type GitLogs struct {
	rawLogs []byte
}

// gets the logs from the git repository
func (g *GitLogs) getLogs() {
	cmd := exec.Command("git", "log")
	bLogs, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	g.rawLogs = bLogs
}

// returns the short hash of the last commit
func (g *GitLogs) LastCommitHash() string {
	g.getLogs()
	logLines := strings.Split(string(g.rawLogs), "\n")
	commitHash := strings.Split(logLines[0], " ")
	return commitHash[1][:6]
}
