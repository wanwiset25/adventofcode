package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// for each line
	// parse string
	// compare num and calc points
	// add up points
	// readFile, err := os.Open("example.txt")
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	winning_nums := make([][]int, 0)
	nums := make([]map[int]bool, 0)

	output := 0
	count := 0
	_ = count
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line_a := strings.Split(line, "|")
		w_nums_a, nums_a := line_a[0], line_a[1]
		w_nums_a = strings.Split(w_nums_a, ":")[1]

		w_nums_list := strings.Split(w_nums_a, " ")
		nums_list := strings.Split(nums_a, " ")

		// fmt.Println(w_nums_list)
		// fmt.Println(nums_list)

		count++

		w_nums_data_list := make([]int, 0)
		nums_map := make(map[int]bool)

		for _, e := range w_nums_list {
			if e != "" {
				e_int, _ := strconv.Atoi(e)
				w_nums_data_list = append(w_nums_data_list, e_int)
			}
		}
		for _, e := range nums_list {
			if e != "" {
				e_int, _ := strconv.Atoi(e)
				nums_map[e_int] = true
			}
		}

		winning_nums = append(winning_nums, w_nums_data_list)
		nums = append(nums, nums_map)

	}

	fmt.Println(winning_nums)
	fmt.Println(nums)

	for i, _ := range winning_nums {
		winning_num := winning_nums[i]
		num := nums[i]

		score := 0
		for j, _ := range winning_num {
			_, found := num[winning_num[j]]
			if found {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		output += score
	}

	fmt.Println(output)
}
