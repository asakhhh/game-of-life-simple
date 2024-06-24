package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	tsize "github.com/kopoli/go-terminal-size"
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
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
	BrightRed     = "\033[91m"
)

var (
	flagHelp        *bool
	flagVerbose     *bool
	flagDelayms     *int
	flagFile        *string
	flagEdgesPortal *bool
	flagRandom      *string
	flagFullscreen  *bool
	flagFootprints  *bool
	flagColored     *bool
	tickNumber      int
)

/*
About command line args:
--help conflicts with all other args
--random= and --file= conflict with each other
all other args are ok
*/

func isValidArg(s string) bool {
	if s == "--help" || s == "--verbose" || s == "--edges-portal" || s == "--fullscreen" || s == "--footprints" || s == "--colored" {
		return true
	}
	if len(s) > 10 && s[:11] == "--delay-ms=" {
		return true
	}
	if len(s) > 6 && s[:7] == "--file=" {
		return true
	}
	if len(s) > 8 && s[:9] == "--random=" {
		return true
	}
	return false
}

func parseArgs() {
	var t, errArgs []string
	t = append(t, os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		if isValidArg(os.Args[i]) {
			t = append(t, os.Args[i])
		} else {
			errArgs = append(errArgs, os.Args[i])
		}
	}
	fmt.Printf("The following arguments were incorrect: " + Red)
	for i, v := range errArgs {
		fmt.Printf("%s", v)
		if i != len(errArgs)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Println("\n" + Reset)

	os.Args = t

	flagHelp = flag.Bool("help", false, "")
	flagVerbose = flag.Bool("verbose", false, "")
	flagDelayms = flag.Int("delay-ms", -1, "")
	flagFile = flag.String("file", "", "")
	flagEdgesPortal = flag.Bool("edges-portal", false, "")
	flagRandom = flag.String("random", "", "")
	flagFullscreen = flag.Bool("fullscreen", false, "")
	flagFootprints = flag.Bool("footprints", false, "")
	flagColored = flag.Bool("colored", false, "")

	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) > 10 && os.Args[i][:11] == "--delay-ms=" {
			numstring := os.Args[i][11:]
			if len(numstring) == 0 || stringToInt(numstring) == -1 {
				os.Args[i] = os.Args[i][:11] + "-2"
			}
		}
	}

	flag.Parse()

	if *flagHelp {
		if os.Args[1] == "--help" {
			return
		}
		*flagHelp = false
	}

	if *flagRandom != "" && *flagFile != "" {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i][:7] == "--file=" {
				*flagRandom = ""
				return
			} else if os.Args[i][:9] == "--random=" {
				*flagFile = ""
				return
			}
		}
	}
}

func main() {
	parseArgs()
	if !*flagHelp && *flagDelayms < 0 {
		fmt.Println("Delay in ms was either not set or inputted incorrectly. Default value of 2500 ms will be used.")
		*flagDelayms = 2500
	}

	if *flagHelp {
		printHelp()
		os.Exit(0)
	}

	var matrix [][]bool
	var used [][]bool
	if len(*flagFile) != 0 {
		FileGrid(&matrix)
	} else if len(*flagRandom) != 0 {
		RandomGrid(&matrix)
	} else {
		height, width := readHeightWidth()
		SetSizeToMatrix(&matrix, height, width)
		SetSizeToMatrix(&used, height, width)
		correct := inputMatrix(&matrix)
		if !correct {
			fmt.Printf("Your map input was invalid. Only . and # characters can be entered.\n")
			return
		}
	}
	for i := range matrix {
		used = append(used, make([]bool, len(matrix[0])))
		for j := range matrix[0] {
			used[i][j] = matrix[i][j]
		}
	}
	// if *flagFullscreen {
	// 	termHeight, termWidth := getTerminalSize()
	// 	if len(matrix) > termHeight {
	// 		matrix = matrix[:termHeight]
	// 	}
	// 	for i := range matrix {
	// 		if len(matrix[i]) > termWidth/2 {
	// 			matrix[i] = matrix[i][:termWidth/2]
	// 		}
	// 	}
	// 	SetSizeToMatrix(&used, len(matrix), len(matrix[0]))
	// 	for i := range matrix {
	// 		copy(used[i], matrix[i])
	// 	}
	// }

	game(&matrix, &used)
}

func game(matrix, used *[][]bool) {
	printMatrix(matrix, used)
	time.Sleep(time.Duration(*flagDelayms) * time.Millisecond)

	height := len(*matrix)
	width := len((*matrix)[0])

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
						if *flagEdgesPortal {
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
	} else if !gameContinues {
		printMatrix(matrix, used)
		fmt.Println("\nNo live cells left.")
	} else {
		game(matrix, used)
	}
}

func readHeightWidth() (int, int) {
	fmt.Printf("Enter height and width: ")

	inp := readLine()
	hasSpace := false
	var num1, num2 string

	if len(inp) == 0 {
		fmt.Printf("Please enter two numbers in format: <h w>\n\n")
		return readHeightWidth()
	}

	i := 0
	for i < len(inp) {
		if inp[i] == ' ' {
			if !hasSpace {
				for inp[i] == ' ' {
					i++
				}
				hasSpace = true
			} else {
				for i < len(inp) {
					if inp[i] != ' ' {
						fmt.Printf("Please enter two numbers in format: <h w>\n\n")
						return readHeightWidth()
					}
					i++
				}
			}
		} else if inp[i] < '0' || inp[i] > '9' {
			fmt.Printf("Please enter two numbers in format: <h w>\n\n")
			return readHeightWidth()
		} else {
			if hasSpace {
				num2 += string(inp[i])
			} else {
				num1 += string(inp[i])
			}
			i++
		}
	}

	if len(num1) == 0 || len(num2) == 0 {
		fmt.Printf("Please enter two numbers in format: <h w>\n\n")
		return readHeightWidth()
	}

	h, w := stringToInt(num1), stringToInt(num2)
	if h < 3 || w < 3 {
		fmt.Printf("Height and width should be at least 3.\n\n")
		return readHeightWidth()
	}

	return h, w
}

