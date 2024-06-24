package game

import (
	"fmt"
	"math/rand"
	"os"
)

func RandomGrid(matrix, used *[][]bool) {
	h, w := readRandomHW()

	if h < 3 || w < 3 {
		fmt.Println("invalid size")
		os.Exit(0)
	}

	for i := 0; i < h; i++ {
		*matrix = append(*matrix, make([]bool, 0))
		*used = append(*used, make([]bool, 0))
		for j := 0; j < w; j++ {
			(*matrix)[i] = append((*matrix)[i], rand.Intn(2) == 0)
			(*used)[i] = append((*used)[i], rand.Intn(2) == 0)
		}
	}
}

func readRandomHW() (int, int) {
	for i, v := range *FlagRandom {
		if v == 'x' && i != len(*FlagRandom)-1 {
			return stringToInt((*FlagRandom)[:i]), stringToInt((*FlagRandom)[i+1:])
		}
	}
	return 0, 0
}
