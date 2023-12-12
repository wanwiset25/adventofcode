package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	//part 1
	// readFile, err := os.Open("example.txt")
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	engine := make([][]rune, 0)

	output := 0
	count := 0
	_ = count
	for fileScanner.Scan() {
		line := fileScanner.Text()

		line_slice := make([]rune, 0)
		for _, c := range line {
			line_slice = append(line_slice, c)
		}
		engine = append(engine, line_slice)
		// fmt.Println(string(line_slice))
	}
	// fmt.Println(engine)

	for i, _ := range engine {
		for j, _ := range engine[i] {
			if engine[i][j] == '*' {
				groups := find_groups(i, j, engine)
				// fmt.Println(i, j)
				if len(groups) == 2 {
					group_1, _ := strconv.Atoi(string(groups[0]))
					group_2, _ := strconv.Atoi(string(groups[1]))
					output += group_1*group_2
				}
				// for _, g := range groups {
				// 	fmt.Println(string(g))
				// }
			}
		}
		// os.Exit(1)
	}
	fmt.Println(output)
	// count++
	// if count > 10 {
	// 	os.Exit(1)
	// }
}

func find_groups(i int, j int, engine [][]rune) [][]rune {
	len_v := len(engine)
	len_h := len(engine[0])
	checked := make([][]bool, 0)
	groups := make([][]rune, 0)

	for _, l := range engine {
		line := make([]bool, 0)
		for _, c := range l {
			_ = c
			line = append(line, false)
		}
		checked = append(checked, line)
	}

	if i-1 >= 0 && j-1 >= 0 {
		if unicode.IsDigit(engine[i-1][j-1]) {
			if checked[i-1][j-1] == false {
				group := identify_group(i-1, j-1, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i-1][j-1] = true
	}

	if i-1 >= 0 && j+1 < len_v {
		if unicode.IsDigit(engine[i-1][j+1]) {
			if checked[i-1][j+1] == false {
				group := identify_group(i-1, j+1, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i-1][j+1] = true
	}

	if i+1 < len_h && j-1 >= 0 {
		if unicode.IsDigit(engine[i+1][j-1]) {
			if checked[i+1][j-1] == false {
				group := identify_group(i+1, j-1, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i+1][j-1] = true
	}

	if i+1 < len_h && j+1 < len_v {
		if unicode.IsDigit(engine[i+1][j+1]) {
			if checked[i+1][j+1] == false {
				group := identify_group(i+1, j+1, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i+1][j+1] = true
	}

	if i-1 >= 0 {
		if unicode.IsDigit(engine[i-1][j]) {
			if checked[i-1][j] == false {
				group := identify_group(i-1, j, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i-1][j] = true
	}

	if i+1 < len_h {
		if unicode.IsDigit(engine[i+1][j]) {
			if checked[i+1][j] == false {
				group := identify_group(i+1, j, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i+1][j] = true
	}

	if j-1 >= 0 {
		if unicode.IsDigit(engine[i][j-1]) {
			if checked[i][j-1] == false {
				group := identify_group(i, j-1, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i][j-1] = true
	}

	if j+1 < len_v {
		if unicode.IsDigit(engine[i][j+1]) {
			if checked[i][j+1] == false {
				group := identify_group(i, j+1, engine, checked)
				groups = append(groups, group)
			}
		}
	} else {
		//		checked[i][j+1] = true
	}

	return groups
}

func identify_group(i int, j int, engine [][]rune, checked [][]bool) []rune {
	// len_v := len(engine)
	// len_h := len(engine[0])

	group := make([]rune, 0)
	group = append(group, engine[i][j])

	checked[i][j] = true

	group = identify_left(i, j-1, engine, checked, group)
	group = identify_right(i, j+1, engine, checked, group)

	// fmt.Println(group)

	return group
}

func identify_left(i int, j int, engine [][]rune, checked [][]bool, group []rune) []rune {
	len_v := len(engine)
	len_h := len(engine[0])

	if i >= 0 && i < len_v && j >= 0 && j < len_h {
		checked[i][j] = true
		if unicode.IsDigit(engine[i][j]) {
			group = append([]rune{engine[i][j]}, group...)
			// fmt.Println(group)
			group = identify_left(i, j-1, engine, checked, group)
		} else {
			return group
		}
	} else {
		return group
	}
	return group

}

func identify_right(i int, j int, engine [][]rune, checked [][]bool, group []rune) []rune {
	len_v := len(engine)
	len_h := len(engine[0])
	if i >= 0 && i < len_v && j >= 0 && j < len_h {
		checked[i][j] = true
		if unicode.IsDigit(engine[i][j]) {
			group = append(group, engine[i][j])
			// fmt.Println(group)
			group = identify_right(i, j+1, engine, checked, group)
		} else {
			return group
		}
	} else {
		return group
	}
	return group

}
