package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kdawg500/larry/parse"
)

func main() {
	cmd:= exec.Command("git", "log")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	gitLog := parse.GitLog{RawString: string(output)}
	commit := gitLog.LastCommitHash()
	var b strings.Builder
	b.Write(output)
	fmt.Println(commit)
	// fmt.Printf("`%v`", string(output))
}