package cmd

import (
	"strings"

	"github.com/crlspe/notes-cli-v4/input"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/output"
)

func SearchItems(items model.ItemList, flags model.CliFlags) model.ItemList {
	var selectedItems = make(model.ItemsMap)
	if *flags.IsTask {
		selectedItems = items.GetTasks().ToMap()
	} else if *flags.IsNote {
		selectedItems = items.GetNotes().ToMap()
	} else {
		selectedItems = items.ToMap()
	}

	var searchTerm = ""
	if len(flags.StringArgs) <= 0 {
		searchTerm = input.SinglePrompt("search: ")
	} else {
		searchTerm = strings.Join(flags.StringArgs, " ")
	}

	var itemsFound = selectedItems.FindAll(searchTerm).ToList()
	itemsFound.Print(output.PrintShortTable)
	return itemsFound
}
