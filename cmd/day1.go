package main

import (
	"aoc2019/solution/day1"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		panic("failed to open input")
	}
	defer file.Close()

	var lines []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		float, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("failed to parse text=" + text)
		}

		lines = append(lines, float)
	}

	fuel := day1.TotalFuel(lines)
	fuelRequirement := day1.TotalFuelRequirements(lines)
	println("mass result = " + strconv.FormatInt(int64(fuel), 10))
	println("fuel result = " + strconv.FormatInt(int64(fuelRequirement), 10))
}