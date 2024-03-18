package lab2

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// CalculatePostfix variant 1: postfix format expression calculation
func CalculatePostfix(input string) (string, error) {
	var err error
	var result float64

	expressionArr := strings.Split(input, " ")
	result, err = strconv.ParseFloat(expressionArr[0], 64)
	if err != nil {
		err = errors.New("operand 1 must be a number")
		return "", err
	}

	for i := 1; i < len(expressionArr); i += 2 {
		operand2, err := strconv.ParseFloat(expressionArr[i], 64)
		if err != nil {
			err = errors.New("operand 2 must be a number")
			return "", err
		}

		switch expressionArr[i+1] {
		case "+":
			result += operand2
		case "-":
			result -= operand2
		case "*":
			result *= operand2
		case "/":
			if operand2 == 0 {
				err = errors.New("division by zero")
				return "", err
			}
			result /= operand2
		case "^":
			result = math.Pow(result, operand2)
		default:
			err = errors.New("invalid operator")
			return "", err
		}
	}

	return fmt.Sprintf("%.2f", result), err
}
