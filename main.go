package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		width  = 50
		height = 10
	)

	board := MakeBoard(height, width)

	buf := make([]rune, 0, height*width)
	var (
		px, py int
		vx, vy = 1, 1
		ms     = 100
	)

	for {

		draw(board, buf)
		board[px][py] = false

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}
		board[px][py] = true

		time.Sleep(time.Millisecond * time.Duration(ms))

	}

}

func MakeBoard(h int, w int) [][]bool {
	board := make([][]bool, w)

	for row := range board {
		board[row] = make([]bool, h)
	}

	return board
}

func draw(board [][]bool, b []rune) {
	filledCell := '⚽'
	emptyCell := ' '
	b = b[:0]
	var cell rune

	for y := range board[0] {
		for x := range board {
			cell = emptyCell
			if board[x][y] {
				cell = filledCell
			}

			b = append(b, cell, ' ')

		}
		b = append(b, '\n')
	}
	fmt.Print(string(b))
}
