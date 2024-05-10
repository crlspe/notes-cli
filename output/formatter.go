package output

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/jedib0t/go-pretty/v6/text"
)

type FormatFn func(string) string

type Formatter struct{
	Expression	string
	Format		FormatFn
}

type FormatterList []Formatter

func (f *FormatterList) Add(expression string, formatFn FormatFn) {
	*f = append(*f, Formatter{Expression: expression, Format: formatFn})
} 

func (f FormatterList) Apply(value *string) {
	for _, formater := range f{ 
		*value = f.colorizeExpression(*value, formater.Expression, formater.Format)
	} 
}

func (f *FormatterList) colorizeExpression(text string, expresion string, colorFn func(string) string) string {
	regex := regexp.MustCompile(expresion)
	modifiedText := regex.ReplaceAllStringFunc(text, colorFn)
	return modifiedText
}


func TrucateContent(length int) FormatFn { 
	return func(value string) string {
		return truncate(value, length)
	}
}

func WrapContent(length int) FormatFn { 
	return func (value string) string {
		return text.WrapSoft(value,length)
	}
}

func truncate(text string, length int) string {
	if len(text) <= length {
		return text
	}
	truncated := text[:length]
	totalChars := utf8.RuneCountInString(text)
	return fmt.Sprintf("%s...(%d)", truncated, totalChars)
}

func Get(text, expression string) []string {
	regex := regexp.MustCompile(expression)
	scopes := regex.FindAllString(text, -1)
	return scopes
}
