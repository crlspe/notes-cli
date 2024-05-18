package command

import (
	"fmt"

	"github.com/crlspe/notes-cli/constant"
	"github.com/crlspe/notes-cli/input"
	"github.com/crlspe/notes-cli/model"
)

func AddItems(flags model.Flags) {

	var contentValues = GetAddInput(flags)

	var items = storer.Load()
	if *flags.IsTask {
		items.AddTasks(contentValues)
	} else {
		items.AddNotes(contentValues)
	}

	storer.Save(items)
	fmt.Printf("Added %v item(s)", len(contentValues))
}

func GetAddInput(flags model.Flags) []string {
	if len(flags.StringArgs) <= 0 {
		return input.MultiplePrompt(constant.PromptAdd)
	}
	return flags.StringArgs
}
