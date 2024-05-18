package command

import (
	"fmt"

	"github.com/crlspe/notes-cli/model"
)

func RemoveItems(itemsToRemove model.ItemList, isPermanent bool) {
	var items = storer.Load()

	if isPermanent {
		items.RemoveAll(itemsToRemove)
	} else {
		items.UpdateAsRemoved(itemsToRemove, true)
	}

	storer.Save(items)
	fmt.Printf("Deleted %v item(s), permanently: %v", len(itemsToRemove), isPermanent)
}
