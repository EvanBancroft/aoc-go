package main

import (
	"fmt"
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
		matrix, unique := formatInput(input)
		antinodes := make(map[string]bool)

		for _, letter := range unique {
			locationCombos := genPermutations(letter.locations, 2)

			for _, coords := range locationCombos {
				rightAntenna := make([]int, 2)
				leftAntenna := make([]int, 2)
				if coords[0][0] > coords[1][0] {
					rightAntenna = coords[0]
					leftAntenna = coords[1]
				}
				if coords[0][0] < coords[1][0] {
					rightAntenna = coords[1]
					leftAntenna = coords[0]
				}

				result := recoverAllAntinodes(leftAntenna, rightAntenna, matrix)

				for k, v := range result {
					antinodes[k] = v
				}
			}
		}
		return len(antinodes)
	}
	matrix, unique := formatInput(input)
	antinodes := make(map[string]bool)

	for _, letter := range unique {
		locationCombos := genPermutations(letter.locations, 2)

		for _, coords := range locationCombos {
			var leftAntinode []int
			var rightAntinode []int

			rightAntenna := make([]int, 2)
			leftAntenna := make([]int, 2)
			diff := make([]int, 2)

			if coords[0][0] > coords[1][0] {
				rightAntenna = coords[0]
				leftAntenna = coords[1]
			}
			if coords[0][0] < coords[1][0] {
				rightAntenna = coords[1]
				leftAntenna = coords[0]
			}

			diff = getDiff(rightAntenna, leftAntenna)

			for i := range diff {
				rightAntinode = append(rightAntinode, rightAntenna[i]+diff[i])
				leftAntinode = append(leftAntinode, leftAntenna[i]-diff[i])
			}

			if isInBounds(leftAntinode, matrix) {
				if matrix[-leftAntinode[1]][leftAntinode[0]] == "." {
					matrix[-leftAntinode[1]][leftAntinode[0]] = "#"
				}
				antinodes[arrayToString(leftAntinode, "")] = true
			}
			if isInBounds(rightAntinode, matrix) {
				if matrix[-rightAntinode[1]][rightAntinode[0]] == "." {
					matrix[-rightAntinode[1]][rightAntinode[0]] = "#"
				}
				antinodes[arrayToString(rightAntinode, "")] = true
			}

		}
	}

	if len(unique) < 5 {
		for i := range matrix {
			fmt.Println(matrix[i])
		}
	}
	return len(antinodes)
}

func recoverAllAntinodes(leftAntenna []int, rightAntenna []int, matrix [][]string) map[string]bool {
	antinodes := make(map[string]bool)
	diff := getDiff(rightAntenna, leftAntenna)

	antinodes[arrayToString(leftAntenna, "")] = true

	var recover func(path []int, result map[string]bool, matrix [][]string, leftStart []int, rightStart []int)
	recover = func(path []int, result map[string]bool, matrix [][]string, leftStart []int, rightStart []int) {
		nextLeft := make([]int, 2)
		nextRight := make([]int, 2)
		for i := range path {
			nextLeft[i] = leftStart[i] - path[i]
			nextRight[i] = rightStart[i] + path[i]
		}
		rightValid := isInBounds(nextRight, matrix)
		leftValid := isInBounds(nextLeft, matrix)

		if rightValid {
			result[arrayToString(nextRight, "")] = true
		}
		if leftValid {
			result[arrayToString(nextLeft, "")] = true
		}

		if !rightValid && !leftValid {
			return
		}

		recover(path, antinodes, matrix, nextLeft, nextRight)
	}

	recover(diff, antinodes, matrix, leftAntenna, leftAntenna)

	return antinodes
}

func isInBounds(slice []int, matrix [][]string) bool {
	if slice[0] >= 0 && slice[0] < len(matrix) && slice[1] <= 0 && slice[1] > -len(matrix) {
		return true
	}
	return false
}

func genPermutations[T any](slice []T, length int) [][]T {
	// Validate input
	if length < 0 || length > len(slice) {
		return [][]T{}
	}

	var result [][]T

	var generate func(current []T, remaining []T)
	generate = func(current []T, remaining []T) {
		// If we've reached the desired length, add to results
		if len(current) == length {
			permutation := make([]T, len(current))
			copy(permutation, current)
			result = append(result, permutation)
			return
		}

		if len(remaining) == 0 {
			return
		}

		// Recursive permutation generation
		for i, elem := range remaining {
			// Add current element to permutation
			newCurrent := append(current, elem)

			// Create new remaining slice without the current element
			newRemaining := make([]T, 0, len(remaining)-1)
			newRemaining = append(newRemaining, remaining[:i]...)
			newRemaining = append(newRemaining, remaining[i+1:]...)

			// Recursive call
			generate(newCurrent, newRemaining)
		}
	}

	// Start the recursive generation
	generate([]T{}, slice)

	return result
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func getDiff(rightAntenna []int, leftAntenna []int) []int {
	diff := make([]int, 2)
	for i := range rightAntenna {
		diff[i] = rightAntenna[i] - leftAntenna[i]
	}
	return diff
}

type antenna struct {
	locations [][]int
}

func formatInput(input string) ([][]string, map[string]antenna) {
	rawSplit := strings.Split(input, "\n")
	matrix := make([][]string, len(rawSplit)-1)
	unique := make(map[string]antenna)

	for i := range rawSplit {
		if len(rawSplit[i]) == 0 {
			break
		}
		line := strings.Split(rawSplit[i], "")
		matrix[i] = line
		for x := range line {
			if matrix[i][x] == "." {
				continue
			}
			loc := make([]int, 2)
			loc[0] = x
			loc[1] = -i
			updatedLocations := append(unique[line[x]].locations, loc)
			unique[line[x]] = antenna{locations: updatedLocations}
		}
	}
	return matrix, unique
}
