package storage

import "github.com/crlspe/notes-cli-v4/model"

type IStorage interface {
	Load() model.ItemList
	Save(model.ItemList)
}
