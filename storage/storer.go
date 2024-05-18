package storage

import "github.com/crlspe/notes-cli/model"

type Storer interface {
	Load() model.ItemList
	Save(model.ItemList)
}
