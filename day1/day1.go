package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func spelledToInt(spelled string) int {
	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	output, ok := m[spelled]
	if !ok {
		output, _ := strconv.Atoi(spelled)
		// fmt.Println("input:", spelled, "output:", (output))
		return output
	}
	// fmt.Println("input:", spelled, "output:", (output))
	return output

}

func reverseString(input string) string {
	output := ""
	for _, c := range input {
		output = string(c) + output
	}
	return output
}

func main() {
	//part 1
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	output := 0
	re_string := "\\d"
	r, _ := regexp.Compile(re_string)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		match := r.FindString(line)
		// fmt.Println(match)

		match_b := r.FindString(reverseString(line))
		// fmt.Println(match_b)

		line_value, _ := strconv.Atoi(match + match_b)
		// fmt.Println(line_value)

		output += line_value
	}

	fmt.Println(output)

	readFile.Close()

	//part 2
	readFile, err = os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	output = 0
	re_string = "(?:zero|one|two|three|four|five|six|seven|eight|nine|\\d)"
	temp_b := reverseString("zero|one|two|three|four|five|six|seven|eight|nine")
	re_string_b := fmt.Sprintf("(?:%s|\\d)", temp_b)
	r, _ = regexp.Compile(re_string)
	r_b, _ := regexp.Compile(re_string_b)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		match := r.FindString(line)
		match_b := r_b.FindString(reverseString(line))
		match_b = reverseString(match_b)

		// fmt.Println(line)
		// fmt.Print(match)
		// fmt.Println(spelledToInt(match))
		// fmt.Print(match_b)
		// fmt.Println(spelledToInt(match_b))

		num := strconv.Itoa(spelledToInt(match))
		num_b := strconv.Itoa(spelledToInt(match_b))
		// fmt.Println(num, num_b)
		line_value, _ := strconv.Atoi(num + num_b)
		// fmt.Println(line_value)

		output += line_value
	}

	fmt.Println(output)

	readFile.Close()
}
