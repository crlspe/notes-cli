package command

import (
	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/crlspe/notes-cli-v4/input"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/output"
	"github.com/crlspe/notes-cli-v4/storage"
)

func AddItems(flags model.Flags) {
	var items = model.ItemList{}
	items.Load(storage.LoadJsonFile)

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

	items.Save(storage.SaveJsonFile)
	items.Print(output.PrintShortTable)
}
