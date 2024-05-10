package command

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/storage"
)

func RestoreItems(items, selectedItems model.ItemList) {
	items.UpdateAsRemoved(selectedItems, false)
	items.Save(storage.SaveJsonFile)
}
