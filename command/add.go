package command

import (
	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/crlspe/notes-cli-v4/input"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/output"
)


func AddItems(flags model.Flags) {
	var items =	storer.Load()

	var contentValues = flags.StringArgs

	// GET TYPE AND INPUT
	switch {
	case len(contentValues) <= 0:
		contentValues = input.MultiplePrompt(constant.AddPrompt)
		fallthrough
	default:
		if *flags.IsTask {
			items.AddTasks(contentValues)
		} else {
			items.AddNotes(contentValues)
		}
	}

	storer.Save(items)
	items.Print(output.PrintShortTable)
}
