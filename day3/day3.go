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
		is_group := false
		is_part := false
		part_num := ""
		for j, _ := range engine[i] {
			is_digit := unicode.IsDigit(engine[i][j])
			if is_digit {
				part_num += string(engine[i][j])
				if !is_group {
					is_group = true
					is_part = false
				}

				if !is_part {
					is_part = check(i, j, engine)
					// fmt.Println(string(engine[i][j]), is_part)
				}

				if j+1 == len(engine[i]) {	//end of line logic
					is_group = false
					if is_part {
						part_int, _ := strconv.Atoi(part_num)
						is_part = false
						output += part_int
						// fmt.Println(part_int)
						part_num = ""
					}
				}
			} else {
				if is_group {
					is_group = false
					if is_part {
						part_int, _ := strconv.Atoi(part_num)
						is_part = false
						output += part_int
						// fmt.Println(part_int)
					}
					part_num = ""
				}
			}
		}
	}
	fmt.Println(output)
	// count++
	// if count > 10 {
	// 	os.Exit(1)
	// }
}

func check(i int, j int, engine [][]rune) bool {
	len_v := len(engine)
	len_h := len(engine[0])

	// i-1 >= 0
	// i+1 < len_h

	// j-1 >= 0
	// j+1 < len_v

	if i-1 >= 0 && j-1 >= 0 {
		if !unicode.IsDigit(engine[i-1][j-1]) {
			if engine[i-1][j-1] != '.' {
				return true
			}
		}
	}
	if i-1 >= 0 && j+1 < len_v {
		if !unicode.IsDigit(engine[i-1][j+1]) {
			if engine[i-1][j+1] != '.' {
				return true
			}
		}
	}
	if i+1 < len_h && j-1 >= 0 {
		if !unicode.IsDigit(engine[i+1][j-1]) {
			if engine[i+1][j-1] != '.' {
				return true
			}
		}
	}
	if i+1 < len_h && j+1 < len_v {
		if !unicode.IsDigit(engine[i+1][j+1]) {
			if engine[i+1][j+1] != '.' {
				return true
			}
		}
	}
	if i-1 >= 0 {
		if !unicode.IsDigit(engine[i-1][j]) {
			if engine[i-1][j] != '.' {
				return true
			}
		}
	}
	if i+1 < len_h {
		if !unicode.IsDigit(engine[i+1][j]) {
			if engine[i+1][j] != '.' {
				return true
			}
		}
	}
	if j-1 >= 0 {
		if !unicode.IsDigit(engine[i][j-1]) {
			if engine[i][j-1] != '.' {
				return true
			}
		}
	}
	if j+1 < len_v {
		if !unicode.IsDigit(engine[i][j+1]) {
			if engine[i][j+1] != '.' {
				return true
			}
		}
	}
	return false
}
