package main

import (
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
	rows := strings.Split(input, "\n")

	matrix := make([][]byte, len(rows)-1)
	for i := range matrix {
		matrix[i] = make([]byte, len(rows[0]))
	}

	for i := range matrix {
		row := strings.Split(rows[i], "")

		if len(row) == 0 {
			continue
		}

		for j := range matrix[i] {
			matrix[i][j] = []byte(row[j])[0]
		}
	}

	if part2 {
		rowCount := len(matrix)
		columnCount := len(matrix[0])
		var runningCount int

		for i := 1; i < rowCount-1; i++ {
			for j := 1; j < columnCount-1; j++ {
				if matrix[i][j] == 'A' {
					var coords []byte
					coords = append(coords, matrix[i-1][j-1])
					coords = append(coords, matrix[i-1][j+1])
					coords = append(coords, matrix[i+1][j-1])
					coords = append(coords, matrix[i+1][j+1])

					var sCount int
					var mCount int

					for i := range coords {
						if coords[i] == 'S' {
							sCount++
						}
						if coords[i] == 'M' {
							mCount++
						}
					}

					if sCount == 2 && mCount == 2 {
						if !(coords[0] == 'S' && coords[3] == 'S' || coords[1] == 'S' && coords[2] == 'S') {
							runningCount++
						}
					}

				}
			}
		}
		return runningCount
	}

	word := "XMAS"

	columnTotal := searchColumns(matrix, word)
	rowTotal := searchRows(matrix, word)
	leftDiagonalTotal := diagSearch(matrix, word)

	// solve part 1 here
	return columnTotal + rowTotal + leftDiagonalTotal
}

func searchColumns(matrix [][]byte, word string) int {
	runningTotal := 0
	reverseWord := reverseString(word)

	for i := 0; i < len(matrix[0]); i++ {
		var column string
		for _, row := range matrix {
			column += string(row[i])
		}

		runningTotal += wordCountHelper(string(column), word)
		runningTotal += wordCountHelper(string(column), reverseWord)
	}
	return runningTotal
}

func searchRows(matrix [][]byte, word string) int {
	runningTotal := 0
	reverseWord := reverseString(word)

	for i := range matrix {
		row := string(matrix[i])

		runningTotal += wordCountHelper(row, word)
		runningTotal += wordCountHelper(row, reverseWord)

	}
	return runningTotal
}

func diagSearch(letters [][]byte, word string) int {
	reverseWord := reverseString(word)
	count := 0
	rows := len(letters)
	cols := len(letters[0])

	// Top Left
	for start := 0; start < rows+cols-1; start++ {
		var line string
		for x, y := max(0, start-cols+1), max(0, cols-1-start); x < rows && y < cols; x, y = x+1, y+1 {
			line += string(letters[x][y])
		}
		count += wordCountHelper(line, word)
		count += wordCountHelper(line, reverseWord)
	}

	// Top Right
	for start := 0; start < rows+cols-1; start++ {
		var line string
		for x, y := max(0, start-cols+1), min(cols-1, start); x < rows && y >= 0; x, y = x+1, y-1 {
			line += string(letters[x][y])
		}
		count += wordCountHelper(line, word)
		count += wordCountHelper(line, reverseWord)
	}

	return count
}

// Utility max function for integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Utility min function for integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func wordCountHelper(line, word string) int {
	count := 0
	for i := 0; i <= len(line)-len(word); i++ {
		if line[i:i+len(word)] == word {
			count++
		}
	}
	return count
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
