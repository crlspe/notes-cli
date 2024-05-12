package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

func readString() string {
	var inputReader = bufio.NewReader(os.Stdin)
	var input, _ = inputReader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readRune() rune {
	var inputReader = bufio.NewReader(os.Stdin)
	var input, _, _ = inputReader.ReadRune()
	return input
}


func SinglePrompt(label string) string {
	fmt.Print(text.FgGreen.Sprint(label))
	return readString()
}

func YesOrNoPrompt(label string) bool {
	fmt.Print(text.FgGreen.Sprint(label))
	var input = readRune()
	if input == 'y' || input == 'Y' {
		return true
	}
	return false
}

func MultiplePromptC(label string, color text.Color) []string {
	var inputs = []string{}
	for {
			fmt.Print(text.FgGreen.Sprint(label))
			var input = readString()
			if input == "" {
				break
			}
			inputs = append(inputs, input)
	}
	return inputs
}

func MultiplePrompt(label string) []string {
	return MultiplePromptC(label, text.FgGreen)
}
