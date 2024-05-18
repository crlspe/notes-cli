package main

import (
	"github.com/crlspe/notes-cli/command"
	"github.com/crlspe/notes-cli/output"
)

func main() {
	output.ClearScreen()
	output.PrintApplicationInfo()
	command.Cli{}.Run()
}
