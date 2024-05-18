package command

import (
	"strings"

	"github.com/crlspe/notes-cli/constant"
	"github.com/crlspe/notes-cli/input"
	"github.com/crlspe/notes-cli/model"
)

func SearchItems(flags model.Flags) model.ItemList {
	var items = storer.Load()
	applyFlagFilters(&items, flags)
	applyTypeFilters(&items, flags)

	var selectedItems model.ItemsMap = items.ToMap()
	var itemsFound = selectedItems.Find(GetSearchInput(flags)).ToList()

	return itemsFound
}

func applyFlagFilters(items *model.ItemList, flags model.Flags) {
	switch {
	case *flags.Restore:
		fallthrough
	case *flags.ShowRemoved:
		*items = items.Filter(func(x model.Item) bool { return x.Removed })
	default:
		*items = items.Filter(func(x model.Item) bool { return !x.Removed })
	}
}

func applyTypeFilters(items *model.ItemList, flags model.Flags) {
	switch {
	case *flags.IsTask:
		*items = items.GetTasks()
	case *flags.IsNote:
		*items = items.GetNotes()
	}
}

func GetSearchInput(flags model.Flags) string {
	switch {
	case len(flags.StringArgs) <= 0:
		return input.SinglePrompt(constant.PromptSearch)
	default:
		return strings.Join(flags.StringArgs, constant.StrSpace)
	}
}
