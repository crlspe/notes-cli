package output

import (
	"github.com/jedib0t/go-pretty/v6/text"
)

func Yellow(value string) string {
	return text.Color.Sprint(text.FgYellow, value)
}

func YellowI(value int) string {
	return text.Color.Sprint(text.FgYellow, value)
}

func Green(value string) string {
	return text.Color.Sprint(text.FgGreen, value)
}

func Red(value string) string {
	return text.Color.Sprint(text.FgRed, value)
}

func White(value string) string {
	return text.Color.Sprint(text.FgWhite, value)
}

func HiWhite(value string) string {
	return text.Color.Sprint(text.FgHiWhite, value)
}

func Cyan(value string) string {
	return text.Color.Sprint(text.FgCyan, value)
}

func Blue(value string) string {
	return text.Color.Sprint(text.FgBlue, value)
}

func Magenta(value string) string {
	return text.Color.Sprint(text.FgMagenta, value)
}

func Black(value string) string {
	return text.Color.Sprint(text.FgBlack, value)
}

func BgGreen(value string) string {
	return text.Color.Sprint(text.BgGreen, value)
}

func BgBlue(value string) string {
	return text.Color.Sprint(text.BgBlue, value)
}

func BgYellow(value string) string {
	return text.Color.Sprint(text.BgYellow, value)
}

func Bold(value string) string {
	return text.Color.Sprint(text.Bold, value)
}
