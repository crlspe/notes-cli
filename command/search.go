package command

import (
	"strings"

	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/crlspe/notes-cli-v4/input"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/output"
	"github.com/crlspe/notes-cli-v4/storage"
)

func SearchItems(flags model.Flags) (model.ItemList, model.ItemList) {
	var items = model.ItemList{}
	items.Load(storage.LoadJsonFile)

	if !*flags.ShowRemoved {
		items = items.Filter(func(x model.Item) bool { return !x.Removed })
	}

	var selectedItems = make(model.ItemsMap)

	// GET TYPE
	if *flags.IsTask {
		selectedItems = items.GetTasks().ToMap()
	} else if *flags.IsNote {
		selectedItems = items.GetNotes().ToMap()
	} else {
		selectedItems = items.ToMap()
	}

	// GET INPUT
	var searchTerm = constant.Empty
	if len(flags.StringArgs) <= 0 {
		searchTerm = input.SinglePrompt("search: ")
	} else {
		searchTerm = strings.Join(flags.StringArgs, constant.Space)
	}

	var itemsFound = selectedItems.Find(searchTerm).ToList()
	itemsFound.Print(output.PrintShortTable)

	return items, itemsFound
}
