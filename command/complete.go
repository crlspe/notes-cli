package command

import (
	"fmt"

	"github.com/crlspe/notes-cli/model"
)

func SetTaskStatusAs(selectedItems model.ItemList, areCompleted bool) {
	var items = storer.Load()
	items.SetAsCompleted(selectedItems, areCompleted)
	storer.Save(items)
	fmt.Printf("Updated %v item(s) to complete: %v", len(selectedItems), areCompleted)
}
