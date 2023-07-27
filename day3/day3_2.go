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
	//elf counter
	counter := 0

	//set of items of first elf
	var set map[rune]struct{}
	set = make(map[rune]struct{})

	//set of common items of first 2 elfs
	var commonSet map[rune]struct{}
	commonSet = make(map[rune]struct{})

	var member struct{}

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}

		if counter%3 == 0 {
			//first elf in a group
			set = make(map[rune]struct{})

			for index, character := range input {
				set[character] = member
				fmt.Println(index)
			}
		} else if counter%3 == 1 {
			//second elf in a group
			commonSet = make(map[rune]struct{})
			for index, character := range input {
				if _, ok := set[character]; ok {
					commonSet[character] = member
				}
				fmt.Println(index)
			}
		} else {
			//third elf in a group
			for index, character := range input {
				if _, ok := commonSet[character]; ok {
					ascii := int(character)
					if ascii < lowercaseAValue {
						sum += ascii - uppercaseAValue + 27
					} else {
						sum += ascii - lowercaseAValue + 1
					}
					break
				}
				fmt.Println(index)
			}
		}
		counter++
	}

	fmt.Println(sum)
}
