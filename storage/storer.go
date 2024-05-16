package storage

import "github.com/crlspe/notes-cli-v4/model"

type Storer interface {
	Load() model.ItemList
	Save(items model.ItemList)
}
