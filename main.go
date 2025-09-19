package main

import (
	"fmt"

	"github.com/kdawg500/larry/parse"
)

func main() {
	logs := parse.GitLogs{}
	tags := parse.GitTags{}
	commit := logs.LastCommitHash()
	tag := tags.LastTag()
	tag.IncrementPatch()
	
	fmt.Println(tag.String())
	fmt.Println(commit)
}
