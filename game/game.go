package game

import (
	"bytes"
	"fmt"
	"time"
)

const (
	Reset         = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	Gray          = "\033[37m"
	White         = "\033[97m"
	WhiteBG       = "\033[48;5;7m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
	BrightRed     = "\033[91m"
)

var (
	FlagHelp        *bool
	FlagVerbose     *bool
	FlagDelayms     *int
	FlagFile        *string
	FlagEdgesPortal *bool
	FlagRandom      *string
	FlagFullscreen  *bool
	FlagFootprints  *bool
	FlagColored     *bool
	FlagAlternative *bool
	TickNumber      int
	buf             *(bytes.Buffer)
)

func Game(matrix, used *[][]bool) {
	buf = new(bytes.Buffer)
	for {
		printMatrix(matrix, used)
		time.Sleep(time.Duration(*FlagDelayms) * time.Millisecond)

		height, width := len(*matrix), len((*matrix)[0])

		neighbourCount := make([][]int, height)
		for i := 0; i < height; i++ {
			neighbourCount[i] = make([]int, width)
		}

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if (*matrix)[i][j] {
					for y := i - 1; y <= i+1; y++ {
						for x := j - 1; x <= j+1; x++ {
							if y == i && x == j {
								continue
							}
							if *FlagEdgesPortal {
								neighbourCount[(y+height)%height][(x+width)%width]++
							} else {
								if y < 0 || x < 0 || y >= height || x >= width {
									continue
								}
								neighbourCount[y][x]++
							}
						}
					}
				}
			}
		}

		gameContinues, gameChanged := false, false
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if (*matrix)[i][j] {
					if neighbourCount[i][j] == 2 || neighbourCount[i][j] == 3 {
						gameContinues = true
					} else {
						(*matrix)[i][j] = false
						gameChanged = true
					}
				} else {
					if neighbourCount[i][j] == 3 {
						(*matrix)[i][j] = true
						(*used)[i][j] = true
						gameContinues = true
						gameChanged = true
					}
				}
			}
		}

		if !gameChanged {
			fmt.Println("\nThe cell evolution has stopped at this state.")
			break
		} else if !gameContinues {
			printMatrix(matrix, used)
			fmt.Println("\nNo live cells left.")
			break
		}
	}
}
