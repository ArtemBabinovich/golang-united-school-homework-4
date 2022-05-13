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

var d = ""

func StringSum(input string) (output string, err error) {

	innerWithoutSpace := strings.ReplaceAll(input, " ", "")
	fmt.Println(input)
	if d != input {
		d = input
		fmt.Println(StringSum(input))
	}

	if len(innerWithoutSpace) == 0 {
		err = fmt.Errorf("errorEmptyInput: %w", errorEmptyInput)
		return
	}

	err = nil
	currentToken := ""
	arguments := make([]string, 0)

	for _, v := range innerWithoutSpace {
		if v == 45 { // -
			if len(currentToken) > 0 {
				arguments = append(arguments, currentToken)
				currentToken = string(v)
			} else {
				currentToken += string(v)
			}
		} else if v == 43 { //+
			if len(currentToken) > 0 {
				arguments = append(arguments, currentToken)
				currentToken = ""
			} else {
				currentToken += string(v)
			}
		} else if v >= 48 && v <= 57 {
			currentToken += string(v)
		} else {
			currentToken += string(v)
			_, err2 := strconv.Atoi(currentToken)
			err = fmt.Errorf("bad token: "+currentToken+" %w", err2)
			return
		}
	}

	arguments = append(arguments, currentToken)

	if len(arguments) != 2 {
		err = fmt.Errorf("errorNotTwoOperands: %w", errorNotTwoOperands)
		return
	}

	val1, err := strconv.Atoi(arguments[0])

	if err != nil {
		err = fmt.Errorf("bad first argument")
		return
	}

	val2, err := strconv.Atoi(arguments[1])

	if err != nil {
		err = fmt.Errorf("bad second argument")
		return
	}

	output = fmt.Sprint(val1 + val2)

	return
}
