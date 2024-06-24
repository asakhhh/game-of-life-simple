package main

// bonus flag for dynamic printmatrix
import (
	"fmt"
	"os"

	"crunch03/game"
)

func main() {
	game.ParseArgs()

	if *game.FlagHelp { // --help flag
		game.PrintHelp()
		os.Exit(0)
	}

	var matrix [][]bool
	var used [][]bool

	if len(*game.FlagFile) != 0 { // File input
		game.FileGrid(&matrix, &used)
	} else if len(*game.FlagRandom) != 0 { // Random generation
		game.RandomGrid(&matrix, &used)
	} else {
		game.InputMatrix(&matrix, &used) // Custom input
	}

	if *game.FlagDelayms < 0 {
		fmt.Println(game.Blue + "Delay in ms was either not set or inputted incorrectly. Default value of 2500 ms will be used." + game.Reset)
		*game.FlagDelayms = 2500
	}
	if *game.FlagAlternative && !*game.FlagColored {
		fmt.Println(game.Blue + "Alternative visualization will still include colors." + game.Reset)
		*game.FlagColored = true
	}

	game.Game(&matrix, &used)
}
