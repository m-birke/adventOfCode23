package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, _ := os.Open("3_input.txt")
	scanner := bufio.NewScanner(file)

	topLine := ""
	midLine := ""
	botLine := ""

	scanner.Scan()
	midLine = scanner.Text()
	scanner.Scan()
	botLine = scanner.Text()

	var starNeighborhoods [][]bool

	sum := 0

	for scanner.Scan() {
		topLine = midLine
		midLine = botLine
		botLine = scanner.Text()

		starNeighborhoods = ExtractStarNeighborhoods(midLine)

		for _, starNeighborhood := range starNeighborhoods {

			var allStarNeighbors []int

			allStarNeighbors = append(allStarNeighbors, ExtractNeighbors(topLine, starNeighborhood)...)
			allStarNeighbors = append(allStarNeighbors, ExtractNeighbors(midLine, starNeighborhood)...)
			allStarNeighbors = append(allStarNeighbors, ExtractNeighbors(botLine, starNeighborhood)...)

			if len(allStarNeighbors) == 2 {
				sum += (allStarNeighbors[0] * allStarNeighbors[1])
			}
		}

	}
	fmt.Println(sum)
}

func ExtractStarNeighborhoods(line string) [][]bool {

	symbolPattern := regexp.MustCompile(`[*]`)
	symbolMatches := symbolPattern.FindAllStringIndex(line, -1)

	var neighborhoods [][]bool

	for _, symbolMatch := range symbolMatches {
		neighborhood := make([]bool, len(line))

		neighborhood[symbolMatch[0]] = true

		if symbolMatch[0] > 0 {
			neighborhood[symbolMatch[0]-1] = true
		}

		if symbolMatch[0] < (len(line) - 1) {
			neighborhood[symbolMatch[0]+1] = true
		}

		neighborhoods = append(neighborhoods, neighborhood)
	}

	return neighborhoods
}

func ExtractNeighbors(targetLine string, neighborhood []bool) []int {
	pattern := regexp.MustCompile(`[0-9]+`)
	matches := pattern.FindAllStringIndex(targetLine, -1)

	var neighbors []int

	for _, match := range matches {
		for i := match[0]; i < match[1]; i++ {
			if neighborhood[i] {
				tmp := targetLine[match[0]:match[1]]
				number, _ := strconv.Atoi(tmp)
				neighbors = append(neighbors, number)
				break
			}
		}
	}
	return neighbors
}
