package main

import (
	"fmt"
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
		lines := formatInput(input)
		var count int
		for _, line := range lines {
			testValue, vars := formatValues(line)
			validOption := false

			for perm := range GenerateOperatorPermutationsChannel(len(vars)-1, true) {
				if doTheMath(vars, testValue, perm) {
					validOption = true
					break
				}
			}

			if validOption {
				count = count + testValue
			}
		}
		return count

	}
	lines := formatInput(input)
	var count int

	for _, line := range lines {
		testValue, vars := formatValues(line)
		validOption := false

		for perm := range GenerateOperatorPermutationsChannel(len(vars)-1, false) {
			if doTheMath(vars, testValue, perm) {
				validOption = true
				break
			}
		}

		if validOption {
			count = count + testValue
		}
	}
	return count
}

func doTheMath(vars []int, testValue int, perm string) bool {
	splitPerm := strings.Split(perm, "")
	var total int
	var adjusted int
	for i := range splitPerm {
		i = i + adjusted
		if i == len(splitPerm) {
			break
		}

		if i == 0 {
			if splitPerm[i] == "+" {
				total = vars[0] + vars[1]
				continue
			}
			if splitPerm[i] == "*" {
				total = vars[0] * vars[1]
				continue
			}
			if splitPerm[i] == "|" {
				combinedInt, err := strconv.Atoi(strconv.Itoa(vars[0]) + strconv.Itoa(vars[1]))
				if err != nil {
					fmt.Println("Error opening file:", err)
					return false
				}
				total = combinedInt
				continue
			}
		}

		if splitPerm[i] == "+" {
			total = total + vars[i+1]
		}
		if splitPerm[i] == "*" {
			total = total * vars[i+1]
		}
		if splitPerm[i] == "|" {
			combinedInt, err := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(vars[i+1]))
			if err != nil {
				fmt.Println("Error opening file:", err)
				return false
			}
			total = combinedInt
		}
	}
	if total == testValue {
		return true
	}
	return false
}

// GenerateOperatorPermutationsChannel provides a channel-based generator
func GenerateOperatorPermutationsChannel(length int, combo bool) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)

		var generate func(string)
		generate = func(current string) {
			if len(current) == length {
				output <- current
				return
			}
			// Explicitly handle single-length case
			if length == 1 {
				output <- "+"
				output <- "*"
				if combo {
					output <- "|"
				}
				return
			}

			generate(current + "+")
			generate(current + "*")
			if combo {
				generate(current + "|")
			}
		}

		generate("")
	}()

	return output
}

func formatValues(line string) (int, []int) {
	items := strings.Split(line, ":")
	testValue, _ := strconv.Atoi(items[0])
	stringVars := strings.Split(strings.Trim(items[1], " "), " ")
	vars := make([]int, 0)

	for i := range stringVars {
		stringVars[i] = strings.Trim(stringVars[i], " ")
		int, _ := strconv.Atoi(stringVars[i])
		vars = append(vars, int)
	}

	return testValue, vars
}

func formatInput(input string) []string {
	rawSplit := strings.Split(input, "\n")
	finalSplit := make([]string, 0)

	for i := range rawSplit {
		if len(rawSplit[i]) > 0 {
			finalSplit = append(finalSplit, rawSplit[i])
		}
	}

	return finalSplit
}
