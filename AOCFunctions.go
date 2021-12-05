package AOCFunctions

import (
	"io/ioutil"
	"strings"
)

//MinInt returns the mminimum value among a provided list of integers
func MinInt(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

//MaxInt returns the maxiumum value among a provided list of integers
func MaxInt(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

//ReadFile reads the input.txt and parses it into a slice of strings
func ReadFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}
