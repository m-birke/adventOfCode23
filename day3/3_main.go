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

	midLine := ""
	botLine := ""

	var topNeighborhood []bool
	var midNeighborhood []bool
	var botNeighborhood []bool

	scanner.Scan()
	midLine = scanner.Text()
	scanner.Scan()
	botLine = scanner.Text()

	midNeighborhood = ExtractNeighborhood(midLine)
	botNeighborhood = ExtractNeighborhood(botLine)
	neighborhood := OrOnSlice(midNeighborhood, botNeighborhood)

	sum := ExtractNumbersInNeighborhoodSum(midLine, neighborhood)

	for scanner.Scan() {
		midLine = botLine
		botLine = scanner.Text()

		topNeighborhood = midNeighborhood
		midNeighborhood = botNeighborhood
		botNeighborhood = ExtractNeighborhood(botLine)

		neighborhood = OrOnSlice(botNeighborhood, OrOnSlice(midNeighborhood, topNeighborhood))

		sum += ExtractNumbersInNeighborhoodSum(midLine, neighborhood)
	}

	neighborhood = OrOnSlice(botNeighborhood, midNeighborhood)
	sum += ExtractNumbersInNeighborhoodSum(botLine, neighborhood)

	fmt.Println(sum)
}

func ExtractNeighborhood(line string) []bool {

	symbolPattern := regexp.MustCompile(`[@#$%&*+?/=-]`)
	symbolMatches := symbolPattern.FindAllStringIndex(line, -1)

	neighborhood := make([]bool, len(line))

	for _, symbolMatch := range symbolMatches {
		neighborhood[symbolMatch[0]] = true

		if symbolMatch[0] > 0 {
			neighborhood[symbolMatch[0]-1] = true
		}

		if symbolMatch[0] < (len(line) - 1) {
			neighborhood[symbolMatch[0]+1] = true
		}
	}

	return neighborhood
}

func ExtractNumbersInNeighborhoodSum(targetLine string, neighborhood []bool) int {
	pattern := regexp.MustCompile(`[0-9]+`)
	matches := pattern.FindAllStringIndex(targetLine, -1)

	sum := 0

	for _, match := range matches {
		for i := match[0]; i < match[1]; i++ {
			if neighborhood[i] {
				tmp := targetLine[match[0]:match[1]]
				number, _ := strconv.Atoi(tmp)
				sum += number
				break
			}
		}
	}
	return sum
}

func OrOnSlice(left []bool, right []bool) []bool {
	result := make([]bool, len(left))

	for i := 0; i < len(left); i++ {
		result[i] = left[i] || right[i]
	}

	return result
}
