package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

func SinglePromptC(label string, color text.Color) string {
	fmt.Print(text.Color.Sprint(color, label))
	var inputReader = bufio.NewReader(os.Stdin)
	var input, _ = inputReader.ReadString('\n')
	return strings.TrimSpace(input)
}

func SinglePrompt(label string) string {
	return SinglePromptC(label, text.FgGreen)
}

func MultiplePromptC(label string, color text.Color) []string {
	var inputs = []string{}
	for {
		var input = SinglePromptC(label, color)
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
