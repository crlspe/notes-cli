package command

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/storage"
)

func RestoreItems(selectedItems model.ItemList) {
	var items = model.ItemList{}
	items.Load(storage.LoadJsonFile)
	items.UpdateAsRemoved(selectedItems, false)
	items.Save(storage.SaveJsonFile)
}
