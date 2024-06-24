package game

import (
	"fmt"
	"os"

	tsize "github.com/kopoli/go-terminal-size"
)

func getTerminalSize() (int, int) {
	size, err := tsize.GetSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		os.Exit(1)
	}
	return size.Height, size.Width
}

func stringToInt(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(rune(c)-'0')
	}
	return n
}
