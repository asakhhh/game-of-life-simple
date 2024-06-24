package game

import "fmt"

func printMatrix(matrix, used *[][]bool) {
	height, width := len(*matrix), len((*matrix)[0])
	var termHeight, termWidth int

	termHeight, termWidth = getTerminalSize()

	fmt.Println()
	for x := 0; x < termWidth && (!*FlagFullscreen || x < width*2); x++ {
		fmt.Printf("=")
	}
	fmt.Println()

	TickNumber++

	if *FlagVerbose {
		aliveCells := 0
		for _, row := range *matrix {
			for _, cell := range row {
				if cell {
					aliveCells++
				}
			}
		}
		printVerbose(len(*matrix), len((*matrix)[0]), aliveCells)
	}

	if *FlagFullscreen {
		verboseTrim := 0
		if *FlagVerbose {
			verboseTrim = 5
		}
		for y := 0; y < height && y < termHeight-verboseTrim; y++ {
			for x := 0; x < width && x*2 < termWidth; x++ {
				if (*matrix)[y][x] {
					if *FlagColored {
						fmt.Print(Green)
					}
					if !*FlagAlternative {
						fmt.Printf("× " + Reset)
					} else {
						fmt.Printf("██" + Reset)
					}
				} else if *FlagFootprints && (*used)[y][x] {
					if *FlagColored {
						fmt.Print(Yellow)
					}
					if !*FlagAlternative {
						fmt.Printf("∘ " + Reset)
					} else {
						fmt.Printf("██" + Reset)
					}
				} else {
					if !*FlagAlternative {
						fmt.Printf("· ")
					} else {
						fmt.Printf("██")
					}
				}
			}
			if (termHeight-verboseTrim > height && y != height-1) || (termHeight-verboseTrim <= height && y != termHeight-1-verboseTrim) {
				fmt.Println()
			}
		}
		for y := height; y < termHeight-1-verboseTrim; y++ {
			fmt.Println()
		}
	} else {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if (*matrix)[y][x] {
					if *FlagColored {
						fmt.Print(Green)
					}
					if !*FlagAlternative {
						fmt.Printf("× " + Reset)
					} else {
						fmt.Printf("██" + Reset)
					}
				} else if *FlagFootprints && (*used)[y][x] {
					if *FlagColored {
						fmt.Print(Yellow)
					}
					if !*FlagAlternative {
						fmt.Printf("∘ " + Reset)
					} else {
						fmt.Printf("██" + Reset)
					}
				} else {
					if !*FlagAlternative {
						fmt.Printf("· ")
					} else {
						fmt.Printf("██")
					}
				}
			}
			if y != height-1 {
				fmt.Println()
			}
		}
	}
}

func printVerbose(height, width, aliveCells int) {
	fmt.Printf("Tick: %d\n", TickNumber)
	fmt.Printf("Grid Size: %dx%d\n", height, width)
	fmt.Printf("Live cells: %d\n", aliveCells)
	fmt.Printf("DelayMs: %dms\n\n", *FlagDelayms)
}

func PrintHelp() {
	fmt.Println("Welcome to the ", Green+"Game of Life"+Reset+"!")
	fmt.Println("This simulation models the evolution of cells on a grid. Each cell can be alive or dead, and its state changes based on its neighbors.")
	fmt.Println()
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println(Blue + "  --help   " + Reset + "      : Show the help message and exit")
	fmt.Println(Blue + "  --verbose" + Reset + "      : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	fmt.Println(Blue + "  --delay-ms=X" + Reset + "   : Set the animation speed in milliseconds. Default is 2500 milliseconds")
	fmt.Println(Blue + "  --file=X" + Reset + "       : Load the initial grid from a specified file")
	fmt.Println(Blue + "  --edges-portal" + Reset + " : Enable portal edges where cells that exit the grid appear on the opposite side")
	fmt.Println(Blue + "  --random=HxW" + Reset + "   : Generate a random grid of the specified height (H) and width (W)")
	fmt.Println(Blue + "  --fullscreen" + Reset + "   : Adjust the grid to fit the terminal size with empty cells")
	fmt.Println(Blue + "  --footprints" + Reset + "   : Add traces of visited cells, displayed as '∘'")
	fmt.Println(Blue + "  --colored" + Reset + "      : Add color to live cells and traces if footprints are enabled")
	fmt.Println(Blue + "  --alternative" + Reset + "  : Alternative visualization for the game. --colored arg is automatically included for this option.")
	fmt.Println()
}
