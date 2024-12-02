package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	rough_slice := strings.Fields(input)
	left_slice := []string{}
	right_slice := []string{}

	total_distance := 0

	for index, value := range rough_slice {
		if index%2 == 1 {
			right_slice = append(right_slice, value)
			continue
		}
		left_slice = append(left_slice, value)

	}
	slices.Sort(left_slice)
	slices.Sort(right_slice)

	for index, value := range left_slice {

		left_num, init_err := strconv.Atoi(value)
		if init_err != nil {
			fmt.Println("Conversion error:", init_err)
			continue
		}

		right_num, next_err := strconv.Atoi(right_slice[index])
		if next_err != nil {
			fmt.Println("Conversion error:", next_err)
			continue
		}

		distance := left_num - right_num

		if distance < 0 {
			distance = distance * -1
		}

		total_distance = total_distance + distance
	}

	if part2 {
		total_score := 0
		for _, value := range left_slice {
			matches := []string{}
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Conversion error:", err)
				continue
			}

			for _, right_value := range right_slice {
				if value == right_value {
					matches = append(matches, right_value)
				}
			}

			similarity_score := len(matches) * num

			total_score = total_score + similarity_score

		}

		return total_score
	}
	// solve part 1 here
	return total_distance
}
