package string_sum

import (
	"errors"
	"fmt"
	"regexp"
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
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", fmt.Errorf("e1: %w", errorEmptyInput)
	}

	lessPatern := `^[\s\+-]{0,}\d{1,}$`
	re := regexp.MustCompile(lessPatern)
	less := re.FindAllString(input, 1)
	if len(less) == 1 {
		return "", fmt.Errorf("e2: %w", errorNotTwoOperands)
	}

	pattern := `^[\s\+-]{0,}[0-9a-z]{1,}[\s\+-]{0,}[0-9a-z]{1,}`
	re = regexp.MustCompile(pattern)
	remain := re.ReplaceAllString(input, "")
	if len(remain) > 0 {
		return "", fmt.Errorf("e3: %w", errorNotTwoOperands)
	}

	var x, y int
	var started, signMinus, definedX bool

	for _, v := range input {
		val := string(v)
		if !definedX {
			_, err := strconv.Atoi(val)
			if err != nil {
				if val == "-" {
					signMinus = true
				} else if val == "+" || val == " " {
				} else {
					return "", fmt.Errorf("e3: %w",
						&strconv.NumError{
							Func: "Atoi",
							Num:  strconv.Itoa(x) + val,
							Err:  strconv.ErrSyntax,
						})
				}
				if started {
					definedX = true
				}
			} else {
				x, _ = strconv.Atoi(strconv.Itoa(x) + val)
				if signMinus {
					x = -x
				} else {
					x = x
				}
				signMinus = false
				started = true
			}
		} else if definedX {
			_, err := strconv.Atoi(val)
			if err != nil {
				if val == "-" {
					signMinus = true
				} else if val == "+" || val == " " {
				} else {
					return "", fmt.Errorf("e3: %w",
						&strconv.NumError{
							Func: "Atoi",
							Num:  strconv.Itoa(y) + val,
							Err:  strconv.ErrSyntax,
						})
				}
			} else {
				y, _ = strconv.Atoi(strconv.Itoa(y) + val)
				if signMinus {
					y = -y
				} else {
					y = y
				}
				signMinus = false

			}
		}
	}

	return strconv.Itoa(x + y), nil
}
