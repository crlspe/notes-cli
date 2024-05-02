package model

import "strings"

type ItemsMap map[string]Item

func (itemsMap *ItemsMap) exists(item Item) bool {
	var items = *itemsMap
	var _, exists = items[item.Id]
	return exists
}

func (itemsMap *ItemsMap) Add(item Item) {
	var items = *itemsMap
	if !items.exists(item) {
		items[item.Id] = item
	}
}

// Use Generics here
func (itemsMap ItemsMap) AddList(items ItemList) {
	for _, item := range items {
		itemsMap.Add(item)
	}
}

func (itemsMap ItemsMap) AddMap(items ItemsMap) {
	for _, item := range items {
		itemsMap.Add(item)
	}
}

func (itemsMap *ItemsMap) Remove(item Item) {
	var items = *itemsMap
	if items.exists(item) {
		delete(items, item.Id)
	}
}

func (itemsMap ItemsMap) ToList() ItemList {
	var items ItemList
	for _, item := range itemsMap {
		items = append(items, item)
	}
	return items
}

func (itemsMap ItemsMap) FindAll(searchText string) ItemsMap {

	if len(strings.TrimSpace(searchText)) <= 0 {
		return itemsMap
	}

	var result = make(ItemsMap)
	var searchTerms = strings.Split(searchText, " ")

	for _, searchTerm := range searchTerms {
		if strings.Contains(searchTerm, "|") {
			var found = itemsMap.Filter(func(x Item) bool { return x.ContainsAll(searchTerm) })
			result.AddMap(found)
		} else {
			result.AddMap(itemsMap.Contains(searchTerm))
		}
	}

	return result
}

func (itemsMap ItemsMap) Contains(searchTerm string) ItemsMap {
	return itemsMap.Filter(func(x Item) bool { return x.Contains(searchTerm) })
}

func (itemsMap ItemsMap) Map(mapFn func(Item) interface{}) map[string]interface{} {
	var result = make(map[string]interface{}, len(itemsMap))
	for key, item := range itemsMap {
		result[key] = mapFn(item)
	}
	return result
}

func (itemsMap ItemsMap) Filter(filterFn func(Item) bool) ItemsMap {
	var result = make(ItemsMap)
	for _, item := range itemsMap {
		if filterFn(item) {
			result[item.Id] = item
		}
	}
	return result
}
