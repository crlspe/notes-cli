package utils

import (
	"regexp"

	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/jedib0t/go-pretty/v6/text"
)

func color(value interface{}, color text.Color) string {
	return text.Color.Sprint(color, value)
}

func Yellow(value interface{}) string {
	return color(value, text.FgYellow)
}

func yellow(value string) string {
	return Yellow(value)
}

func Green(value interface{}) string {
	return color(value, text.FgGreen)
}

func Red(value interface{}) string {
	return color(value, text.FgRed)
}

func White(value interface{}) string {
	return color(value, text.FgWhite)
}

func HiWhite(value interface{}) string {
	return color(value, text.FgHiWhite)
}

func Cyan(value interface{}) string {
	return color(value, text.FgCyan)
}

func Blue(value interface{}) string {
	return color(value, text.FgBlue)
}

func Magenta(value interface{}) string {
	return color(value, text.FgMagenta)
}

func Black(value interface{}) string {
	return color(value, text.FgBlack)
}

func BgGreen(value interface{}) string {
	return color(value, text.BgGreen)
}

func BgBlue(value interface{}) string {
	return color(value, text.BgBlue)
}

func BgYellow(value interface{}) string {
	return color(value, text.BgYellow)
}

func Bold(value interface{}) string {
	return color(value, text.Bold)
}

func ColorScopes(text string) string {
	regex := regexp.MustCompile(constant.ScopeRegEx)
	modifiedText := regex.ReplaceAllStringFunc(text, yellow)
	return modifiedText
}
