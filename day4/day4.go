package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var firstElf []string
	var secondElf []string
	sum := 0

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}

		splitted := strings.Split(input, ",")
		firstElf = strings.Split(splitted[0], "-")
		secondElf = strings.Split(splitted[1], "-")

		firstElfA, err := strconv.Atoi(firstElf[0])
		if err != nil {
			log.Fatal(err)
		}
		firstElfB, err := strconv.Atoi(firstElf[1])
		if err != nil {
			log.Fatal(err)
		}
		secondElfA, err := strconv.Atoi(secondElf[0])
		if err != nil {
			log.Fatal(err)
		}
		secondElfB, err := strconv.Atoi(secondElf[1])
		if err != nil {
			log.Fatal(err)
		}

		if (firstElfA >= secondElfA && firstElfB <= secondElfB) || (firstElfA <= secondElfA && firstElfB >= secondElfB) {
			sum++
		}
	}

	fmt.Println(sum)
}
