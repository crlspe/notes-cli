package command

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/storage"
)

func RemoveItems(itemsToRemove model.ItemList, isPermanent bool) {
	var items = model.ItemList{}
	items.Load(storage.LoadJsonFile)

	if isPermanent {
		items.RemoveAll(itemsToRemove)
	} else {
		items.UpdateAsRemoved(itemsToRemove, true)
	}
	items.Save(storage.SaveJsonFile)
}
