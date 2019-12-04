package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		panic("failed to open input")
	}
	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("failed to parse text=" + text)
		}

		lines = append(lines, value)
	}
}