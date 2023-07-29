package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fmt.Println(findStartOfPacketMarker(input))
	fmt.Println(findStartOfMessageMarker(input))
}

func findStartOfMessageMarker(input string) int {
	list := ""
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(list); j++ {
			if list[j] == input[i] {
				list = list[j+1:]
				break
			}
		}
		list += string(input[i])
		if len(list) == 14 {
			return i + 1
		}
	}
	return -1
}

func findStartOfPacketMarker(input string) int {
	threeAgo := input[0]
	twoAgo := input[1]
	oneAgo := input[2]

	for i := 3; i < len(input); i++ {
		if oneAgo == '-' || oneAgo == input[i] {
			oneAgo = input[i]
			twoAgo = '-'
			threeAgo = '-'
		} else if twoAgo == '-' || twoAgo == input[i] {
			twoAgo = oneAgo
			oneAgo = input[i]
			threeAgo = '-'
		} else if threeAgo == '-' || threeAgo == input[i] {
			threeAgo = twoAgo
			twoAgo = oneAgo
			oneAgo = input[i]
		} else {
			return i + 1
		}
	}
	return -1
}
