package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func north(x int, y int, array [][]bool) bool {
	x -= 1
	if x < 0 {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func east(x int, y int, array [][]bool) bool {
	y += 1
	if y >= len(array[x]) {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func west(x int, y int, array [][]bool) bool {
	y -= 1
	if y < 0 {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func south(x int, y int, array [][]bool) bool {
	x += 1
	if x >= len(array) {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func nEast(x int, y int, array [][]bool) bool {
	x -= 1
	y += 1
	if x < 0 || y >= len(array[x]) {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func sEast(x int, y int, array [][]bool) bool {
	x += 1
	y += 1
	if x >= len(array) || y >= len(array[x]) {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func sWest(x int, y int, array [][]bool) bool {
	x += 1
	y -= 1
	if x >= len(array) || y < 0 {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func nWest(x int, y int, array [][]bool) bool {
	x -= 1
	y -= 1
	if x < 0 || y < 0 {
		return false
	}
	if array[x][y] {
		return true
	}
	return false
}
func checkCell(x int, y int, array [][]bool) bool {
	count := 0
	if north(x, y, array) {
		count++
	}
	if east(x, y, array) {
		count++
	}
	if west(x, y, array) {
		count++
	}
	if south(x, y, array) {
		count++
	}
	if nEast(x, y, array) {
		count++
	}
	if sEast(x, y, array) {
		count++
	}
	if nWest(x, y, array) {
		count++
	}
	if sWest(x, y, array) {
		count++
	}

	if array[x][y] {
		if count < 2 {
			return false
		}
		if count == 2 || count == 3 {
			return true
		}
		if count > 3 {
			return false
		}
	} else {
		if count == 3 {
			return true
		}
	}
	return false
}

func newMatrix(rows int, cols int) [][]bool {
	m := make([][]bool, rows)
	for r := range m {
		m[r] = make([]bool, cols)
	}
	return m
}

func nextGeneration(rows int, cols int, array [][]bool) [][]bool {
	nextGen := newMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			nextGen[i][j] = checkCell(i, j, array)
		}
	}
	return nextGen
}
func printGeneration(rows int, cols int, array [][]bool) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if array[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}
func main() {
	rows := 12
	cols := 12
	array := newMatrix(rows, cols)
	array[0][1] = true
	array[1][2] = true
	array[2][0] = true
	array[2][1] = true
	array[2][2] = true

	array[8][2] = true
	array[8][3] = true
	array[8][4] = true
	array[8][7] = true
	array[8][8] = true
	array[8][9] = true

	for {
		printGeneration(rows, cols, array)
		array = nextGeneration(rows, cols, array)
		time.Sleep(500 * time.Millisecond)
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		err := c.Run()
		if err != nil {
			return
		}
	}
}
