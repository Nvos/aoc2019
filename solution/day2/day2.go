package day2

import (
	"github.com/sirupsen/logrus"
	"math/rand"
)
func Solve(values []int, noun int, verb int) []int {
	logger := logrus.New()
	index := 0

	values[1] = noun
	values[2] = verb

	var buff []int

	for index < len(values) {
		if index + 4 > len(values) {
			break
		}

		buff = []int{
			values[index],
			values[index+1],
			values[index+2],
			values[index+3],
		}

		if buff[0] == 99 {
			break
		}

		result := 0
		switch buff[0] {
			case 1:
				result = values[buff[1]] + values[buff[2]]
				break
			case 2:
				result = values[buff[1]] * values[buff[2]]
				break
		}

		values[buff[3]] = result
		index += 4
	}

	logger.Info("Final result is: ", values[0])
	return values
}

func BruteForce(values []int) int {
	logger := logrus.New()
	var result int
	noun := 2
	verb := 12

	for result != 19690720 {
		noun = rand.Intn(100)
		verb = rand.Intn(100)

		input := make([]int, len(values))
		copy(input, values)

		result = Solve(input, noun, verb)[0]
	}

	logger.Info("noun =", noun, ", verb = ", verb, ", result = ", 100 * noun + verb)

	return 100 * noun + verb
}