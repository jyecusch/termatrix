package matrix

import "math/rand"

var possibleChars = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
	'-', '_', '+', '=', '|', '\\', '/', '?', '.', ',',
	'[', ']', '{', '}', '(', ')', '<', '>',
	'`', '~', '\'', '"', ':', ';',
}

func RandomRune() rune {
	return possibleChars[rand.Intn(len(possibleChars))]
}

func newRuneGrid(width, height int) [][]rune {
	grid := make([][]rune, width)
	for i := range grid {
		grid[i] = make([]rune, height)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	return grid
}
