package lexicon

import "regexp"

// The Type of a token can be
const (
	TokenOperator TokenType = iota
	TokenNumber
	TokenVariable
)

type TokenType int8

// Token a valid element on the alfhabeth with contex
type Token struct {
	Value string
	Type  TokenType
}

type Context int8

// All the Context that can be
const (
	ContextFunction Context = iota
	ContextVoid
	ContextBinaryOperator
	ContextReadValue
	ContextReadOperator
)

var normalOperators = []string{
	"+",
	"-",
}

var priorityOperators = []string{
	"*",
	"/",
	"^",
}

var regexAnyNumber, _ = regexp.Compile(`^-?\d*\.\d*`)

// LexicalAnalysis return all valid token find on a string
// Is based os math expressions
func LexicalAnalysis(value string) ([]Token, error) {
	result := []Token{}

	currentContext := ContextVoid

	currentCollun := 0

	for currentCollun < len(value) {

		switch currentContext {
		case ContextVoid:
			if regexAnyNumber.Match([]byte(value[currentCollun:])) {

			}

		}
	}

	return nil, result
}
