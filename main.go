package main

import "fmt"

func main() {
	height, width := readTwoNumbers()

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

func readTwoNumbers() (int, int) {
	inp := readLine()
	has_space := false
	var num1, num2 string

	if len(inp) == 0 {
		return -1, -1
	}

	for i, c := range inp {
		if c == ' ' {
			if has_space || i == 0 {
				return -1, -1
			}
			has_space = true
		} else if c < '0' || c > '9' {
			return -1, -1
		} else {
			if has_space {
				num2 += string(c)
			} else {
				num1 += string(c)
			}
		}
	}

	if len(num2) == 0 {
		return -1, -1
	}

	return stringToInt(num1), stringToInt(num2)
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
