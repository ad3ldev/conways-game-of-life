package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var directions [][]int

func checkCell(x int, y int, array [][]bool) bool {
	count := 0
	for i := 0; i < len(directions); i++ {
		nextX := x + directions[i][0]
		nextY := y + directions[i][1]
		if nextX < 0 || nextX >= len(array) ||
			nextY < 0 || nextY >= len(array[x]) {
			continue
		}
		if !array[nextX][nextY] {
			continue
		}
		if array[nextX][nextY] {
			count++
		}
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
func game() {
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
func main() {
	directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
}
