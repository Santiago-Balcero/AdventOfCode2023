package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("Advent of Code 2023 - 01")
	file, err := os.Open("input.txt")
	if err != nil {
		panic("File not found")
	}

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstNum := ""
		lastNum := ""
		for _, r := range line {
			if unicode.IsDigit(r) {
				if firstNum == "" {
					firstNum = string(r)
				}
				lastNum = string(r)
			}
		}
		lineValue, _ := strconv.ParseInt(firstNum+lastNum, 10, 64)
		result += int(lineValue)

	}
	file.Close()
	fmt.Println(result)
}
