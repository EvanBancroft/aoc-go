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
	if part2 {
		part2Files := genFilesChunks(input)
		formatted := formatChunks(part2Files)
		total := calcCheckSum(formatted)
		return total
	}

	files := genFiles(input)
	formatted := formatDisk(files)
	total := calcCheckSum(formatted)
	return total
}

func calcCheckSum(formattedDisk []string) int {
	var total int
	for i, file := range formattedDisk {
		// continue loop once we hit empty memory
		if file == "." {
			continue
		}
		idString := string(file)
		fileId, _ := strconv.Atoi(idString)
		total = (fileId * i) + total
	}
	return total
}

func formatChunks(files [][]string) []string {
	copyFiles := make([][]string, len(files))
	copy(copyFiles, files)

	lastFileId := len(files)
	var sortDisk func([][]string) [][]string
	sortDisk = func(r [][]string) [][]string {
		var emptyIndex int
		var lastChunkIndex int

		for i := len(r) - 1; i >= 0; i-- {
			currId, _ := strconv.Atoi(r[i][0])
			if r[i][0] != "." && currId < lastFileId {
				lastChunkIndex = i
				lastFileId = currId
				break
			}
		}
		for i := range r {
			if i > lastChunkIndex {
				emptyIndex = -1
				break
			}
			if r[i][0] == "." && len(r[i]) >= len(r[lastChunkIndex]) {
				emptyIndex = i
				break
			}
		}

		if lastChunkIndex == 0 { // end loop once we hit the end, could probably be more greedy about this
			return r
		}

		if emptyIndex > 0 {
			diff := len(r[emptyIndex]) - len(r[lastChunkIndex])
			if diff > 0 {
				remainder := r[emptyIndex][:diff]
				r[emptyIndex] = r[emptyIndex][diff:]
				r[emptyIndex], r[lastChunkIndex] = r[lastChunkIndex], r[emptyIndex]
				r = slices.Insert(r, emptyIndex+1, remainder)
			}

			if diff == 0 {
				r[emptyIndex], r[lastChunkIndex] = r[lastChunkIndex], r[emptyIndex]
			}
		}
		return sortDisk(r)
	}

	copyFiles = sortDisk(copyFiles)

	var result []string

	for _, subSlice := range copyFiles {
		result = append(result, subSlice...)
	}

	return result
}

func formatDisk(files []string) []string {
	copyFiles := make([]string, len(files))
	copy(copyFiles, files)

	var lastFileIndex int
	var sortDisk func([]string) []string

	sortDisk = func(r []string) []string {
		emptyIndex := slices.Index(r, ".") // find first empty item in our slice
		reversedSlice := reverse(r)

		for i := range reversedSlice {
			if reversedSlice[i] != "." {
				lastFileIndex = int((i - len(r) + 1) * -1)
				break
			}
		}

		// exit the loop if our next empty pointer is > our first file
		if emptyIndex > lastFileIndex {
			return r
		}

		r[emptyIndex], r[lastFileIndex] = r[lastFileIndex], r[emptyIndex]
		return sortDisk(r)
	}

	copyFiles = sortDisk(copyFiles)
	return copyFiles
}

func genFilesChunks(input string) [][]string {
	output := make([][]string, 0)
	for i, char := range strings.Split(input, "") {
		amount, _ := strconv.Atoi(char)
		newChunk := make([]string, amount)

		if i%2 == 0 {
			currentId := strconv.Itoa(i / 2)

			for j := 0; j < amount; j++ {
				newChunk[j] = currentId
			}
			output = append(output, newChunk)
			continue
		}

		for j := 0; j < amount; j++ {
			newChunk[j] = "."
		}

		if len(newChunk) == 0 {
			continue
		}
		output = append(output, newChunk)
	}
	return output
}

func genFiles(input string) []string {
	output := make([]string, 0)
	for i, char := range strings.Split(input, "") {
		amount, _ := strconv.Atoi(char)

		if i%2 == 0 {
			currentId := strconv.Itoa(i / 2)

			for i := 0; i < amount; i++ {
				output = append(output, currentId)
			}
			continue
		}

		for i := 0; i < amount; i++ {
			output = append(output, ".")
		}
	}

	return output
}

func reverse(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
