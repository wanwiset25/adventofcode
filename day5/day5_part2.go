package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// first create all mappings
	// how to keep track of min? we need to use heap? (unfamiliar, prob just use regular variables)
	// 	- no need, we can collapse the mappings
	// after all mappings created, reduce one by one until seed-location map
	// 	- make a function for this

	// readFile, err := os.Open("example.txt")
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	mappings_map := make([][][]int, 0)
	map_name_list := []string{}

	var input []int
	_ = input
	line_num := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line_num++

		if line_num == 1 {
			seeds := line
			_ = seeds
			seeds2 := strings.Split(line, ":")[1]
			seeds3 := strings.Split(seeds2, " ")
			seeds4 := []int{}
			for _, s := range seeds3 {
				if s == "" || s == " " {
					continue
				}
				seed_int, _ := strconv.Atoi(s)
				seeds4 = append(seeds4, seed_int)
				input = seeds4
			}

			continue
		}

		inner_line_num := 0
		map_name := ""
		_ = map_name
		mapping_list := make([][]int, 0)
		for line != "" {
			inner_line_num++
			if inner_line_num == 1 {
				title := line
				map_name = strings.Split(title, " ")[0]
				map_name_list = append(map_name_list, map_name)

			} else {
				mapping_str := strings.Split(line, " ")
				mapping_int := make([]int, 3)

				for i, num_str := range mapping_str {
					mapping_int[i], _ = strconv.Atoi(num_str)
				}
				mapping_list = append(mapping_list, mapping_int)
			}

			// fmt.Println(line)
			fileScanner.Scan()
			line = fileScanner.Text()
			line_num++
		}
		if len(mapping_list) != 0 {
			mappings_map = append(mappings_map, mapping_list)
		}

	}

	//parse
	start_end_map := make([][]start_end, 0)

	count := 0
	for _, mappings := range mappings_map {
		count++

		src_range := make([]start_end, 0)
		dst_range := make([]start_end, 0)
		for _, list := range mappings {
			dst_start := list[0]
			src_start := list[1]
			range_num := list[2]

			dst_se := start_end{start: dst_start, end: dst_start + range_num}
			src_se := start_end{start: src_start, end: src_start + range_num, target: &dst_se}

			src_range = append(src_range, src_se)
			dst_range = append(dst_range, dst_se)
		}

		sort_se(src_range)
		sort_se(dst_range)

		start_end_map = append(start_end_map, src_range)
		start_end_map = append(start_end_map, dst_range)

	}
	for i := 0; i < len(start_end_map); i++ {
		// fmt.Println(start_end_map[i])
	}
	fmt.Println(map_name_list)

	// test_list := []int{79, 14, 55, 13}
	// test_list := []int{82, 5, 55, 13}
	test_list := input
	test_range := []start_end{}
	i := 0
	for i < len(test_list) {
		se := start_end{start: test_list[i], end: test_list[i] + test_list[i+1]}
		test_range = append(test_range, se)
		i += 2
	}

	fmt.Println(test_range)
	k := 0
	for i := 0; i < len(start_end_map); i++ {
		// fmt.Println(map_name_list[k])
		// fmt.Println(test_range)

		k++
		test_range = new_transform_range(test_range, start_end_map[i])
		test_range = process(test_range)
		i++
		// if i >= 14 {
			// break
		// }
	}
	fmt.Println()
	fmt.Println(test_range)
}

type start_end struct {
	start  int
	end    int
	target *start_end
}

func sort_se(se_list []start_end) {
	sort.Slice(se_list, func(i, j int) bool {
		ans := se_list[i].start < se_list[j].start
		if se_list[i].start == se_list[j].start {
			ans = se_list[i].end < se_list[j].end
		}
		return ans
	})
}

func process(input []start_end) []start_end {
	output := []start_end{}
	i := 0
	sort_se(input)
	for i < len(input) {
		if i+1 < len(input) {

			if input[i] == input[i+1] {
				i++
				continue
			}
		}
		output = append(output, input[i])
		i++

	}
	// output2:=[]start_end{}
	fmt.Println(output)
	i = 0
	for i < len(output) {
		if i+1 < len(output) {
			// fmt.Println(i, output[i], output[i+1])
			if output[i].end >= output[i+1].start {
				remove := output[i]
				remove2 := output[i+1]
				output = append(output[:i], output[i+1:]...)
				output = append(output[:i], output[i+1:]...)
				end := 0
				if remove.end > remove2.end {
					end = remove.end
				} else {
					end = remove2.end
				}
				add := start_end{start: remove.start, end: end}
				output = append(output[:i+1], output[i:]...)
				output[i] = add
				continue
			}
		}
		i++
	}
	fmt.Println(output)
	return output
}

func new_transform_range(input []start_end, mapping []start_end) []start_end {
	sort_se(input)
	sort_se(mapping)
	fmt.Println("transform", input, mapping)
	remove := []start_end{}
	output := []start_end{}
	// remain := []start_end{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(mapping); j++ {
			n := input[i]
			m := mapping[j]
			if n.end > m.start && n.start < m.end {
				points := []int{n.start, n.end, m.start, m.end}
				sort.Ints(points)
				for k := 0; k < 3; k++ {
					if points[k] >= m.start && points[k+1] <= m.end {
						if points[k] >= n.start && points[k+1] <= n.end {
							transform := start_end{start: m.target.start + (points[k] - m.start), end: m.target.start + (points[k] - m.start) + (points[k+1] - points[k])}
							output = append(output, transform)
							remove = append(remove, start_end{start: points[k], end: points[k+1]})
							// fmt.Println("output", output)
							// fmt.Println("remove", remove)
						}
					}
				}
			}
		}
	}
	remain := input
	for i := 0; i < 10; i++ {
		remain = clean_remaining_input(remain, mapping, remove)
		// fmt.Println(remain)
	}
	// prev := []start_end{}
	// equal := false
	// for !equal {
	// }

	output = append(output, remain...)
	return output
}
func clean_remaining_input(input []start_end, mapping []start_end, remove []start_end) []start_end {
	// fmt.Println("clean_input", input, remove, mapping )
	remaining_total := []start_end{}
	for i := 0; i < len(input); i++ {
		remaining_input := []start_end{}
		for j := 0; j < len(remove); j++ {
			n := input[i]
			r := remove[j]
			if n.start < r.end && n.end > r.start {
				if n.start < r.start && n.end <= r.end {
					remain := start_end{start: n.start, end: r.start}
					remaining_input = append(remaining_input, remain)
					// fmt.Println("appending", remain)
				}
				if n.start < r.start && n.end > r.end {
					remain := start_end{start: n.start, end: r.start}
					remain2 := start_end{start: r.end, end: n.end}
					remaining_input = append(remaining_input, remain)
					remaining_input = append(remaining_input, remain2)
					// fmt.Println("appending", remain)
					// fmt.Println("appending", remain2)

				}
				if n.start >= r.start && n.end <= r.end {
					// no remain
				}
				if n.start >= r.start && n.end > r.end {
					remain := start_end{start: r.end, end: n.end}
					remaining_input = append(remaining_input, remain)
					// fmt.Println("appending", remain)
				}
			}
		}
		found := 0
		for j := 0; j < len(mapping); j++ {
			n := input[i]
			m := mapping[j]
			if n.start < m.end && n.end > m.start {
				// fmt.Println("found == 1")
				found = 1
			}
		}
		if found == 0 {
			// fmt.Println("append remaining total", input[i])
			remaining_total = append(remaining_total, input[i])
		}
		remaining_total = append(remaining_total, remaining_input...)
	}
	return remaining_total
}

