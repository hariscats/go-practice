package main

import (
	"fmt"
	"math/rand"
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
