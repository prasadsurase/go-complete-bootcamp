package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		width  = 50
		height = 10
		blank  = ' '
		ball   = '⚾'
	)

	var (
		cell rune
		px   int
		py   int
	)
	vx, vy := 1, 1
	buff := make([]rune, 0, width*height)

	board := make([][]bool, width)
	for col := range board {
		board[col] = make([]bool, height)
	}

	board[px][py] = true

	for i := 0; i <= 1200; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		buff = buff[:0]
		time.Sleep(100 * time.Millisecond)

		for y := range board[0] {
			for x := range board {
				cell = blank
				if board[x][y] {
					cell = ball
				}
				// fmt.Print(string(cell))
				// fmt.Print(" ")
				buff = append(buff, cell, ' ')
			}
			// fmt.Println()
			buff = append(buff, '\n')
		}
		fmt.Print(string(buff))
	}
}
