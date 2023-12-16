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

	// mappings_map := make(map[string][][]int)
	mappings_map := make([][][]int, 0)
	// output := 0
	var input []int
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
				// fmt.Println(map_name)

			} else {
				// fmt.Println(line)
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
	// fmt.Println(mappings_map)

	//parse
	start_end_map := make([][]start_end, 0)

	count := 0
	for _, mappings := range mappings_map {
		count++
		// if count > 2 {
		// 	break
		// }

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
		// fmt.Println(src_range)
		// fmt.Println(dst_range)

		start_end_map = append(start_end_map, src_range)
		start_end_map = append(start_end_map, dst_range)

	}
	for i := 0; i < len(start_end_map); i++ {
		// fmt.Println(start_end_map[i])
	}
	// test_range := combine_range(start_end_map[1], start_end_map[2])
	// test_list := []int{79, 14, 55, 13}
	test_list := input
	temp_list := []int{}
	for i := 0; i < len(start_end_map); i++ {
		for j := 0; j < len(test_list); j++ {
			test_traverse := traverse(test_list[j], start_end_map[i])
			// fmt.Println(test_traverse)
			temp_list = append(temp_list, test_traverse)
		}
		test_list = temp_list
		temp_list = []int{}
		fmt.Println(test_list)
		// if i == 2 {
		// 	break
		// }
		i++
	}

	min := test_list[0]
	for i := 0; i < len(test_list); i++ {
		if test_list[i] < min {
			min = test_list[i]
		}
	}
	fmt.Println(len(start_end_map))
	fmt.Println(min)

}

type start_end struct {
	start  int
	end    int
	target *start_end
}

func traverse(num int, mapping []start_end) int {

	// fmt.Println(num, mapping)

	var result int
	for i := 0; i < len(mapping); i++ {
		if mapping[i].start <= num && num < mapping[i].end {
			diff := num - mapping[i].start
			target := mapping[i].target
			result = target.start + diff
			return result
		}
	}
	return num
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

// func combine_range(src_range []start_end, dst_range []start_end) []start_end {
// 	fmt.Println(src_range)
// 	fmt.Println(dst_range)

// 	output_range := make([]start_end, 0)
// 	// for _, s_r := src_range{
// 	// 	pointer := s_r.start
// 	// 	for _, d_r := dst_range{
// 	// 		if pointer < d_r.end{
// 	// 			output_range = append(output_range, start_end{start:pointer, end:d_r.end})
// 	// 			pointer = d_r.end
// 	// 		}
// 	// 		}
// 	// 	}
// 	combine := append(src_range, dst_range...)
// 	sort_se(combine)
// 	fmt.Println(combine)
// 	combine_2 := make([]start_end, 0)
// 	for i, _ := range combine {
// 		if i+1 >= len(combine) {
// 			break
// 		}
// 		if combine[i].end > combine[i+1].start {
// 			// fmt.Print(combine[i])
// 			// fmt.Println(combine[i+1])

// 			var target *start_end
// 			if combine[i].target != nil {
// 				if combine[i+1].target != nil {
// 					fmt.Println("wtf")
// 				}
// 				target = combine[i].target
// 			} else {
// 				if combine[i+1].target == nil {
// 					fmt.Println("wtf")
// 				}
// 				target = combine[i+1].target
// 			}
// 			// if combine[i].target == nil {
// 			// 	if combine[i+1].target == nil {
// 			// 		fmt.Println("wtf")
// 			// 	} else {
// 			// 		target = combine[i+1].target
// 			// 	}
// 			// } else {
// 			// 	target = combine[i].target
// 			// }

// 			combine_2 = append(combine_2, start_end{start: combine[i].start, end: combine[i+1].start, target: target})

// 			var first start_end
// 			var last start_end
// 			if combine[i].end < combine[i+1].end {
// 				first = combine[i]
// 				last = combine[i+1]
// 			} else {
// 				first = combine[i+1]
// 				last = combine[i]
// 			}
// 			combine_2 = append(combine_2, start_end{start: combine[i+1].start, end: first.end, target: target})
// 			target = last.target
// 			combine_2 = append(combine_2, start_end{start: first.end, end: last.end, target: target})

// 		} else {
// 			combine_2 = append(combine_2, combine[i])
// 		}
// 	}
// 	sort_se(combine_2)
// 	fmt.Println(combine_2)
// 	combine_3 := make([]start_end, 0)
// 	for i := 0; i < len(combine_2); i++ {
// 		if combine_2[i].start == combine_2[i].end {
// 			continue
// 		}
// 		if i+1 < len(combine_2) {
// 			if combine_2[i].start == combine_2[i+1].start && combine_2[i].end == combine_2[i+1].end {
// 				var with_target = start_end{}
// 				fmt.Println(combine_2[i], combine_2[i+1])
// 				if combine_2[i].target != nil {
// 					with_target = combine_2[i]
// 				} else {
// 					with_target = combine_2[i+1]
// 				}
// 				fmt.Println(with_target)

// 				// if combine[i].target == nil {

// 				// 	if combine[i+1].target == nil {
// 				// 		fmt.Println("wtf")
// 				// 	} else {
// 				// 		with_target = combine[i]
// 				// 	}
// 				// } else {
// 				// 	with_target = combine[i+1]
// 				// }
// 				combine_3 = append(combine_3, with_target)
// 				i++
// 				continue
// 			}
// 		}

// 		combine_3 = append(combine_3, combine_2[i])
// 	}
// 	// fmt.Println(combine_3)
// 	sort_se(combine_3)
// 	fmt.Println(combine_3)
// 	return combine_3
// 	return output_range
// }
