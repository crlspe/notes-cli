package command

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/storage"
)

func SetItemsAsCompleted(items model.ItemList, selectedItems model.ItemList, areCompleted bool) {
	items.SetAsCompleted(selectedItems, areCompleted)
	items.Save(storage.SaveJsonFile)
}
