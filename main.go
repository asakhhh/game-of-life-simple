package main

import (
	"fmt"
	"flag"
)

/*
About command line args:
--help conflicts with all other args
--random= and --file= conflict with each other
all other args are ok
*/

func main() {
	height, width := readHeightWidth()

	matrix := make([][]bool, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}

	// correct := InputMatrix(&matrix)
	// if !correct {
	// 	//
	// }

	game(&matrix)
}

func game(matrix *[][]bool) {
	printMatrixSimple(matrix)
	
	height := len(*matrix)
	width := len((*matrix)[0])

	neighbourCount := make([][]int, height)
	for i := 0; i < height; i++ {
		neighbourCount[i] = make([]int, width)
	}
	
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (*matrix)[i][j] {
				for y := i - 1; y <= i + 1; y++ {
					for x := j - 1; x <= j + 1; x++ {
						if y == i && x == j {
							continue
						}
						if y < 0 || x < 0 || y >= height || x >= width {
							continue
						}
						neighbourCount[y][x]++
					}
				}
			}
		}
	}
	
	gameContinues := false
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (*matrix)[i][j] {
				if neighbourCount[i][j] == 2 || neighbourCount[i][j] == 3 {
					gameContinues = true
				} else {
					(*matrix)[i][j] = false
				}
			} else {
				if neighbourCount[i][j] == 3 {
					(*matrix)[i][j] = true
					gameContinues = true
				}
			}
		}
	}
	
	if !gameContinues {
		printMatrixSimple(matrix)
	} else {
		game(matrix)
	}
}

func printMatrixSimple(matrix *[][]bool) { //
	height := len(*matrix)
	width := len((*matrix)[0])
	
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (*matrix)[i][j] {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("=======================")
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

func printUsage() {
	fmt.Printf(`Usage: go run main.go [options]

Options:
  --help        : Show the help message and exit
  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name
  --delay-ms=X  : Set the animation speed in milliseconds. Default is 2500 milliseconds
  --file=X      : Load the initial grid from a specified file
  --edges-portal: Enable portal edges where cells that exit the grid appear on the opposite side
  --random=HxW  : Generate a random grid of the specified width (W) and height (H)
  --fullscreen  : Adjust the grid to fit the terminal size with empty cells
  --footprints  : Add traces of visited cells, displayed as '∘'
  --colored     : Add color to live cells and traces if footprints are enabled
`)
}
