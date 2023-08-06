package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	grid := parseInput()

	boolGrid := make([][]bool, 0)
	for i := 0; i < len(grid); i++ {
		tmp := make([]bool, len(grid[0]))
		boolGrid = append(boolGrid, tmp)
	}

	//bad
	for i := 0; i < len(grid); i++ {
		max := grid[i][0]
		boolGrid[i][0] = true
		for j := 1; j < len(grid[i])-1; j++ {
			if max < grid[i][j] {
				boolGrid[i][j] = true
				max = grid[i][j]
			}
		}
		boolGrid[i][len(grid[i])-1] = true
		max = grid[i][len(grid[i])-1]
		for j := len(grid[i]) - 2; j > 0; j-- {
			if max < grid[i][j] {
				boolGrid[i][j] = true
				max = grid[i][j]
			}
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		max := grid[0][i]
		boolGrid[0][i] = true
		for j := 1; j < len(grid)-1; j++ {
			if max < grid[j][i] {
				boolGrid[j][i] = true
				max = grid[j][i]
			}
		}
		boolGrid[len(grid)-1][i] = true
		max = grid[len(grid)-1][i]
		for j := len(grid) - 2; j > 0; j-- {
			if max < grid[j][i] {
				boolGrid[j][i] = true
				max = grid[j][i]
			}
		}
	}

	sum := 0
	for _, row := range boolGrid {
		for _, i := range row {
			if i {
				sum++
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(maxScenic(grid))
}

func maxScenic(grid [][]int) int {
	max := 0

	for i, row := range grid {
		for j, tree := range row {
			maxCandidate := 1
			localMax := 0
			for x := j - 1; x >= 0; x-- {
				if grid[i][x] >= tree {
					localMax++
					break
				}
				localMax++
			}
			maxCandidate *= localMax
			localMax = 0
			for x := j + 1; x < len(row); x++ {
				if grid[i][x] >= tree {
					localMax++
					break
				}
				localMax++
			}
			maxCandidate *= localMax
			localMax = 0
			for x := i - 1; x >= 0; x-- {
				if grid[x][j] >= tree {
					localMax++
					break
				}
				localMax++
			}
			maxCandidate *= localMax
			localMax = 0
			for x := i + 1; x < len(grid); x++ {
				if grid[x][j] >= tree {
					localMax++
					break
				}
				localMax++
			}
			maxCandidate *= localMax
			if maxCandidate > max {
				max = maxCandidate
			}
		}
	}

	return max
}

func parseInput() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 0)

	for scanner.Scan() {
		inputLine := scanner.Text()
		if inputLine == "" {
			break
		}

		gridLevel := make([]int, 0)
		for _, num := range inputLine {
			n, _ := strconv.Atoi(string(num))
			gridLevel = append(gridLevel, n)
		}
		grid = append(grid, gridLevel)
	}

	return grid
}
