package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []*int64
	operator  string
	val       string
	divisible int64
	t         int64
	f         int64
	inspects  int64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var monkeys []*monkey

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			scanner.Scan()
			input = scanner.Text()
		}
		newMonkey := monkey{}

		scanner.Scan()
		input = scanner.Text()
		if input == "exit" {
			break
		}
		input = strings.TrimSpace(input)
		splittedLine := strings.Split(input, " ")
		for i := 2; i < len(splittedLine)-1; i++ {
			item, err := strconv.ParseInt(splittedLine[i][:len(splittedLine[i])-1], 10, 32)
			if err != nil {
				panic(err)
			}
			newMonkey.items = append(newMonkey.items, &item)
		}
		item, err := strconv.ParseInt(splittedLine[len(splittedLine)-1], 10, 32)
		if err != nil {
			panic(err)
		}
		newMonkey.items = append(newMonkey.items, &item)

		scanner.Scan()
		input = scanner.Text()
		input = strings.TrimSpace(input)
		splittedLine = strings.Split(input, " ")
		println(splittedLine[4])
		newMonkey.operator = splittedLine[4]
		newMonkey.val = splittedLine[5]

		scanner.Scan()
		input = scanner.Text()
		input = strings.TrimSpace(input)
		splittedLine = strings.Split(input, " ")
		divisor, err := strconv.ParseInt(splittedLine[3], 10, 32)
		if err != nil {
			panic(err)
		}
		newMonkey.divisible = divisor

		scanner.Scan()
		input = scanner.Text()
		input = strings.TrimSpace(input)
		splittedLine = strings.Split(input, " ")
		trueMonkey, err := strconv.ParseInt(splittedLine[5], 10, 32)
		if err != nil {
			panic(err)
		}
		newMonkey.t = trueMonkey

		scanner.Scan()
		input = scanner.Text()
		input = strings.TrimSpace(input)
		splittedLine = strings.Split(input, " ")
		falseMonkey, err := strconv.ParseInt(splittedLine[5], 10, 32)
		if err != nil {
			panic(err)
		}
		newMonkey.f = falseMonkey

		newMonkey.inspects = 0

		monkeys = append(monkeys, &newMonkey)
	}

	var factor int64
	factor = 1
	for _, m := range monkeys {
		factor *= m.divisible
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			val := monkey.val
			for _, item := range monkey.items {
				switch monkey.operator {
				case "*":
					if val == "old" {
						*item *= *item
					} else {
						v, _ := strconv.ParseInt(val, 10, 32)
						*item *= v
					}
				case "+":
					if val == "old" {
						*item += *item
					} else {
						v, _ := strconv.ParseInt(val, 10, 32)
						*item += v
					}
				case "-":
					if val == "old" {
						*item -= *item
					} else {
						v, _ := strconv.ParseInt(val, 10, 32)
						*item -= v
					}
				case "/":
					if val == "old" {
						*item /= *item
					} else {
						v, _ := strconv.ParseInt(val, 10, 32)
						*item /= v
					}
				}
				//*item /= 3
				*item %= factor

				if *item%monkey.divisible == 0 {
					monkeys[monkey.t].items = append(monkeys[monkey.t].items, item)
				} else {
					monkeys[monkey.f].items = append(monkeys[monkey.f].items, item)
				}

				monkey.items = monkey.items[1:]
				monkey.inspects++
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspects > monkeys[j].inspects
	})

	fmt.Println(monkeys[0].inspects * monkeys[1].inspects)
}
