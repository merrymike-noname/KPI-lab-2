package lab2

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePostfixWith2Operands(t *testing.T) {
	res, err := CalculatePostfix("2 2 *")
	if assert.Nil(t, err) {
		assert.Equal(t, "4.00", res)
	}
}

func TestCalculatePostfixWith3Operands(t *testing.T) {
	res, err := CalculatePostfix("2 2.5 * 4 /")
	if assert.Nil(t, err) {
		assert.Equal(t, "1.25", res)
	}
}

func TestCalculatePostfixWithNegativePower(t *testing.T) {
	res, err := CalculatePostfix("2 -2 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, "0.25", res)
	}
}

func TestCalculatePostfixWith7Operands(t *testing.T) {
	res, err := CalculatePostfix("8 2 / 9 * 32 - 3 ^ 20 + 4 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "80.00", res)
	}
}

func TestCalculatePostfixWith8Operands(t *testing.T) {
	res, err := CalculatePostfix("8 2 / 9 * 32 - 3 ^ 20 + 4 - 10 /")
	if assert.Nil(t, err) {
		assert.Equal(t, "8.00", res)
	}
}

func TestCalculatePostfixWith10Operands(t *testing.T) {
	res, err := CalculatePostfix("8 2 / 9 * 32 - 3 ^ 20 + 4 - 10 / -1 ^ 50 *")
	if assert.Nil(t, err) {
		assert.Equal(t, "6.25", res)
	}
}

func TestCalculatePostfixDivisionByZero(t *testing.T) {
	_, err := CalculatePostfix("2 2 + 0 /")
	expectedError := errors.New("division by zero")
	assert.EqualError(t, err, expectedError.Error())
}

func TestCalculatePostfixIncorrectOperand1(t *testing.T) {
	_, err := CalculatePostfix("a 2 +")
	expectedError := errors.New("operand 1 must be a number")
	assert.EqualError(t, err, expectedError.Error())
}

func TestCalculatePostfixIncorrectOperand2(t *testing.T) {
	_, err := CalculatePostfix("2 b +")
	expectedError := errors.New("operand 2 must be a number")
	assert.EqualError(t, err, expectedError.Error())
}

func TestCalculatePostfixIncorrectOperator(t *testing.T) {
	_, err := CalculatePostfix("2 2 $")
	expectedError := errors.New("invalid operator")
	assert.EqualError(t, err, expectedError.Error())
}

func TestCalculatePostfixEmptyInput(t *testing.T) {
	_, err := CalculatePostfix("")
	expectedError := errors.New("empty input")
	assert.EqualError(t, err, expectedError.Error())
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 + 3 *")
	fmt.Println(res)

	// Output:
	// 12.00
}
