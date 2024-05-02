package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/crlspe/notes-cli-v4/utils"
)

type ItemType string

const (
	NOTE ItemType = "note"
	TASK ItemType = "task"
)

type Item struct {
	Id          string   `json:"id"`
	Content     string   `json:"content"`
	Type        ItemType `json:"isTask"`
	Completed   bool     `json:"completed"`
	CompletedAt string   `json:"completedAt"`
	CreatedAt   string   `json:"createdAt"`
}

func (item *Item) Complete(completed bool) {
	item.Completed = completed
	item.CompletedAt = ""
	if completed {
		item.CompletedAt = time.Now().Format(constant.DateFormat)
	}
}

func NewNote(content string) Item {
	return NewItem(content, NOTE, false)
}

func NewTask(content string, completed bool) Item {
	return NewItem(content, TASK, completed)
}

func NewItem(
	content string,
	itemType ItemType,
	completed bool) Item {

	if itemType != TASK {
		itemType = NOTE
	}

	var id = utils.GenerateId()
	var createdAt = time.Now().Format(constant.DateFormat)

	var completedAt = ""
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

func (item Item) ContainsAll(searchTerms string) bool {
	var terms = strings.Split(searchTerms, "|")
	var containsAll = true
	for _, term := range terms {
		containsAll = containsAll && item.Contains(term)
	}
	return containsAll
}
