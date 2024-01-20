package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	patterns := parse("input.txt")
	// patterns := parse("example.txt")
	for _, p := range patterns {
		for _, l := range p {
			fmt.Println(string(l))
		}
		fmt.Println()
	}

	record := make([]int, 0)
	for _, p := range patterns {
		kind, index := v_or_h(p)
		// fmt.Println(string(kind), index)
		if kind == 'v' {
			record = append(record, 0)
			record = append(record, index)
		}
		if kind == 'h' {
			record = append(record, 1)
			record = append(record, index)
		}
	}

	fmt.Println(record)

	output := 0
	for i, p := range patterns {
		kind, index := find_smudge(p, record[2*i], record[2*i+1])
		fmt.Println(string(kind), index, record[2*i], record[2*i+1])
		fmt.Println()
		if kind == 'v' {
			output += index
		}
		if kind == 'h' {
			output += 100 * index
		}
		if kind == 'x' {
			for _, l := range p {
				fmt.Println(string(l))
			}

			os.Exit(1)
		}
	}
	fmt.Println(output)
}

func find_smudge(pattern [][]rune, kind_int int, index int) (rune, int) {
	var kind rune
	if kind_int == 0 {
		kind = 'v'
	}
	if kind_int == 1 {
		kind = 'h'
	}

	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[0]); j++ {
			var new [][]rune
			for _, l := range pattern {
				var line []rune
				for _, c := range l {
					line = append(line, c)
				}
				new = append(new, line)
			}

			if new[i][j] == '#' {
				new[i][j] = '.'
			} else {
				new[i][j] = '#'
			}

			reflect_kind, reflect_index := v_or_h2(new)
			for i, _ := range reflect_index {
				// fmt.Println(string(reflect_kind[i]), reflect_index[i])
				if reflect_kind[i] == kind && reflect_index[i] == index {
					// fmt.Println("found old")
				} else {
					// fmt.Println("found new")
					return reflect_kind[i], reflect_index[i]
				}
			}
			// for _, l := range new {
			// 	fmt.Println(string(l))
			// }
			// fmt.Println()

			// for _, l := range pattern {
			// 	fmt.Println(string(l))
			// }
			// fmt.Println()
			// fmt.Println(string(new_kind), new_index)
			// fmt.Println(i, j)
		}
	}
	fmt.Println("smudge not found, impossible")
	return 'x', 0
}

func v_or_h2(pattern [][]rune) ([]rune, []int) {
	var recordhv []rune
	var record []int
	kind := 'v'
	for i := 0; i < len(pattern[0]); i++ {
		if index_reflect(pattern, i) {
			recordhv = append(recordhv, kind)
			record = append(record, i)
		}
	}
	kind = 'h'
	pattern = transpose(pattern)
	for i := 0; i < len(pattern[0]); i++ {
		if index_reflect(pattern, i) {
			recordhv = append(recordhv, kind)
			record = append(record, i)
		}
	}
	return recordhv, record

}
func v_or_h(pattern [][]rune) (rune, int) {
	kind := 'v'
	output := find_reflection(pattern)
	if output == 0 {
		output = find_reflection(transpose(pattern))
		kind = 'h'
	}
	if output == 0 {
		// fmt.Println("impossible case")
		// os.Exit(1)
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

func index_reflect(pattern [][]rune, index int) bool {
	if index == 0 || index == len(pattern[0]) {
		return false
	}
	index_r := index
	index_l := index - 1
	for index_r < len(pattern[0]) && index_l >= 0 {
		if !compare_lines(pattern, index_l, index_r) {
			return false
		}
		index_l--
		index_r++
	}
	return true
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
