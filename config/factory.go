package config

import "github.com/crlspe/notes-cli/output"

func GetDependency(name string) interface{} {

	switch name {
	case "PrintShortTable":
		return output.PrintShortTable
	}

	return nil
}
