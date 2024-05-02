package cmd

import (
	"github.com/crlspe/notes-cli-v4/input"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/output"
	"github.com/crlspe/notes-cli-v4/storage"
)

const ADD_PROMPT = "Add: "

func AddItems(items model.ItemList, flags model.CliFlags) {
	var contentValues = flags.StringArgs

	switch {
	case len(contentValues) <= 0:
		contentValues = input.MultiplePrompt(ADD_PROMPT)
		fallthrough
	default:
		if *flags.IsTask {
			items.AddTasks(contentValues)
		} else {
			items.AddNotes(contentValues)
		}
	}

	items.Save(storage.SaveJsonFile)
	items.Print(output.PrintShortTable)
}
