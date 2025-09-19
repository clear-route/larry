package main

import (
	"fmt"
	"os/exec"

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
	
	cmd= exec.Command("git", "tag")
	bTags, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	tag := parse.LastTag(bTags)
	

	fmt.Println(tag)
	fmt.Println(commit)
}