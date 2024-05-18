package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/crlspe/notes-cli/constant"
	"github.com/crlspe/notes-cli/model"
)

type JsonFile struct{}

func (f *JsonFile) Load() model.ItemList {
	var items = model.ItemList{}

	file, _ := os.OpenFile(constant.FilePath, os.O_RDONLY|os.O_CREATE, constant.FileRWPermissions)
	defer file.Close()

	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&items)

	return items
}

func (f *JsonFile) Save(items model.ItemList) {
	file, err := os.Create(constant.FilePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(items); err != nil {
		fmt.Println(err)
	}
}
