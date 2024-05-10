package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/crlspe/notes-cli-v4/model"
)

var HomeFolder, _ = os.UserHomeDir()
var FilePath = HomeFolder + constant.FileName

type JsonFile struct{}

func (f JsonFile) Load() model.ItemList {
	return LoadJsonFile()
}

func (f JsonFile) Save(items model.ItemList) {
	SaveJsonFile(items)
}

func LoadJsonFile() model.ItemList {
	var items = model.ItemList{}

	file, _ := os.OpenFile(FilePath, os.O_RDONLY|os.O_CREATE, constant.FileRWPermissions)
	defer file.Close()

	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&items)

	return items
}

func SaveJsonFile(items model.ItemList) {
	file, err := os.Create(FilePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(items); err != nil {
		fmt.Println(err)
	}
}
