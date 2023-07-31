package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const DISK_SPACE = 70000000
const REQUIRED_SPACE = 30000000

type Node struct {
	name     string
	size     uint `default:"0"`
	parent   *Node
	children []*Node
}

type Tree struct {
	root *Node
}

func find(node *Node, dir string) *Node {
	if node == nil {
		return nil
	}
	if node.name == dir {
		return node
	} else {
		for _, n := range node.children {
			find(n, dir)
		}
	}
	return nil
}

func calcSum(node *Node, sum *uint, toDelete uint, delete *uint) {
	if node == nil {
		return
	}
	if node.size <= 100000 {
		*sum += node.size
	}
	if node.size >= toDelete && node.size < *delete {
		*delete = node.size
	}
	for _, n := range node.children {
		calcSum(n, sum, toDelete, delete)
	}
}

func add(node *Node, val uint) {
	if node == nil {
		return
	}
	node.size += val
	add(node.parent, val)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	currNode := &Node{
		name: "/",
	}
	tree := Tree{
		root: currNode,
	}

	scanner.Scan()
	for scanner.Scan() {
		in := scanner.Text()
		if in == "" {
			break
		}

		if strings.HasPrefix(in, "$") {
			userInput := strings.Split(in[2:], " ")
			userCommand := userInput[0]
			switch userCommand {
			case "cd":
				if userInput[1] == ".." {
					currNode = currNode.parent
				} else {
					alreadyExists := find(tree.root, userInput[1])
					if alreadyExists == nil {
						newNode := &Node{
							name:   userInput[1],
							parent: currNode,
						}
						currNode.children = append(currNode.children, newNode)
						currNode = newNode
					} else {
						currNode = alreadyExists
					}
				}
			case "ls":
				continue
			}
		} else {
			userInput := strings.Split(in, " ")
			if userInput[0] != "dir" {
				s, _ := strconv.Atoi(userInput[0])
				add(currNode, uint(s))
			}
		}
	}
	var sum uint
	sum = 0
	var toDelete uint
	toDelete = REQUIRED_SPACE - (DISK_SPACE - tree.root.size)
	var delete uint
	delete = math.MaxUint32
	calcSum(tree.root, &sum, toDelete, &delete)
	fmt.Println(sum, delete)
}
