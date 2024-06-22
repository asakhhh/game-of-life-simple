package main

import "fmt"

func main() {
	height, width := ReadTwoNumbers()

	matrix := make([][]bool, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}

	correct := InputMatrix(&matrix)
	if !correct {
		//
	}

	Game(&matrix)
}

func Game(matrix *[][]bool) {
	PrintMatrix(matrix)

	if alive {
		Game()
	}
}

func ReadTwoNumbers() (int, int) {
	inp := ReadLine()
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

	return StringToInt(num1), StringToInt(num2)
}

func ReadLine() string {
	var inp string
	var r rune
	fmt.Scanf("%c", &r)
	for r != '\n' {
		inp += string(r)
		fmt.Scanf("%c", &r)
	}
	return inp
}

func StringToInt(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(rune(c)-'0')
	}
	return n
}
