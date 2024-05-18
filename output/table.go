package output

import (
	"fmt"
	"strings"

	c "github.com/crlspe/notes-cli/constant"
	"github.com/crlspe/notes-cli/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var CompletedColor = text.Colors{text.FgGreen}
var IncompletedColor = text.Colors{text.FgRed}
var NoteColor = text.Colors{text.FgWhite}
var TaskColor = text.Colors{text.FgCyan}

func PrintShortTable(items model.ItemList) {
	itemTable := initializeTable(items)

	itemTable.AppendHeader(table.Row{
		c.StrSpace,
		formatHeader(c.HeaderType),
		formatHeader(c.HeaderContent),
		formatHeader(c.HeaderScopes),
	})

	for _, item := range items {
		itemTable.AppendRow(table.Row{
			formatStatus(item),
			item.Type,
			formatContent(item.Content),
			formatScopes(item.Content),
		})
	}
	fmt.Println(itemTable.Render())
	fmt.Println()
}

func rowPainter(row table.Row) text.Colors {
	if row[1] == model.TASK {
		return TaskColor
	}
	return NoteColor
}

func initializeTable(items model.ItemList) table.Writer {
	var itemTable = table.NewWriter()
	itemTable.SetStyle(table.StyleRounded)
	itemTable.SetTitle(title(items))
	itemTable.SetRowPainter(rowPainter)
	return itemTable
}

func formatScopes(content string) string {
	var scopes = strings.Join(Get(content, c.ReScope), c.StrSpace) +
		strings.Join(Get(content, c.ReTag), c.StrSpace)
	var scopeFormatter = FormatterList{
		{Expression: c.ReScope, Format: Yellow},
		{Expression: c.ReTag, Format: Magenta},
		{Expression: c.ReAnyString, Format: WrapContent(20)},
	}
	scopeFormatter.Apply(&scopes)
	return scopes
}

func formatContent(content string) string {
	var contentFormatter = FormatterList{
		{Expression: c.ReScope, Format: Yellow},
		{Expression: c.ReTag, Format: Magenta},
		{Expression: c.ReAnyString, Format: TrucateContent(120)},
		{Expression: c.ReAnyString, Format: WrapContent(60)},
	}
	contentFormatter.Apply(&content)
	return content
}

func formatHeader(header string) string {
	var headerFormater = FormatterList{
		{Expression: c.ReAnyString, Format: Cyan},
	}
	headerFormater.Apply(&header)
	return header
}

func formatStatus(item model.Item) string {
	switch {
	case item.Type == model.TASK && item.Completed:
		return Green(c.TaskCompleted)
	case item.Type == model.TASK && !item.Completed:
		return Red(c.TaskIncompleted)
	default:
		return c.ItemNone
	}
}

func title(items model.ItemList) string {
	return fmt.Sprint(
		c.HTitleTotalCount, YellowI(len(items)), c.HTitleSeparator,
		c.HTitleNoteCount, YellowI(len(items.GetNotes())), c.HTitleSeparator,
		c.HTitleTaskCount, YellowI(len(items.GetTasks())),
	)
}
