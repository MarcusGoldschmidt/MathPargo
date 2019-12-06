package utils

import (
	"mathPargo/utils/errors"
	"regexp"
	"strconv"
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
	leftOperator   string
	operation      string
	rightOperator  string
	value          float64
	variable       string
	paramsFunction string
	operationType  OperationType
}

// SeparateSymbols in a binary operation on a function
// Return left, operator, right and a OperationType
func SeparateSymbols(value *string) (*SeparateSymbolsResult, error) {

	if *value == "" {
		return nil, errors.GenerateNoExpressionEnteredError()
	}

	blockSize, operationType, err := IdentifyBlock(value)

	if err != nil {
		return nil, err
	}

	switch operationType {
	case BinaryExpression:
		return &SeparateSymbolsResult{
			leftOperator:  (*value)[0:blockSize],
			operation:     string((*value)[blockSize]),
			rightOperator: (*value)[blockSize+1:],
			operationType: BinaryExpression,
		}, nil
	case Value:
		floatValue, err := strconv.ParseFloat((*value)[0:blockSize], 64)
		if err != nil {
			return nil, err
		}
		return &SeparateSymbolsResult{
			value:         floatValue,
			operationType: Value,
		}, nil
	case Variable:
		return &SeparateSymbolsResult{
			variable:      (*value)[0:blockSize],
			operationType: Variable,
		}, nil
	case Function:
		return &SeparateSymbolsResult{
			operation:      (*value)[0:blockSize],
			paramsFunction: GetFunctionParams(value),
			operationType:  Function,
		}, nil
	}

	return &SeparateSymbolsResult{}, nil
}

// REGEX for ^(-?\(?\d|\.\)?)*([\*|\/|\^]{1})?(-?\d|\.)*

// IdentifyBlock
func IdentifyBlock(value *string) (int, OperationType, error) {
	regexFunctions, _ := regexp.Compile(`^\w{3}\(.*\)`)
	regexValue, _ := regexp.Compile(`^-?(\(?\d*\.\d*\)?)`)
	regexBinaryExpression, _ := regexp.Compile(`^-?(-?\(.*\)|-?\d*\.\d*|\w{3}\(.*\)|\w)`)
	regexVariableOperation, _ := regexp.Compile(`^-?\w*`)

	valueBytes := []byte(*value)

	if block := regexVariableOperation.Find(valueBytes); block != nil {
		return len(string(block)), Variable, nil
	}

	if block := regexValue.Find(valueBytes); block != nil {
		return fixPriorityOperatorSize(block, value, regexBinaryExpression), Value, nil
	}

	if block := regexBinaryExpression.Find(valueBytes); block != nil {
		return fixPriorityOperatorSize(block, value, regexBinaryExpression), BinaryExpression, nil
	}

	if block := regexFunctions.Find(valueBytes); block != nil {
		return len(string(block)), Function, nil
	}

	return 0, 0, errors.GenerateNotValidExpression(*value)
}

func ValidateGeneralExpression(value string) error {
	if !ValidateParentheses(value) {
		return errors.GenerateInvalidSyntaxParenthesis(value)
	}
	return nil
}

func GetFunctionParams(value *string) string {
	start, end := 0, 0
	state := false

	for i, el := range *value {
		if !state && el == '(' {
			state = true
			start = i
		}
		if state && el == ')' {
			end = i
		}
	}

	return (*value)[start:end]
}

func ValidateParentheses(value string) bool {
	numberOf := 0
	for _, el := range value {
		if el == '(' {
			numberOf++
			continue
		}

		if el == ')' {
			numberOf--
			continue
		}
	}
	return numberOf == 0
}

func fixPriorityOperatorSize(block []byte, value *string, regexBinaryExpression *regexp.Regexp) int {
	// Identify Block with operator priority
	fistBlockSize := len(string(block))

	// Find next valid block
	secondBlockSize := 0
	if len(*value) < fistBlockSize-1 && verifyPriorityOperator(string((*value)[fistBlockSize+1])) {
		block = regexBinaryExpression.Find([]byte((*value)[fistBlockSize+1:]))
		// Add one because the operator
		secondBlockSize = len(string(block)) + 1
	}

	// No need to update the order
	if fistBlockSize + secondBlockSize == len(*value){
		return fistBlockSize
	}

	return fistBlockSize + secondBlockSize
}

func verifyPriorityOperator(value string) bool {
	return value == "*" || value == "/" || value == "^"
}
