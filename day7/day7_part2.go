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
	// hands := parse("example.txt")
	hands := parse("input.txt")
	sort_hands(hands)
	for i := 0; i < len(hands); i++ {
		hands[i].rank = len(hands) - i
	}
	fmt.Println(hands)

	output := 0
	for i := 0; i < len(hands); i++ {
		output += hands[i].rank * hands[i].bid
	}
	fmt.Println(output)
}

type hand struct {
	value    string
	win_type int
	bid      int
	rank     int
}

func parse(filename string) []hand {
	output := []hand{}

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		hand_str := fileScanner.Text()
		hand_str_list := strings.Split(hand_str, " ")
		hand_value := hand_str_list[0]
		bid, _ := strconv.Atoi(hand_str_list[1])

		hand := hand{value: hand_value, bid: bid}

		hand.find_win_type()

		output = append(output, hand)

	}
	return output
}

func sort_hands(hands []hand) {
	sort.Slice(hands, func(i, j int) bool {
		var ans bool
		if hands[i].win_type != hands[j].win_type {
			ans = hands[i].win_type < hands[j].win_type
		} else {
			cmp_res := 0
			for k := 0; k < 5; k++ {
				cmp_res = compare_cards(rune(hands[i].value[k]), rune(hands[j].value[k]))
				if cmp_res == 1 {
					ans = true
					break
				}
				if cmp_res == 2 {
					ans = false
					break
				}
			}
			if cmp_res == 0 {
				fmt.Println("impossible case", hands[i], hands[j])
			}
		}
		return ans
	})

}

//	func sort_se(se_list []start_end) {
//		sort.Slice(se_list, func(i, j int) bool {
//			ans := se_list[i].start < se_list[j].start
//			if se_list[i].start == se_list[j].start {
//				ans = se_list[i].end < se_list[j].end
//			}
//			return ans
//		})
func compare_cards(i rune, j rune) int {
	value := map[rune]int{
		'A': 1,
		'K': 2,
		'Q': 3,
		'J': 15, // special joker
		'T': 5,
		'9': 6,
		'8': 7,
		'7': 8,
		'6': 9,
		'5': 10,
		'4': 11,
		'3': 12,
		'2': 13,
		'1': 14,
	}
	if value[i] < value[j] {
		return 1
	}
	if value[i] > value[j] {
		return 2
	}
	if value[i] == value[j] {
		return 3
	}
	fmt.Println("impossible case", i, j)
	return 0
}

func (h *hand) find_win_type() {
	chars := map[rune]int{}
	max := 0
	j_count := 0
	for _, c := range h.value {
		if c == 'J' {
			j_count++
			continue
		}
		_, ok := chars[c]
		if !ok {
			chars[c] = 1
			if chars[c] > max {
				max++
			}
		} else {
			chars[c]++
			if chars[c] > max {
				max++
			}
		}
	}

	max += j_count

	// fmt.Println(h.value, chars, max)
	if max == 5 {
		h.win_type = 1 //five of a kind
		return
	}
	if max == 4 {
		h.win_type = 2 //four of a kind
		return
	}
	if max == 3 && len(chars) == 2 {
		h.win_type = 3 //full house
		return
	}
	if max == 3 && len(chars) == 3 {
		h.win_type = 4 //three of a kind
		return
	}
	if max == 2 && len(chars) == 3 {
		h.win_type = 5 //two pair
		return
	}
	if max == 2 && len(chars) == 4 {
		h.win_type = 6 //one pair
		return
	}
	if max == 1 {
		h.win_type = 7
		return
	}
	fmt.Println("impossible case!", h.value, chars)
	return
}



// 1 J
//highcard > pair
//one pair > three
//three > four
// two pair > full
//three > four
//four > five