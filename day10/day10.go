package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cycle int64
	var x int64
	x = 1
	var signal int64
	signal = 0

	crt := make([][]string, 6)
	fillCrt(crt)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}

		if input == "noop" {
			if cycle%40 >= x-1 && cycle%40 <= x+1 {
				crt[cycle/40][cycle%40] = "#"
			}
			cycle++
			if cycle == 20 || (cycle-20)%40 == 0 {
				signal += x * cycle
			}
			continue
		}
		value, err := strconv.ParseInt(strings.Split(input, " ")[1], 10, 32)
		if err != nil {
			panic(err)
		}

		if (cycle < 20 && cycle+2 >= 20) || ((cycle-20)/40 != (cycle-20+2)/40) {
			if cycle+2 == 20 || (cycle-20+2)%40 == 0 {
				signal += x * (cycle + 2)
			} else {
				signal += x * (cycle + 2 - (cycle+2)%10)
			}
		}
		if cycle%40 >= x-1 && cycle%40 <= x+1 {
			crt[cycle/40][cycle%40] = "#"
		}
		cycle++
		if cycle%40 >= x-1 && cycle%40 <= x+1 {
			crt[cycle/40][cycle%40] = "#"
		}
		cycle++
		x += value
	}

	fmt.Printf("signal = %d\n", signal)
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			fmt.Printf("%s", crt[i][j])
		}
		fmt.Printf("\n")
	}
}

func fillCrt(crt [][]string) {
	for i := 0; i < 6; i++ {
		tmp := make([]string, 40)
		for j := 0; j < 40; j++ {
			tmp[j] = "."
		}
		crt[i] = tmp
	}
}
