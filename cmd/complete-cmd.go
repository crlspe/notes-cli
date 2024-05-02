package cmd

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/output"
	"github.com/crlspe/notes-cli-v4/storage"
)

func CompleteItems(items model.ItemList, selectedItems model.ItemList) {
	var itemsUpdated = items.CompleteAll(selectedItems)
	itemsUpdated.Print(output.PrintShortTable)
	items.Save(storage.SaveJsonFile)
}

func IncompleteItems(items model.ItemList, selectedItems model.ItemList) {
	var itemsUpdated = items.IncompleteAll(selectedItems)
	itemsUpdated.Print(output.PrintShortTable)
	items.Save(storage.SaveJsonFile)
}
