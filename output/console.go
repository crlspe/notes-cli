package output

import (
	"fmt"

	"github.com/crlspe/notes-cli-v4/model"
)

func PrintConsole(items model.ItemList) {
	for _, item := range items {
		fmt.Printf("Id: %.8s\t Content: %s\t Type: %s\t Completed: %t\t CreatedAt: %s\t CompletedAt: %s\n",
			item.Id[len(item.Id)-8:], item.Content, item.Type, item.Completed, item.CreatedAt, item.CompletedAt)
	}
}
