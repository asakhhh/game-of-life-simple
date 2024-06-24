package game

import (
	"fmt"
	"os"
)

func FileGrid(matrix, used *[][]bool) {
	if grid, err := os.ReadFile(*FlagFile); err != nil {
		fmt.Println("Game-of-Life: " + *FlagFile + ": No such file or directory")
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
			*matrix = append(*matrix, make([]bool, len(line)))
			*used = append(*used, make([]bool, len(line)))
			for j := range line {
				(*matrix)[len(*matrix)-1][j] = line[j]
				(*used)[len(*used)-1][j] = line[j]
			}
			line = []bool{}
		}
	}
}
