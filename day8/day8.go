package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// instructions, g := parse("example.txt")
	// instructions, g := parse("example2.txt")
	instructions, g := parse("input.txt")

	// fmt.Println(g)
	for k, v := range g {
		fmt.Println(k, v)
	}

	cur := g["AAA"]
	count := 0
	for cur.name != "ZZZ" {
		cur = traverse(cur, instructions, g)
		fmt.Println(cur)
		count++
	}
	fmt.Println(count * len(instructions))
}

type node struct {
	name       string
	left_name  string
	left       *node
	right_name string
	right      *node
}

// type graph struct{

// }

func parse(filename string) (string, map[string]*node) {
	graph := map[string]*node{}

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	instructions := fileScanner.Text()
	fmt.Println(instructions)

	fileScanner.Scan() //skip 1 line

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line_1 := strings.Split(line, "=")
		name := strings.Trim(line_1[0], " ")
		line_2 := strings.Split(line_1[1], ",")
		left := strings.Trim(line_2[0], " ")
		left = strings.Trim(left, "(")
		right := strings.Trim(line_2[1], " ")
		right = strings.Trim(right, ")")

		graph[name] = &node{name: name, left_name: left, right_name: right}
	}

	for _, n := range graph {
		ref, found := graph[n.left_name]
		if !found {
			fmt.Println("error bad logic")
			os.Exit(1)
		}
		n.left = ref

		ref, found = graph[n.right_name]
		if !found {
			fmt.Println("error bad logic")
			os.Exit(1)
		}
		n.right = ref
	}
	return instructions, graph
}

func traverse(cur *node, instructions string, graph map[string]*node) *node {
	for _, c := range instructions {
		if c == 'L' {
			cur = cur.left
		}
		if c == 'R' {
			cur = cur.right
		}
	}
	return cur
}
