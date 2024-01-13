package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show() {
	for _, row := range u {
		for _, cell := range row {
			if cell {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	for i := 0; i < (width*height)/4; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		u[y][x] = true
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(j == 0 && i == 0) && u.Alive(x+i, y+j) {
				count++
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	alive := u.Alive(x, y)
	neighbors := u.Neighbors(x, y)

	return neighbors == 3 || alive && neighbors == 2
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b[y][x] = a.Next(x, y)
		}
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 10; i++ {
		Step(a, b)
		fmt.Print("\x0c") // Clear screen
		b.Show()
		time.Sleep(time.Second / 2)
		a, b = b, a // Swap universes
	}
}
