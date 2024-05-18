package output

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/crlspe/notes-cli/constant"
	"github.com/crlspe/notes-cli/model"
)

func PrintConsole(items model.ItemList) {
	for _, item := range items {
		fmt.Printf("Id: %.8s\t Content: %s\t Type: %s\t Completed: %t\t CreatedAt: %s\t CompletedAt: %s\n",
			item.Id[len(item.Id)-8:], item.Content, item.Type, item.Completed, item.CreatedAt, item.CompletedAt)
	}
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

func PrintApplicationInfo() {
	fmt.Println(Green(constant.ApplicationName), constant.Space, Magenta(constant.Version))
}
