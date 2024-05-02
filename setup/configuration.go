package setup

import (
	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/output"
	"github.com/crlspe/notes-cli-v4/storage"
)

type Options struct {
	ApplicationName string
	Version         string
	Storage         storage.IStorage
	PrintOutput     func(items model.ItemList)
}

type Configuration struct {
	Options
}

type optionSetup func(conf *Options)

func NewConfig(setupFns ...optionSetup) *Configuration {
	var configuration = defaultConfiguration()
	for _, setup := range setupFns {
		setup(&configuration)
	}
	return &Configuration{
		configuration,
	}
}

func defaultConfiguration() Options {
	return Options{
		ApplicationName: constant.ApplicationName,
		Version:         constant.Version,
		Storage:         storage.JsonFile{},
		PrintOutput:     output.PrintShortTable,
	}
}
