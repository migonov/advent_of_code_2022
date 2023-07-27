package main

import "fmt"

func main() {
	var opponent, output string

	var result map[string]int
	result = make(map[string]int)
	result["X"] = 0
	result["Y"] = 3
	result["Z"] = 6

	var shape map[string]string
	shape = make(map[string]string)
	shape["AX"] = "C"
	shape["AY"] = "A"
	shape["AZ"] = "B"
	shape["BX"] = "A"
	shape["BY"] = "B"
	shape["BZ"] = "C"
	shape["CX"] = "B"
	shape["CY"] = "C"
	shape["CZ"] = "A"

	var value map[string]int
	value = make(map[string]int)
	value["A"] = 1
	value["B"] = 2
	value["C"] = 3

	points := 0
	for {
		fmt.Scan(&opponent, &output)

		points += value[shape[opponent+output]] + result[output]
		fmt.Println(points)
	}
}
