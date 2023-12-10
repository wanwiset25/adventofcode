package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	red   int
	green int
	blue  int
}

func main() {
	//part 1
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	m := make(map[int]*Game)

	output := 0
	count := 0 
	_ = count 
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		line_b := strings.Split(line, ":")
		// fmt.Println(line_b[0], line_b[1])

		game_num, _ := strconv.Atoi(strings.Split(line_b[0], " ")[1])
		// fmt.Println(game_num)

		line_c := strings.Split(line_b[1], ";")
		for _, c := range line_c {
			// fmt.Println(c)
			line_d := strings.Split(c, ",")
			for _, d := range line_d {
				// fmt.Println(d)
				line_e := strings.Split(d, " ")
				num, _ := strconv.Atoi(line_e[1])
				color := line_e[2]
				_, _ = num, color
				// fmt.Println(num, color)

				_, exist := m[game_num]
				if !exist {
					m[game_num] = &Game{}
				}

				switch color {
				case "red":
					if num > m[game_num].red {
						m[game_num].red = num
					}
				case "green":
					if num > m[game_num].green {
						m[game_num].green = num
					}
				case "blue":
					if num > m[game_num].blue {
						m[game_num].blue = num
					}
				}
			}
		}
	}
	for k, v := range m {
		// fmt.Println(k, v)
		possible := true
		if v.red > 12 {
			possible = false
		}
		if v.green > 13 {
			possible = false
		}
		if v.blue > 14 {
			possible = false
		}
		if possible {
			output += k
			// fmt.Println(k)
		}
	}
	fmt.Println(output)


	// part 2
	output2 := 0
	for _, v := range m {
		power := v.red*v.green*v.blue
		output2 += power
	}
	fmt.Println(output2)



	// count++
	// if count > 10 {
	// os.Exit(1)
	// }
}
