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

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func prepareStr(input string) string {
	inputWhiteSpacesOff := strings.ReplaceAll(input, " ", "")
	if strings.HasPrefix(inputWhiteSpacesOff, "+") {
		inputWhiteSpacesOff = inputWhiteSpacesOff[1:]
	}
	return inputWhiteSpacesOff
}

func getSumOfStrNums(input []string, binarSign bool) (string, error) {
	firstOper, err := strconv.Atoi(input[0])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	secondOper, err := strconv.Atoi(input[1])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	if binarSign {
		return strconv.Itoa(firstOper + secondOper), nil
	}
	return strconv.Itoa(secondOper - firstOper), nil
}

func getSubOfStrNums(input []string, binarSign bool) (string, error) {
	firstOper, err := strconv.Atoi(input[0])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	secondOper, err := strconv.Atoi(input[1])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	if binarSign {
		return strconv.Itoa(firstOper - secondOper), nil
	}
	return strconv.Itoa(-(firstOper + secondOper)), nil
}

func StringSum(input string) (output string, err error) {
	preparedString := prepareStr(input)
	if len(preparedString) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	binSign := true
	if strings.HasPrefix(preparedString, "-") {
		binSign = false
		preparedString = preparedString[1:]
	}

	inputApartPlus := strings.Split(preparedString, "+")
	if len(inputApartPlus) > 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	if len(inputApartPlus) == 2 {
		return getSumOfStrNums(inputApartPlus, binSign)
	}

	inputApartMinus := strings.Split(preparedString, "-")
	if len(inputApartMinus) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	return getSubOfStrNums(inputApartMinus, binSign)
}
