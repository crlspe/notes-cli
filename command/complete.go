package command

import (
	"github.com/crlspe/notes-cli-v4/model"
)

func SetTaskStatusAs(selectedItems model.ItemList, areCompleted bool) {
	var items = storer.Load() 
	items.SetAsCompleted(selectedItems, areCompleted)
	storer.Save(items)
}
