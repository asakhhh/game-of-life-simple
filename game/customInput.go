package game

import (
	"fmt"
	"os"
)

func InputMatrix(matrix, used *[][]bool) {
	height, width := readHeightWidth()
	for i := 0; i < height; i++ {
		*matrix = append(*matrix, make([]bool, width))
		*used = append(*used, make([]bool, width))
		if line := readLine(); len(line) == width && checkLineOfMatrix(&line) {
			for j, v := range line {
				(*matrix)[i][j] = v == '#'
				(*used)[i][j] = v == '#'
			}
		} else {
			fmt.Printf("Your map input was invalid. Only . and # characters can be entered.\n")
			os.Exit(1)
		}
	}
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

func checkLineOfMatrix(line *string) bool {
	for _, v := range *line {
		if v != '.' && v != '#' {
			return false
		}
	}
	return true
}
