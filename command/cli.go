package command

import (
	"github.com/crlspe/notes-cli/input"
	"github.com/crlspe/notes-cli/model"
	"github.com/crlspe/notes-cli/output"
	"github.com/crlspe/notes-cli/storage"
	"github.com/spf13/pflag"
)

const confirmRestore = "Are you sure you want to RESTORE these items Yes/No? "
const confirmRemove = "Are you sure you want to REMOVE these items Yes/No? "
const confirmPermanentRemove = "Are you sure you want to permanently REMOVE these items Yes/No? "

var storer = storage.JsonFile{}

type Cli struct {
	Flags model.Flags
}

func (cli *Cli) initilize() model.Flags {
	cli.Flags.Search = pflag.BoolP("search", "s", true, "Search items.")
	cli.Flags.Add = pflag.BoolP("add", "a", false, "Add item(s) to the list.")
	cli.Flags.SetAsCompleted = pflag.BoolP("complete", "X", false, "Mark task(s) as completed")
	cli.Flags.SetAsIncompleted = pflag.BoolP("incomplete", "O", false, "Mark task(s) as incompleted")
	cli.Flags.IsTask = pflag.BoolP("task", "t", false, "Set type to task.")
	cli.Flags.IsNote = pflag.BoolP("note", "n", false, "Set type to note.")
	cli.Flags.ShowRemoved = pflag.Bool("showRemoved", false, "Include Removed items on results")
	cli.Flags.Remove = pflag.Bool("remove", false, "Soft Remove item(s) from the list.")
	cli.Flags.IsPermanent = pflag.Bool("permanent", false, "Hard Remove item(s) from the list")
	cli.Flags.Restore = pflag.Bool("restore", false, "Restore soft removed items")

	pflag.Parse()
	cli.Flags.StringArgs = pflag.Args()

	return cli.Flags
}

func (cli Cli) handleFlags() {
	var selectedItems = model.ItemList{}

	switch {

	case IsAddItems(cli.Flags):
		AddItems(cli.Flags)

	case IsSearchItems(cli.Flags):
		selectedItems = SearchItems(cli.Flags)
		fallthrough

	default:

		output.PrintShortTable(selectedItems)

		if IsTaskStatusChange(cli.Flags) {
			SetTaskStatusAs(selectedItems, *cli.Flags.SetAsCompleted)
		}

		if IsRemoveItems(cli.Flags) {
			RemoveItems(selectedItems, *cli.Flags.IsPermanent)
		}

		if IsRestoreItems(cli.Flags) {
			RestoreItems(selectedItems)
		}
	}
}

func (cli Cli) Run() {
	cli.initilize()
	cli.handleFlags()
}

func IsAddItems(flags model.Flags) bool {
	return *flags.Add
}

func IsSearchItems(flags model.Flags) bool {
	return *flags.Search && !*flags.Add
}

func IsTaskStatusChange(flags model.Flags) bool {
	return *flags.SetAsCompleted != *flags.SetAsIncompleted
}

func IsRemoveItems(flags model.Flags) bool {
	return IsRemove(flags) && (IsNotPermanent(flags) || IsPermanent(flags))
}

func IsRestoreItems(flags model.Flags) bool {
	return *flags.Restore && !*flags.Remove && input.YesNoPrompt(confirmRestore)
}

func IsRemove(flags model.Flags) bool {
	return *flags.Remove && !*flags.Restore
}

func IsNotPermanent(flags model.Flags) bool {
	return (!*flags.IsPermanent && input.YesNoPrompt(confirmRemove))
}

func IsPermanent(flags model.Flags) bool {
	return (*flags.IsPermanent && input.YesNoPrompt(confirmPermanentRemove))
}
