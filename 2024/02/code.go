package main

import (
	"log"
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
	reports := strings.Split(input, "\n")

	safe_amount := 0
	dampened_safe_amount := 0

	for _, report := range reports {
		separated_report := strings.Fields(report)

		ints := make([]int, len(separated_report))

		for i, s := range separated_report {
			ints[i] = numbers(s)
		}

		if len(ints) == 0 {
			break
		}

		if check_report_safety(ints) {
			safe_amount = safe_amount + 1
		}

	}
	if part2 {
		for _, report := range reports {
			separated_report := strings.Fields(report)

			ints := make([]int, len(separated_report))

			for i, s := range separated_report {
				ints[i] = numbers(s)
			}

			if len(ints) == 0 {
				break
			}

			is_safe := check_report_safety(ints)
			if !is_safe && check_dampened_safety(ints) {
				dampened_safe_amount = dampened_safe_amount + 1
			}
			if is_safe {
				dampened_safe_amount = dampened_safe_amount + 1
			}

		}
		return dampened_safe_amount
	}
	// solve part 1 here

	return safe_amount
}

func check_dampened_safety(record []int) bool {
	do_any_work := false

	for i := range record {

		new_record := RemoveIndex(record, i)

		if check_report_safety(new_record) {
			do_any_work = true
			break
		}
		continue
	}

	return do_any_work
}

func check_report_safety(report []int) bool {
	is_safe := false

	is_descending := false

	for i, amount := range report {
		next_amount := report[i+1]

		if !safe_increase(amount, next_amount) {
			is_safe = false
			break
		}

		if i == 0 && amount > next_amount {
			is_descending = true
			continue
		}
		if i == 0 && amount < next_amount {
			is_descending = false
			continue
		}

		if is_descending {
			if amount > next_amount {
				if i+2 == len(report) {
					is_safe = true
					break
				}
			} else {
				is_safe = false
				break
			}
		} else {
			if amount < next_amount {
				if i+2 == len(report) {
					is_safe = true
					break
				}
			} else {
				is_safe = false
				break
			}
		}

	}

	return is_safe
}

func numbers(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)
		return 0
	}
	return num
}

func safe_increase(input1 int, input2 int) bool {
	distance := input1 - input2

	if distance < 0 {
		distance = distance * -1
	}

	if distance < 1 || distance > 3 {
		return false
	}
	return true
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
