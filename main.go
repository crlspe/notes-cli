package main

import (
	"github.com/crlspe/notes-cli-v4/cmd"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/setup"
	"github.com/crlspe/notes-cli-v4/utils"
)

func main() {
	var config = setup.NewConfig()

	var items = model.ItemList{}
	items.Load(config.Storage.Load)

	var flags = model.CliFlags{}
	flags.InitilizeCliFlags()

	utils.ClearScreen()
	cmd.HandleCommands(flags, items)
}
