package command

import (
	"fmt"

	"github.com/crlspe/notes-cli/model"
)

func RestoreItems(selectedItems model.ItemList) {
	var items = storer.Load()
	items.UpdateAsRemoved(selectedItems, false)
	storer.Save(items)
	fmt.Printf("Restored %v item(s)", len(selectedItems))
}
