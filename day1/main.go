package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f_num_maps := make(map[string]string)
	f_num_maps["twone"] = "21"
	f_num_maps["oneight"] = "18"
	f_num_maps["eightwo"] = "82"
	f_num_maps["eighthree"] = "83"
	f_num_maps["threeight"] = "38"
	f_num_maps["fiveight"] = "58"
	f_num_maps["sevenine"] = "79"
	f_num_maps["nineight"] = "98"

	num_maps := make(map[string]string)
	num_maps["one"] = "1"
	num_maps["two"] = "2"
	num_maps["three"] = "3"
	num_maps["four"] = "4"
	num_maps["five"] = "5"
	num_maps["six"] = "6"
	num_maps["seven"] = "7"
	num_maps["eight"] = "8"
	num_maps["nine"] = "9"

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		for word, num := range f_num_maps {
			line = strings.Replace(line, word, num, -1)
		}
		for word, num := range num_maps {
			line = strings.Replace(line, word, num, -1)
		}

		firstDigit := -1
		lastDigit := -1

		for _, char := range string(line) {
			if !(unicode.IsDigit(char)) {
				continue
			}
			if firstDigit < 0 {
				firstDigit = int(char)
				continue
			}
			lastDigit = int(char)
		}

		if lastDigit < 0 {
			lastDigit = firstDigit
		}

		lineNumber, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		sum += lineNumber
	}
	fmt.Println(sum)
	file.Close()
}
