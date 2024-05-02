package model

import (
	flag "github.com/spf13/pflag"
)

type CliFlags struct {
	Search           *bool
	Add              *bool
	IsTask           *bool
	IsNote           *bool
	SetAsCompleted   *bool
	SetAsIncompleted *bool
	StringArgs       []string
}

func (flags *CliFlags) InitilizeCliFlags() CliFlags {
	flags.Search = flag.BoolP("search", "s", true, "Search items.")
	flags.Add = flag.BoolP("add", "a", false, "Add elements to the list.")

	flags.SetAsCompleted = flag.BoolP("complete", "X", false, "Mark tasks as completed")
	flags.SetAsIncompleted = flag.BoolP("incomplete", "O", false, "Mark tasks as incompleted")

	flags.IsTask = flag.BoolP("task", "t", false, "Set type to task.")
	flags.IsNote = flag.BoolP("note", "n", false, "Set type to note.")

	flag.Parse()
	flags.StringArgs = flag.Args()

	return *flags
}
