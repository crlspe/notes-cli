package setup

import (
	"github.com/crlspe/notes-cli/config"
	"github.com/crlspe/notes-cli/constant"
	"github.com/crlspe/notes-cli/storage"
)

type Options struct {
	ApplicationName string
	Version         string
	Storage         storage.Storer
	StorageString   string
	PrintOutput     interface{}
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
		Storage:         &storage.JsonFile{},
		StorageString:   constant.HomeFolder + constant.FileName,
		PrintOutput:     config.GetDependency("PrintShortTable"),
	}
}
