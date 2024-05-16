package command

import (
	"github.com/crlspe/notes-cli-v4/model"
)

func RemoveItems(itemsToRemove model.ItemList, isPermanent bool) {
	var items = storer.Load()

	if isPermanent {
		items.RemoveAll(itemsToRemove)
	} else {
		items.UpdateAsRemoved(itemsToRemove, true)
	}

	storer.Save(items)
}
