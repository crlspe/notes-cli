package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/crlspe/notes-cli-v4/constant"
	"github.com/jedib0t/go-pretty/v6/text"
)

func GenerateId() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 8)

	if _, err := rand.Read(randomBytes); err != nil {
		return fmt.Sprintf("%d", timestamp)
	}

	randomHex := hex.EncodeToString(randomBytes)
	id := fmt.Sprintf("%d%s", timestamp, randomHex)
	return id
}

func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func GetScopes(text string) []string {
	regex := regexp.MustCompile(constant.ScopeRegEx)
	scopes := regex.FindAllString(text, -1)
	return scopes
}

func Truncate(text string, length int) string {
	if len(text) <= length {
		return text
	}
	truncated := text[:length]
	totalChars := utf8.RuneCountInString(text)
	return fmt.Sprintf("%s...(%d)", truncated, totalChars)
}

func FormatContent(value string, length int) string {
	return text.WrapSoft(ColorScopes(Truncate(value, length*2)), length)
}

func FormatScopes(value string, length int) string {
	return text.WrapSoft(strings.Join(GetScopes(value), ", "), length)
}
