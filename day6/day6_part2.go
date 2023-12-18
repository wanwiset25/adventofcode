package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// time, distance := parse("example.txt")
	time, distance := parse("input.txt")

	_, _ = time, distance
	fmt.Println(time, distance)

	output := search(time, distance)
	fmt.Println(output)
}
func brute(time int, distance int) int {
	count := 0
	for t := 0; t <= time; t++ {
		if t*(time-t) > distance {
			count++
		}
	}
	return count
}
func search(time int, distance int) int {
	found := false
	min := 0
	max := time
	cur := time / 2
	prev := 0
	for !found {
		if cur*(time-cur) > distance {
			prev = cur
			max = cur
			cur = (min + max) / 2
		} else if cur*(time-cur) < distance {
			prev = cur
			min = cur
			cur = (min + max) / 2
		} else {
			break
		}
		if prev == cur {
			break
		}
		fmt.Println(min, max, prev, cur)
	}
	return time - 2*cur -1
	// 42250895
}

func parse(filename string) (time int, distance int) {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	time_str := fileScanner.Text()
	fileScanner.Scan()
	distance_str := fileScanner.Text()

	time_str_list := strings.Split(time_str, ":")[1]
	distance_str_list := strings.Split(distance_str, ":")[1]
	time_str_list = strings.ReplaceAll(time_str_list, " ", "")
	distance_str_list = strings.ReplaceAll(distance_str_list, " ", "")
	// fmt.Println(time_str_list, distance_str_list)
	time, _ = strconv.Atoi(time_str_list)
	distance, _ = strconv.Atoi(distance_str_list)

	return time, distance
}
