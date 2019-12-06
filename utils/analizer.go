package utils

import (
	"mathPargo/utils/errors"
	"regexp"
)

// The Value of a a token can be
const (
	BinaryExpression OperationType = iota
	Value
	Variable
	Function
)

// OperationType type of a operation can be, the result from a lexicon or a content of node
type OperationType int8

// SeparateSymbolsResult is the result from a analyses of a math expression
type SeparateSymbolsResult struct {
	leftOperator  *string
	operation     *string
	rightOperator *string
	operationType OperationType
}

// SeparateSymbols in a binary operation on a function
// Return left, operator, right and a OperationType
func SeparateSymbols(value *string) (*SeparateSymbolsResult, error) {

	if *value == "" {
		return nil, errors.GenerateNoExpressionEnteredError()
	}

	return &SeparateSymbolsResult{}, nil
}

// REGEX for ^(-?\(?\d|\.\)?)*([\*|\/|\^]{1})?(-?\d|\.)*

func IndentifyBlock(value *string) (int, OperationType, error) {
	regexFunctions, _ := regexp.Compile(`^\w{3}\(.*\)`)
	regexValue, _ := regexp.Compile(`^-?(\d*.\d*|\w)`)
	regexParenthesesBlock, _ := regexp.Compile(`^\(?.*\)?`)

	valueBytes := []byte(*value)

	if block := regexValue.Find(valueBytes); block != nil {
		return len(block), Value, nil
	}

	if block := regexFunctions.Find(valueBytes); block != nil {
		return len(block), Function, nil
	}

	if block := regexParenthesesBlock.Find(valueBytes); block != nil {
		return len(block), BinaryExpression, nil
	}

	return 0, 0, errors.GenerateNotValidExpression(*value)
}

func ValidateGeneralExpression(value *string) error  {
	if ValidateParentheses(value) {
		return errors.GenerateInvalidSyntaxParenthesis(*value)
	}
	return nil
}

func ValidateParentheses(value *string) bool {
	numberOf := 0
	for _, el := range *value{
		if el == '(' {
			numberOf++
		}

		if el == ')'{
			numberOf--
		}
	}
	return numberOf == 0
}
