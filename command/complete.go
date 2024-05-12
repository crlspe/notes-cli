package command

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/storage"
)

func SetTaskStatusAs(selectedItems model.ItemList, areCompleted bool) {
	var items = model.ItemList{}
	items.Load(storage.LoadJsonFile)
	items.SetAsCompleted(selectedItems, areCompleted)
	items.Save(storage.SaveJsonFile)
}
