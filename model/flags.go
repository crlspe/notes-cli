package model

type Flags struct {
	Search           *bool
	Add              *bool
	Remove           *bool
	IsTask           *bool
	IsNote           *bool
	IsPermanent      *bool
	Restore          *bool
	SetAsCompleted   *bool
	SetAsIncompleted *bool
	ShowRemoved      *bool
	StringArgs       []string
}

type Flag struct {
	Name      string
	Alias     string
	IsCommand bool
	Value     interface{}
}

type FlagList []Flag
