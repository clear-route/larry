package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd:= exec.Command("git", "log")
	output, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(output))
}