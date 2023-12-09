package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToIntArray(ns string) ([]int, error) {
	stringNumbers := strings.Split(ns, ",")
	numbers := make([]int, len(stringNumbers))
	for ind, value := range stringNumbers {
		number, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("not a integer number")
		}
		numbers[ind] = number
	}

	return numbers, nil
}

func StringToFloat64Array(ns string) ([]float64, error) {
	stringNumbers := strings.Split(ns, ",")
	numbers := make([]float64, len(stringNumbers))
	for ind, value := range stringNumbers {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("not a number for an id")
		}
		numbers[ind] = number
	}

	return numbers, nil
}
