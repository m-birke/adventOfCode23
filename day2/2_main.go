package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	maxNumColorMap := make(map[string]int)
	maxNumColorMap["red"] = 12
	maxNumColorMap["green"] = 13
	maxNumColorMap["blue"] = 14

	file, _ := os.Open("2_input.txt")
	scanner := bufio.NewScanner(file)

	//possibleGameSum := 0
	gameWightSum := 0

	for scanner.Scan() {
		line := scanner.Text() // 1 game

		noiseRemoved := strings.Split(line, "Game ")[1]
		gameIdResultSplit := strings.Split(noiseRemoved, ": ")
		//gameId, _ := strconv.Atoi(gameIdResultSplit[0])
		experimentResults := gameIdResultSplit[1]

		maxColorCount := make(map[string]int)
		maxColorCount["red"] = 0
		maxColorCount["green"] = 0
		maxColorCount["blue"] = 0

		//isGamePossible := true
		experimentsSplit := strings.Split(experimentResults, "; ")
		for _, experiment := range experimentsSplit {
			//fmt.Println(experiment)
			colorNames, colorCounts := CountColorsOfExperiment(experiment)
			//isExpPossible := CheckExperimentPossible(colorNames, colorCounts, maxNumColorMap)
			//if !(isExpPossible) {
			//	isGamePossible = false
			//	continue
			//}

			for idx, colorName := range colorNames {
				if maxColorCount[colorName] < colorCounts[idx] {
					maxColorCount[colorName] = colorCounts[idx]
				}
			}
		}

		gameWeight := 1
		for _, maxColorCount := range maxColorCount {
			gameWeight *= maxColorCount
		}

		gameWightSum += gameWeight
		//if isGamePossible {
		//	possibleGameSum += gameId
		//}
	}
	//fmt.Println(possibleGameSum)
	fmt.Println(gameWightSum)
}

func CountColorsOfExperiment(drawnColors string) ([]string, []int) {
	colorSplit := strings.Split(drawnColors, ", ")
	numColors := uint8(len(colorSplit))
	colorCounts := make([]int, numColors)
	colorNames := make([]string, numColors)
	for idx, colorStr := range colorSplit {
		colorStrSplit := strings.Split(colorStr, " ")
		colorCountStr := colorStrSplit[0]
		colorName := colorStrSplit[1]
		colorCount, _ := strconv.Atoi(colorCountStr)
		colorNames[idx] = colorName
		colorCounts[idx] = colorCount
	}
	return colorNames, colorCounts
}

func CheckExperimentPossible(colorNames []string, colorCounts []int, maxNumColorMap map[string]int) bool {
	for idx, colorName := range colorNames {
		if colorCounts[idx] > maxNumColorMap[colorName] {
			return false
		}
	}
	return true
}
