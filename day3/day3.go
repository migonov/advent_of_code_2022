package main

import (
	"bufio"
	"fmt"
	"os"
)

const lowercaseAValue = int('a')
const uppercaseAValue = int('A')

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		half := len(input) / 2
		var m map[rune]int
		m = make(map[rune]int)

		for index, character := range input {
			val := m[character]
			if val == 0 {
				m[character] = index + 1
			} else if val <= half && index >= half {
				ascii := int(character)
				if ascii < lowercaseAValue {
					sum += ascii - uppercaseAValue + 27
				} else {
					sum += ascii - lowercaseAValue + 1
				}
				break
			}
		}
	}

	fmt.Println(sum)
}
