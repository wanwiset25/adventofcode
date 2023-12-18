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

	output := 0
	for i := 0; i < len(time); i++ {
		count := 0
		for t := 0; t <= time[i]; t++ {
			traveled := t * (time[i] - t)
			if traveled > distance[i] {
				count++
			}
		}
		if output == 0 {
			output = count
		} else {
			output *= count
		}
	}
	fmt.Println(output)
}

func parse(filename string) (time []int, distance []int) {
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

	time_str_list := strings.Split(strings.Split(time_str, ":")[1], " ")
	distance_str_list := strings.Split(strings.Split(distance_str, ":")[1], " ")
	fmt.Println(time_str_list, distance_str_list)

	for _, w := range time_str_list {
		i, err := strconv.Atoi(w)
		if err != nil {
		} else {
			time = append(time, i)
		}
	}
	for _, w := range distance_str_list {
		i, err := strconv.Atoi(w)
		if err != nil {
		} else {
			distance = append(distance, i)
		}
	}
	return time, distance
}
