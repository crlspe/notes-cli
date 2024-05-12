package main

import (
	"github.com/crlspe/notes-cli-v4/command"
	"github.com/crlspe/notes-cli-v4/output"
)

func main() {
	output.ClearScreen()
	output.PrintApplicationInfo()
	command.Cli{}.Run()
}
