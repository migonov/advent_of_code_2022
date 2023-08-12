package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.in")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	globalFirst := 0
	globalSecond := 0
	globalThird := 0
	currSum := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			converted, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			currSum += converted
		} else {
			// risky but worked in this case
			if currSum > globalFirst {
				globalThird = globalSecond
				globalSecond = globalFirst
				globalFirst = currSum
			}
			currSum = 0
		}
	}

	fmt.Print(globalFirst + globalSecond + globalThird)
}
