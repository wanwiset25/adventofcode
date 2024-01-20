package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	patterns := parse("input.txt")
	// patterns := parse("example.txt")
	fmt.Println(patterns)
	for _, p := range patterns {
		for _, l := range p {
			fmt.Println(string(l))
		}
		fmt.Println()
	}

	output := 0
	for _, p := range patterns {
		kind, index := v_or_h(p)
		fmt.Println(string(kind), index)
		if kind == 'v' {
			output += index
		}
		if kind == 'h' {
			output += 100 * index
		}
	}
	fmt.Println(output)
}

func v_or_h(pattern [][]rune) (rune, int) {
	kind := 'v'
	output := find_reflection(pattern)
	if output == 0 {
		output = find_reflection(transpose(pattern))
		kind = 'h'
	}
	if output == 0 {
		fmt.Println("impossible case")
		os.Exit(1)
	}
	return kind, output
}
func transpose(pattern [][]rune) [][]rune {
	output := make([][]rune, 0)

	for j := 0; j < len(pattern[0]); j++ {
		temp := []rune{}
		for i := 0; i < len(pattern); i++ {
			temp = append(temp, pattern[i][j])
		}
		output = append(output, temp)
	}
	return output
}

func find_reflection(pattern [][]rune) int {
	for j := 1; j < len(pattern[0]); j++ {
		same := true
		index_r := j
		index_l := j - 1

		for index_r < len(pattern[0]) && index_l >= 0 {
			if !compare_lines(pattern, index_l, index_r) {
				same = false
			}
			index_l--
			index_r++
		}
		if same {
			return j
		}
	}
	return 0
}

func compare_lines(pattern [][]rune, first_line int, second_line int) bool {
	for i := 0; i < len(pattern); i++ {
		if pattern[i][first_line] != pattern[i][second_line] {
			return false
		}

	}
	return true
}

func parse(filename string) [][][]rune {
	output := make([][][]rune, 0)
	pattern := make([][]rune, 0)
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			output = append(output, pattern)
			pattern = make([][]rune, 0)
			continue
		}
		temp := []rune{}
		for _, r := range line {
			temp = append(temp, r)
		}
		pattern = append(pattern, temp)
	}
	output = append(output, pattern)
	return output
}

// func parse(filename string) ([][]rune, [][]rune) {
// 	pattern1 := make([][]rune, 0)
// 	pattern2 := make([][]rune, 0)
// 	readFile, err := os.Open(filename)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)
// 	cur := 1
// 	for fileScanner.Scan() {
// 		line := fileScanner.Text()
// 		if line == "" {
// 			cur = 2
// 			continue
// 		}
// 		switch cur {
// 		case 1:
// 			temp := []rune{}
// 			for _, r := range line {
// 				temp = append(temp, r)
// 			}
// 			pattern1 = append(pattern1, temp)
// 		case 2:
// 			temp := []rune{}
// 			for _, r := range line {
// 				temp = append(temp, r)
// 			}
// 			pattern2 = append(pattern2, temp)

// 		}
// 	}
// 	return pattern1, pattern2
// }
