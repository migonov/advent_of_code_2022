package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stacks [][]string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var inputLines []string
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		inputLines = append(inputLines, input)
	}
	stacks = parseStacks(inputLines)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}

		splittedInput := strings.Split(input, " ")
		n, _ := strconv.Atoi(splittedInput[1])
		fromStack, _ := strconv.Atoi(splittedInput[3])
		toStack, _ := strconv.Atoi(splittedInput[5])
		moveCrates(n, fromStack-1, toStack-1)
	}

	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
}

func moveCrates(n int, from int, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-n:]...)
	stacks[from] = stacks[from][:len(stacks[from])-n]
}

func parseStacks(inputLines []string) [][]string {
	parsedInput := make([][]string, (len(inputLines[0])+1)/4)
	for _, line := range inputLines[:len(inputLines)-1] {
		stackNumber := 0
		for index, character := range line {
			if index%4 == 1 {
				if character != ' ' {
					parsedInput[stackNumber] = append([]string{string(character)}, parsedInput[stackNumber]...)
				}
				stackNumber++
			}
		}
	}
	return parsedInput
}
