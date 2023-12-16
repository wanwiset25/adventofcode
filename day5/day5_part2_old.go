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

	readFile, err := os.Open("example.txt")
	// readFile, err := os.Open("input.txt")
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

	test_list := []int{79, 14, 55, 13}
	// test_list := []int{82, 5, 55, 13}
	// test_list := input
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
		fmt.Println(map_name_list[k])
		fmt.Println(test_range)

		k++
		test_range = transform_range(test_range, start_end_map[i])
		i++
		if i >= 4 {
			break
		}
	}

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

func transform_range(input []start_end, mapping []start_end) []start_end {
	sort_se(input)
	sort_se(mapping)
	fmt.Println(input)
	fmt.Println(mapping)

	output := []start_end{}
	i := 0
	j := 0

	for i < len(input) && j < len(mapping) {
		if input[i].end > mapping[j].start && input[i].start < mapping[j].end {
			// var start_first start_end
			// var start_last start_end
			// var end_first start_end
			// var end_last start_end
			// if input[i].start < mapping[j].start {
			// 	start_first = input[i]
			// 	start_last = mapping[j]
			// } else {
			// 	start_first = mapping[j]
			// 	start_last = input[i]
			// }
			// if input[i].end < mapping[j].end {
			// 	end_first = input[i]
			// 	end_last = mapping[j]
			// } else {
			// 	end_first = mapping[j]
			// 	end_last = input[i]
			// }

			cut_points := []int{input[i].start, input[i].end, mapping[j].start, mapping[j].end}
			sort.Ints(cut_points)
			fmt.Println(cut_points)

			r_1 := start_end{start: cut_points[0], end: cut_points[1]}
			r_2 := start_end{start: cut_points[1], end: cut_points[2]}
			r_3 := start_end{start: cut_points[2], end: cut_points[3]}
			if cut_points[0] >= input[i].start && cut_points[1] <= input[i].end {
				temp1, err := transform_se(r_1, input[i], mapping[j])
				if !err {
					if temp1.start != temp1.end {
						output = append(output, temp1)
					}
				}
			}
			if cut_points[1] >= input[i].start && cut_points[2] <= input[i].end {
				temp2, err := transform_se(r_2, input[i], mapping[j])
				if !err {
					if temp2.start != temp2.end {
						output = append(output, temp2)
					}
				}
			}
			if cut_points[2] >= input[i].start && cut_points[3] <= input[i].end {
				temp3, err := transform_se(r_3, input[i], mapping[j])
				if !err {
					if temp3.start != temp3.end {
						output = append(output, temp3)
					}
				}
			}

		} else {
			// output = append(output, input[i])
			// fmt.Println("manual append", input[i])
		}
		if input[i].end > mapping[j].end {
			j++
		} else {
			i++
		}
		// fmt.Println(i,j)
	}
	
	i++
	for i < len(input) {
		fmt.Println("append", input[i])
		output = append(output, input[i])
		i++
	}
	output2 := []start_end{}
	for i := 0; i < len(output); i++ {
		if len(output2) > 0 {
			if output2[len(output2)-1] == output[i] {
				continue
			}
		}
		output2 = append(output2, output[i])
		sort_se(output2)
	}

	sort_se(output2)
	fmt.Println("output2", output2)
	return output2
}

func transform_se(cut start_end, input start_end, mapping start_end) (se start_end, err bool) {
	// if mapping[i].start <= num && num < mapping[i].end {
	// 		diff := num - mapping[i].start
	// 		target := mapping[i].target
	// 		result = target.start + diff
	// 		return result
	// 	}
	// start_end{start: mapping.target.start, end: mapping.target.end}

	// fmt.Println(cut, input, mapping)
	fmt.Println(cut, mapping)

	diff := cut.end - cut.start
	if cut.start == mapping.start {
		if cut.end == mapping.end {
			return start_end{start: mapping.target.start, end: mapping.target.end}, false
		} else if cut.end < mapping.end {
			return start_end{start: mapping.target.start, end: mapping.target.start + diff}, false
		} else {
			fmt.Println("wtf1")
		}

	} else if cut.start > mapping.start {
		if cut.end == mapping.end {
			return start_end{start: mapping.target.end - diff, end: mapping.target.end}, false
		} else if cut.end < mapping.end {
			start_diff := cut.start - mapping.start
			return start_end{start: mapping.target.start + start_diff, end: mapping.target.start + start_diff + diff}, false
		} else {
			// if cut.start >= input.start && cut.end <= input.end {
			// return start_end{start: cut.start, end: cut.end}, false
			// }

			fmt.Println("wtf2", cut, mapping)
		}
	} else {
		if cut.start >= input.start && cut.end <= input.end {
			return start_end{start: cut.start, end: cut.end}, false
		}
		fmt.Println("wtf3", cut)
	}
	return start_end{}, true
}
