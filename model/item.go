package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/crlspe/notes-cli/constant"
	"github.com/crlspe/notes-cli/input"
)

type ItemType string

const (
	NOTE ItemType = "note"
	TASK ItemType = "task"
)

type Item struct {
	Id          string   `json:"id"`
	Content     string   `json:"content"`
	Type        ItemType `json:"type"`
	Completed   bool     `json:"completed"`
	CompletedAt string   `json:"completedAt"`
	CreatedAt   string   `json:"createdAt"`
	Removed     bool     `json:"deleted"`
	RemovedAt   string   `json:"deletedAt"`
}

func NewNote(content string) Item {
	return NewItem(content, NOTE, false)
}

func NewTask(content string, completed bool) Item {
	return NewItem(content, TASK, completed)
}

func NewItem(content string, itemType ItemType, completed bool) Item {

	if itemType != TASK {
		itemType = NOTE
	}

	var id = input.GenerateId()
	var createdAt = time.Now().Format(constant.DateFormat)

	var completedAt = constant.StrEmpty
	if completed {
		completedAt = time.Now().Format(constant.DateFormat)
	}

	return Item{
		Id:          id,
		Content:     content,
		Type:        itemType,
		Completed:   completed,
		CompletedAt: completedAt,
		CreatedAt:   createdAt,
	}
}

func (item Item) Print() {
	fmt.Printf("Id: %.8s\t Content: %s\t Type: %s\t Completed: %t\t CreatedAt: %s\t CompletedAt: %s\n",
		item.Id[len(item.Id)-8:], item.Content, item.Type, item.Completed, item.CreatedAt, item.CompletedAt)
}

func (item Item) Contains(searchTerm string) bool {
	return strings.Contains(strings.ToLower(item.Content), strings.ToLower(searchTerm))
}

func (item Item) MustContainAll(searchTerms string, separator string) bool {
	var terms = strings.Split(searchTerms, separator)
	var containsAll = true
	for _, term := range terms {
		containsAll = containsAll && item.Contains(term)
	}
	return containsAll
}

func (item *Item) Complete(completed bool) {
	item.Completed = completed
	item.CompletedAt = constant.StrEmpty
	if completed {
		item.CompletedAt = time.Now().Format(constant.DateFormat)
	}
}

func (item *Item) SetAsRemoved(isDeleted bool) {
	item.Removed = isDeleted
	item.RemovedAt = constant.StrEmpty
	if isDeleted {
		item.RemovedAt = time.Now().Format(constant.DateFormat)
	}
}
