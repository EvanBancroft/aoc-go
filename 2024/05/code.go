package main

import (
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
	var rules []string
	var updates []string
	var count int
	inputLines := strings.Split(input, "\n")

	for i := range inputLines {
		if strings.ContainsAny(inputLines[i], "|") {
			rules = append(rules, inputLines[i])
		}
		if strings.ContainsAny(inputLines[i], ",") {
			updates = append(updates, inputLines[i])
		}
	}

updateLoop:
	for _, update := range updates {
		nums := strings.Split(update, ",")
		valid := false
		for amountIndex, amount := range nums {
			for i := range rules {
				split := strings.Split(rules[i], "|")

				if amount == split[0] {
					if slices.Contains(nums[:amountIndex], split[1]) {
						valid = false
						continue updateLoop
					}
					valid = true

				}
			}
		}

		if valid {
			middle, _ := strconv.Atoi(nums[len(nums)/2])
			count += middle
		}

	}

	if part2 {
		var count int
		for _, update := range updates {
			nums := strings.Split(update, ",")
			valid := false
			for amountIndex, amount := range nums {
				for i := range rules {
					split := strings.Split(rules[i], "|")

					if amount == split[0] {
						if slices.Contains(nums[:amountIndex], split[1]) {

							rule1Index := slices.Index(nums, split[0])
							rule2Index := slices.Index(nums, split[1])

							nums[rule1Index], nums[rule2Index] = nums[rule2Index], nums[rule1Index]
							nums = resortAndValidate(nums, rules)
							valid = true
						}
					}
				}
			}

			if valid {
				middle, _ := strconv.Atoi(nums[len(nums)/2])
				count += middle
			}
		}
		return count
	}

	// solve part 1 here
	return count
}

func resortAndValidate(nums []string, rules []string) []string {
	reordered := false
	for amountIndex, amount := range nums {
		for i := range rules {
			split := strings.Split(rules[i], "|")

			// rule matches
			if amount == split[0] {
				// order is good
				if !slices.Contains(nums[:amountIndex], split[1]) {
					continue
				}
				// update is out of order
				if slices.Contains(nums[:amountIndex], split[1]) {

					rule1Index := slices.Index(nums, split[0])
					rule2Index := slices.Index(nums, split[1])

					nums[rule1Index], nums[rule2Index] = nums[rule2Index], nums[rule1Index]
					reordered = true
					continue
				}
			}
		}
	}
	// recursively resort if we changed something
	if reordered {
		resortAndValidate(nums, rules)
	}
	return nums
}