func readLine() string {
	var inp string
	var r rune
	fmt.Scanf("%c", &r)
	for r != '\n' {
		inp += string(r)
		fmt.Scanf("%c", &r)
	}
	return inp
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

func printHelp() {
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
	fmt.Println(Blue + "  --random=WxH" + Reset + "   : Generate a random grid of the specified width (W) and height (H)")
	fmt.Println(Blue + "  --fullscreen" + Reset + "   : Adjust the grid to fit the terminal size with empty cells")
	fmt.Println(Blue + "  --footprints" + Reset + "   : Add traces of visited cells, displayed as '∘'")
	fmt.Println(Blue + "  --colored" + Reset + "      : Add color to live cells and traces if footprints are enabled")
	fmt.Println()
}

func inputMatrix(matrix *[][]bool) bool {
	for i := range *matrix {
		if line := readLine(); len(line) == len((*matrix)[i]) && checkLineOfMatrix(&line) {
			for j, v := range line {
				(*matrix)[i][j] = v == '#'
			}
		} else {
			return false
		}
	}
	return true
}

func checkLineOfMatrix(line *string) bool {
	for _, v := range *line {
		if v != '.' && v != '#' {
			return false
		}
	}
	return true
}

func printVerbose(height, width, aliveCells int) {
	fmt.Printf("Tick: %d\n", tickNumber)
	fmt.Printf("Grid Size: %dx%d\n", height, width)
	fmt.Printf("Live cells: %d\n", aliveCells)
	fmt.Printf("DelayMs: %dms\n\n", *flagDelayms)
}

func FileGrid(matrix *[][]bool) {
	if grid, err := os.ReadFile(*flagFile); err != nil {
		fmt.Println("Game-of-Life:" + *flagFile + ": No such file or directory")
		os.Exit(1)
	} else {
		var line []bool
		for i, v := range string(grid) {
			if v != '.' && v != '#' && v != '\n' {
				fmt.Println("Invalid symbol in file")
				os.Exit(1)
			}
			if v != '\n' {
				line = append(line, v == '#')
				if i != len(string(grid))-1 {
					continue
				}
			}
			if len(*matrix) != 0 && len((*matrix)[0]) != len(line) {
				fmt.Println("Wrong size line of grid in file")
				os.Exit(1)
			}
			*matrix = append(*matrix, line)
			line = []bool{}
		}
	}
}

func RandomGrid(matrix *[][]bool) {
	h, w := readRandomWH()
	if h < 3 || w < 3 {
		fmt.Println("invalid size")
		os.Exit(0)
	}
	SetSizeToMatrix(matrix, h, w)
	for i := range *matrix {
		for j := range (*matrix)[i] {
			(*matrix)[i][j] = rand.Intn(2) == 0
		}
	}
}

func SetSizeToMatrix(matrix *[][]bool, h, w int) {
	*matrix = make([][]bool, h)
	for i := range *matrix {
		(*matrix)[i] = make([]bool, w)
	}
}

func readRandomWH() (int, int) {
	for i, v := range *flagRandom {
		if v == 'x' && i != len(*flagRandom)-1 {
			return stringToInt((*flagRandom)[:i]), stringToInt((*flagRandom)[i+1:])
		}
	}
	return 0, 0
}

func getTerminalSize() (int, int) {
	size, err := tsize.GetSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		os.Exit(1)
	}
	return size.Height, size.Width
}

func printMatrix(matrix, used *[][]bool) {
	height, width := len(*matrix), len((*matrix)[0])
	var termHeight, termWidth int

	termHeight, termWidth = getTerminalSize()

	// if *flagFullscreen {
	// 	termHeight, termWidth = getTerminalSize()
	// } else {
	// 	termHeight = height
	// 	termWidth = width
	// }

	fmt.Println()
	for x := 0; x < termWidth; x++ {
		fmt.Printf("=")
	}
	fmt.Println()

	tickNumber++

	if *flagVerbose {
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

	verboseTrim := 0
	if *flagVerbose {
		verboseTrim = 5
	}
	for y := 0; y < height && y < termHeight-verboseTrim; y++ {
		for x := 0; x < width && x*2 < termWidth; x++ {
			if (*matrix)[y][x] {
				if *flagColored {
					fmt.Print(Green)
				}
				fmt.Printf("× " + Reset)
			} else if *flagFootprints && (*used)[y][x] {
				if *flagColored {
					fmt.Print(Yellow)
				}
				fmt.Printf("∘ " + Reset)
			} else {
				fmt.Printf("· ")
			}
		}
		if (termHeight > height && y != height-1) || (termHeight <= height && y != termHeight-1-verboseTrim) {
			fmt.Println()
		}
	}
	for y := height; y < termHeight-1-verboseTrim; y++ {
		fmt.Println()
	}
}
