package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// instructions, g := parse("example_part2.txt")
	instructions, g := parse("input.txt")
	fmt.Println(instructions)

	// fmt.Println(g)
	// for k, v := range g {
	// fmt.Println(k, v)
	// }

	fmt.Println("only A's")
	a := ending_a(g)
	starts := []*node{}
	z_record := []int{}
	for k, v := range a {
		fmt.Println(k, v)
		starts = append(starts, v)
	}
	fmt.Println(starts)

	for i, _ := range starts {
		// if move == 'L' {
		// 	starts[i] = starts[i].left
		// }
		// if move == 'R' {
		// 	starts[i] = starts[i].right
		// }

		// starts[i] = traverse(starts[i], instructions, g)

		temp := starts[i]
		z_count := 0
		for temp.name[2] != 'Z' {
			temp = traverse(temp, instructions, g)
			z_count++
		}
		z_record = append(z_record, z_count)
	}
	fmt.Println(z_record)

	ans := len(instructions)
	for _, z := range z_record {
		ans *= z
	}

	fmt.Println(ans)
}

type node struct {
	name       string
	left_name  string
	left       *node
	right_name string
	right      *node
}

func ending_a(graph map[string]*node) map[string]*node {
	a := map[string]*node{}
	for k, v := range graph {
		if k[2] == 'A' {
			a[k] = v
		}
	}
	return a
}

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
	// fmt.Println(instructions)

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
