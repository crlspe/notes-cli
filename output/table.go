package output

import (
	"fmt"

	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/crlspe/notes-cli-v4/model"
	"github.com/crlspe/notes-cli-v4/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var CompletedColor = text.Colors{text.FgGreen}
var IncompletedColor = text.Colors{text.FgRed}
var NoteColor = text.Colors{text.FgWhite}
var TaskColor = text.Colors{text.FgCyan}

func PrintShortTable(items model.ItemList) {
	itemTable := InitializeTableSettings(items)

	itemTable.AppendHeader(table.Row{
		" ",
		utils.Blue("TYPE"),
		utils.Blue("CONTENT"),
		utils.Blue("SCOPES"),
	})

	for _, item := range items {
		itemTable.AppendRow(table.Row{
			status(item),
			item.Type,
			utils.FormatContent(item.Content, 60),
			utils.FormatScopes(item.Content, 20),
		})
	}

	fmt.Println(itemTable.Render())
	fmt.Println()
}

func InitializeTableSettings(items model.ItemList) table.Writer {
	var itemTable = table.NewWriter()
	itemTable.SetStyle(table.StyleRounded)
	itemTable.SetTitle(
		fmt.Sprintf("%v %v \nTotal: %v | Notes: %v | Tasks: %v",
			utils.Black("Notes-cli"),
			utils.Black(constant.Version),
			utils.Yellow(len(items)),
			utils.Yellow(len(items.GetNotes())),
			utils.Yellow(len(items.GetTasks()))),
	)
	itemTable.SetRowPainter(rowPainter)
	return itemTable
}

func rowPainter(row table.Row) text.Colors {
	switch {
	case row[1] == model.TASK:
		return TaskColor
	default:
		return NoteColor
	}
}

func status(item model.Item) string {
	switch {
	case item.Type == model.TASK && item.Completed:
		return utils.Green(constant.COMPLETED)
	case item.Type == model.TASK && !item.Completed:
		return utils.Red(constant.INCOMPLETED)
	default:
		return constant.NONE
	}
}
