package model

import (
	"fmt"
	"os"

	"github.com/crlspe/notes-cli-v4/constant"
)

var HomeFolder, _ = os.UserHomeDir()
var FilePath = HomeFolder + constant.FileName

type ItemList []Item

func (items *ItemList) Load(loadData func() ItemList) ItemList {
	*items = loadData()
	return *items
}
func (items ItemList) Save(saveData func(ItemList)) {
	saveData(items)
}

func (items *ItemList) Add(newItem Item) {
	*items = append(*items, newItem)
}

func (items *ItemList) AddRange(newItems ItemList) {
	*items = append(*items, newItems...)
}

func (items *ItemList) AddNotes(notes []string) {
	for _, note := range notes {
		*items = append(*items, NewNote(note))
	}
}

func (items *ItemList) AddTasks(tasks []string) {
	for _, task := range tasks {
		*items = append(*items, NewTask(task, false))
	}
}

func (items ItemList) Print(output func(ItemList)) {
	output(items)
}

func (items *ItemList) Update(updatedItem Item) {
	var idx, _ = items.findById(updatedItem)
	if idx != -1 {
		(*items)[idx] = updatedItem
	}
	fmt.Println(fmt.Sprintln("Item ID:%s Not found!" + updatedItem.Id))
}

func (items *ItemList) setCompleteAll(selectedItems ItemList, isCompleted bool) ItemList {
	var itemsUpdated = ItemList{}
	list := *items
	for _, selectedItem := range selectedItems {
		var index, _ = items.findById(selectedItem)
		list[index].Complete(isCompleted)
		itemsUpdated = append(itemsUpdated, list[index])
	}
	return itemsUpdated
}

func (items ItemList) CompleteAll(selectedItems ItemList) ItemList {
	return items.setCompleteAll(selectedItems, true)
}

func (items ItemList) IncompleteAll(selectedItems ItemList) ItemList {
	return items.setCompleteAll(selectedItems, false)
}

func (items ItemList) findById(sItem Item) (int, Item) {
	for idx, item := range items {
		if sItem.Id == item.Id {
			return idx, item
		}
	}
	return -1, Item{}
}

func (items ItemList) filter(filterFn func(Item) bool) ItemList {
	var result = ItemList{}
	for _, item := range items {
		if filterFn(item) {
			result = append(result, item)
		}
	}
	return result
}

func (items ItemList) Map(mapFn func(Item) interface{}) []interface{} {
	var result = make([]interface{}, len(items))
	for idx, item := range items {
		result[idx] = mapFn(item)
	}
	return result
}

func (items ItemList) getByType(itemType ItemType) ItemList {
	return items.filter(func(x Item) bool {
		return x.Type == itemType
	})
}

func (items ItemList) GetTasks() ItemList {
	return items.getByType(TASK)
}

func (items ItemList) GetNotes() ItemList {
	return items.getByType(NOTE)
}

func (items ItemList) Find(search string) ItemList {
	return items.filter(func(x Item) bool { return x.Contains(search) })
}

func (items ItemList) ToMap() map[string]Item {
	var mapItems = make(map[string]Item)
	for _, item := range items {
		mapItems[item.Id] = item
	}
	return mapItems
}
