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

func StringSum(input string) (output string, err error) {
	operand, sumOperand := FindOperand(input)
	if operand == "" && sumOperand == 1 {
		err = fmt.Errorf("%w", errorEmptyInput)
		return "", err
	}
	if operand == "" && sumOperand > 1 {
		err = fmt.Errorf("%w", errorNotTwoOperands)
		return "", err
	}
	nums := strings.Split(input, operand)
	if len(nums) != 2 {
		err = fmt.Errorf("%w", errorEmptyInput)
		return "", err
	}

	num1, err := strconv.Atoi(nums[0])
	if err != nil {
		err = fmt.Errorf("%w", err)
		return "", err
	}
	num2, err := strconv.Atoi(nums[1])
	if err != nil {
		err = fmt.Errorf("%w", err)
		return "", err
	}

	result := Calculation(num1, num2, operand)
	output = strconv.Itoa(result)
	return output, nil
}

func Calculation(num1, num2 int, operand string) (result int) {
	switch operand {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "%":
		result = num1 % num2
	}
	return result
}

func FindOperand(input string) (operand string, count int) {
	for _, value := range input {
		if value == '+' {
			operand = "+"
			count++
		} else if value == '-' {
			operand = "-"
			count++
		} else if value == '*' {
			operand = "*"
			count++
		} else if value == '/' {
			operand = "/"
			count++
		} else if value == '%' {
			operand = "%"
			count++
		}
	}
	if count > 1 {
		operand = ""
		return operand, count
	}
	return operand, count
}
