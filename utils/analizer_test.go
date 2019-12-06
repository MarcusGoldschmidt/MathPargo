package utils

import (
	"testing"
)

type SeparateSymbolsResultTest struct {
	input    string
	expected SeparateSymbolsResult
	obtained *SeparateSymbolsResult
}

func (el *SeparateSymbolsResultTest) HasError() bool {
	return el.expected != *el.obtained
}

func TestShoutThrowErrors(t *testing.T) {
	tests := []SeparateSymbolsResultTest{
		{
			input: "x+x",
			expected: SeparateSymbolsResult{
				leftOperator:  "x",
				operation:     "+",
				rightOperator: "x",
				operationType: BinaryExpression,
			},
		},
		{
			input: "2.5",
			expected: SeparateSymbolsResult{
				value:         2.5,
				operationType: Value,
			},
		},
		{
			input: "x",
			expected: SeparateSymbolsResult{
				variable:      "x",
				operationType: Variable,
			},
		},
		{
			input: "(x+x)-x",
			expected: SeparateSymbolsResult{
				leftOperator:  "(x+x)",
				operation:     "-",
				rightOperator: "x",
				operationType: BinaryExpression,
			},
		},
		{
			input: "x+x*x",
			expected: SeparateSymbolsResult{
				leftOperator:  "x",
				operation:     "+",
				rightOperator: "x*x",
				operationType: BinaryExpression,
			},
		},
		{
			input: "x+x^x",
			expected: SeparateSymbolsResult{
				leftOperator:  "x",
				operation:     "+",
				rightOperator: "x^x",
				operationType: BinaryExpression,
			},
		},
		{
			input: "(x+x)^x",
			expected: SeparateSymbolsResult{
				leftOperator:  "(x+x)",
				operation:     "^",
				rightOperator: "x",
				operationType: BinaryExpression,
			},
		},
	}

	for _, el := range tests {
		var err error
		el.obtained, err = SeparateSymbols(&el.input)

		if err != nil {
			t.Error("Error: ", err)
		}

		if el.HasError() {
			t.Error("\nInput: ", el.input, "\nShould: ", el.expected, "\nThrow: ", el.obtained)
		}
	}
}
