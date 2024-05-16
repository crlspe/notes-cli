package command

import (
	"github.com/crlspe/notes-cli-v4/model"
)

func RestoreItems(selectedItems model.ItemList) {
	var items = storer.Load()
	items.UpdateAsRemoved(selectedItems, false)
	storer.Save(items)
}
