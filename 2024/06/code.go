package main

import (
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
	currentLocation, maze := initializeMaze(input)
	prevLocations := make(map[string]int)
	_, _ = nextMoveLoops(maze, currentLocation, "up", prevLocations)

	uniqueSpots := make(map[string]bool)

	for key := range prevLocations {
		posKey := strings.Split(key, "|")[0]
		uniqueSpots[posKey] = true
	}

	if part2 {
		currentLocation2, maze2 := initializeMaze(input)
		var count int

		for i := range uniqueSpots {

			coords := strings.Split(i, ",")
			x, _ := strconv.Atoi(coords[1])
			y, _ := strconv.Atoi(coords[0])

			maze2[x][y] = "#"

			startingLocation := make([]int, 2)
			copy(startingLocation, currentLocation2)

			prevLocations := make(map[string]int)

			_, loop := nextMoveLoops(maze2, startingLocation, "up", prevLocations)

			if loop {
				count++
			}

			maze2[x][y] = "."

		}

		return count
	}

	return len(uniqueSpots)
}

// Function to detect if the algorithm is stuck in a loop
func isInLoop(visits map[string]int, x, y int, dir string) (bool, map[string]int) {
	threshold := 1
	key := positionKey(x, y, dir)
	visits[key]++
	return visits[key] > threshold, visits
}

func nextMoveLoops(maze [][]string, currentLocation []int, dir string, visitedLocations map[string]int) (map[string]int, bool) {
	if len(visitedLocations) > len(maze)*len(maze[0]) {
		return visitedLocations, false
	}

	loop := false
	loop, visitedLocations = isInLoop(visitedLocations, currentLocation[1], currentLocation[0], dir)

	if loop {
		return visitedLocations, true
	}

	switch dir {
	case "up":
		if currentLocation[0] == 0 {
			break
		}
		nextLocation := maze[currentLocation[0]-1][currentLocation[1]]

		if nextLocation == "." {
			currentLocation[0]--
			return nextMoveLoops(maze, currentLocation, "up", visitedLocations)
		}
		if nextLocation == "#" {
			if currentLocation[1]+1 == len(maze[0]) {
				break
			}
			// check corner case | if right is also blocked
			if maze[currentLocation[0]][currentLocation[1]+1] == "#" {
				currentLocation[0]++
				return nextMoveLoops(maze, currentLocation, "down", visitedLocations)
			}
			currentLocation[1]++
			return nextMoveLoops(maze, currentLocation, "right", visitedLocations)
		}

	case "down":
		if currentLocation[0]+1 == len(maze) {
			break
		}
		nextLocation := maze[currentLocation[0]+1][currentLocation[1]]
		if nextLocation == "." {
			currentLocation[0]++
			return nextMoveLoops(maze, currentLocation, "down", visitedLocations)
		}
		if nextLocation == "#" {
			if currentLocation[1] == 0 {
				break
			}
			// check corner case | if left is also blocked
			if maze[currentLocation[0]][currentLocation[1]-1] == "#" {
				currentLocation[0]--
				return nextMoveLoops(maze, currentLocation, "up", visitedLocations)
			}
			currentLocation[1]--
			return nextMoveLoops(maze, currentLocation, "left", visitedLocations)
		}

	case "left":
		if currentLocation[1] == 0 {
			break
		}
		nextLocation := maze[currentLocation[0]][currentLocation[1]-1]
		if nextLocation == "." {
			currentLocation[1]--
			return nextMoveLoops(maze, currentLocation, "left", visitedLocations)
		}
		if nextLocation == "#" {
			if currentLocation[0] == 0 {
				break
			}
			// check corner case | if up is also blocked
			if maze[currentLocation[0]-1][currentLocation[1]] == "#" {
				currentLocation[1]++
				return nextMoveLoops(maze, currentLocation, "right", visitedLocations)
			}
			currentLocation[0]--
			return nextMoveLoops(maze, currentLocation, "up", visitedLocations)
		}

	case "right":
		if currentLocation[1]+1 == len(maze[0]) {
			break
		}
		nextLocation := maze[currentLocation[0]][currentLocation[1]+1]

		if nextLocation == "." {
			currentLocation[1]++
			return nextMoveLoops(maze, currentLocation, "right", visitedLocations)
		}
		if nextLocation == "#" {
			if currentLocation[0]+1 == len(maze) {
				break
			}
			// check corner case | if down is also blocked
			if maze[currentLocation[0]+1][currentLocation[1]] == "#" {
				currentLocation[1]--
				return nextMoveLoops(maze, currentLocation, "left", visitedLocations)
			}
			currentLocation[0]++
			return nextMoveLoops(maze, currentLocation, "down", visitedLocations)
		}
	}

	return visitedLocations, false
}

func appendPos(visitedLocations map[string]int, currentLocation []int, dir string) map[string]int {
	pos := positionKey(currentLocation[1], currentLocation[0], dir)

	visitedLocations[pos]++

	return visitedLocations
}

// Helper to generate a unique key for a position
func positionKey(x, y int, dir string) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y) + "|" + dir
}

func initializeMaze(input string) ([]int, [][]string) {
	currentLocation := make([]int, 2)
	rows := strings.Split(input, "\n")
	maze := make([][]string, len(rows)-1)

	for i := range rows {
		// trim last empty row
		if len(rows[i]) == 0 {
			continue
		}
		maze[i] = strings.Split(rows[i], "")
	}
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == "^" {

				currentLocation[0] = y
				currentLocation[1] = x

				maze[y][x] = "."

				break
			}
		}
	}

	return currentLocation, maze
}
