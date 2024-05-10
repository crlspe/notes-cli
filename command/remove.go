package command

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/storage"
)

func RemoveItems(items, itemsToRemove model.ItemList, isPermanent bool) {
	if isPermanent {
		items.RemoveAll(itemsToRemove)
	} else {
		items.UpdateAsRemoved(itemsToRemove, true)
	}
	items.Save(storage.SaveJsonFile)
}
