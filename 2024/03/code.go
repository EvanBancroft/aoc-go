package main

import (
	"regexp"
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
	if part2 {
		re := regexp.MustCompile(`(mul[(]\d{1,3},\d{1,3}[)])|(do[(][)])|(don't[(][)])`)
		matches := re.FindAll([]byte(input), -1)
		running_total := 0
		enable_counter := true

		for _, value := range matches {
			str := string(value)

			if strings.Contains(str, "don't") {
				enable_counter = false
				continue
			}
			if strings.Contains(str, "do()") {
				enable_counter = true
				continue
			}
			if !enable_counter {
				continue
			}

			ints := make([]int, 0)
			re := regexp.MustCompile(`(\d{1,3})`)
			numbers := re.FindAll([]byte(str), -1)

			for i := range numbers {
				byteToInt, _ := strconv.Atoi(string(numbers[i]))
				ints = append(ints, byteToInt)
			}

			running_total = running_total + ints[0]*ints[1]
		}

		return running_total
	}

	re := regexp.MustCompile(`(mul[(]\d{1,3},\d{1,3}[)])`)
	matches := re.FindAll([]byte(input), -1)

	running_total := 0

	for _, value := range matches {
		str := string(value)
		re := regexp.MustCompile(`(\d{1,3})`)
		numbers := re.FindAll([]byte(str), -1)

		ints := make([]int, 0)
		for i := range numbers {
			byteToInt, _ := strconv.Atoi(string(numbers[i]))
			ints = append(ints, byteToInt)
		}

		running_total = running_total + ints[0]*ints[1]

	}
	//
	// solve part 1 here
	return running_total
}
