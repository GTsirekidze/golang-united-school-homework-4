package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	t := input

	t = strings.TrimSpace(t)

	if len(t) == 0 {
		return "", errorEmptyInput
	}

	fmt.Println(t)
	isFirstNegative := strings.HasPrefix(t, "-")

	arr := strings.Split(t, "-")
	splitNum := len(arr)

	if isFirstNegative {
		splitNum--
	}

	if splitNum == 2 {
		first, errFirst := strconv.Atoi(strings.TrimSpace(arr[len(arr)-2]))

		second, errSecond := strconv.Atoi(strings.TrimSpace(arr[len(arr)-1]))

		if errFirst == nil && errSecond == nil {
			if isFirstNegative {
				first *= -1
			}
			return strconv.Itoa(first - second), nil
		}

	}

	arr = strings.Split(t, "+")
	splitNum = len(arr)

	if splitNum == 2 {
		first, errFirst := strconv.Atoi(strings.TrimSpace(arr[len(arr)-2]))

		second, errSecond := strconv.Atoi(strings.TrimSpace(arr[len(arr)-1]))

		if errFirst == nil && errSecond == nil {
			return strconv.Itoa(first + second), nil
		}
	}

	return "", errorNotTwoOperands
}
