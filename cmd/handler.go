package cmd

import "github.com/crlspe/notes-cli-v4/model"

func HandleCommands(flags model.CliFlags, items model.ItemList) {
	var selectedItems = items
	switch {
	case *flags.Add:
		AddItems(items, flags)
	case *flags.Search && !*flags.Add:
		selectedItems = SearchItems(items, flags)
		fallthrough
	default:
		if *flags.SetAsCompleted && !*flags.SetAsIncompleted {
			CompleteItems(items, selectedItems)
		}
		if *flags.SetAsIncompleted && !*flags.SetAsCompleted {
			IncompleteItems(items, selectedItems)
		}
	}
}
