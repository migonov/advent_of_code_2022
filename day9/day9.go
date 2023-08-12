package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	headX := 0
	headY := 0
	tailX := 0
	tailY := 0
	visited := make(map[int]map[int]struct{})
	visited[0] = make(map[int]struct{})
	visited[0][0] = struct{}{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		inputSlice := strings.Split(input, " ")
		direction := inputSlice[0]
		steps, _ := strconv.Atoi(inputSlice[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				headX++
				if !isTailClose(headX, headY, tailX, tailY) {
					tailX++
					tailY = headY
				}
			case "L":
				headX--
				if !isTailClose(headX, headY, tailX, tailY) {
					tailX--
					tailY = headY
				}
			case "U":
				headY++
				if !isTailClose(headX, headY, tailX, tailY) {
					tailY++
					tailX = headX
				}
			case "D":
				headY--
				if !isTailClose(headX, headY, tailX, tailY) {
					tailY--
					tailX = headX
				}
			}
			if visited[tailX] == nil {
				visited[tailX] = make(map[int]struct{})
			}
			visited[tailX][tailY] = struct{}{}
		}
	}
	counter := 0
	for _, tmp := range visited {
		counter += len(tmp)
	}
	fmt.Println(counter)
}

func isTailClose(headX int, headY int, tailX int, tailY int) bool {
	return (headX == tailX && isTailYClose(headY, tailY)) || (headX-1 == tailX && isTailYClose(headY, tailY)) || (headX+1 == tailX && isTailYClose(headY, tailY))
}

func isTailYClose(headY int, tailY int) bool {
	return headY == tailY || headY-1 == tailY || headY+1 == tailY
}
