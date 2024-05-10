package command

import (
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/spf13/pflag"
)

type Cli struct {
	Flags model.Flags
}

func (cli *Cli) initilize() model.Flags {
	cli.Flags.Search = pflag.BoolP("search", "s", true, "Search items.")
	cli.Flags.Add = pflag.BoolP("add", "a", false, "Add item(s) to the list.")
	cli.Flags.Remove = pflag.Bool("remove", false, "Soft Remove item(s) from the list.")
	cli.Flags.IsPermanent = pflag.Bool("permanent", false, "Hard Remove item(s) from the list")
	cli.Flags.Restore = pflag.Bool("restore", false, "Restore sof deleted items")
	cli.Flags.ShowRemoved = pflag.Bool("showRemoved", false, "Include Removed items on results")
	cli.Flags.SetAsCompleted = pflag.BoolP("complete", "X", false, "Mark task(s) as completed")
	cli.Flags.SetAsIncompleted = pflag.BoolP("incomplete", "O", false, "Mark task(s) as incompleted")
	cli.Flags.IsTask = pflag.BoolP("task", "t", false, "Set type to task.")
	cli.Flags.IsNote = pflag.BoolP("note", "n", false, "Set type to note.")

	pflag.Parse()
	cli.Flags.StringArgs = pflag.Args()

	return cli.Flags
}

func (cli Cli) handleFlags() {
	var items = model.ItemList{}
	var selectedItems = model.ItemList{}

	switch {

	case *cli.Flags.Add:
		AddItems(cli.Flags)

	case *cli.Flags.Search && !*cli.Flags.Add:
		items, selectedItems = SearchItems(cli.Flags)
		fallthrough

	default:

		if *cli.Flags.SetAsCompleted != *cli.Flags.SetAsIncompleted {
			SetItemsAsCompleted(items, selectedItems, *cli.Flags.SetAsCompleted)
		}

		if *cli.Flags.Remove && !*cli.Flags.Restore {
			//AskforConfirmation
			RemoveItems(items, selectedItems, *cli.Flags.IsPermanent)
		}

		if *cli.Flags.Restore && !*cli.Flags.Remove {
			//Askfor Confirmation
			RestoreItems(items, selectedItems)
		}
	}
}

func (cli Cli) Run() {
	cli.initilize()
	cli.handleFlags()
}
