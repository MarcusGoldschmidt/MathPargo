package mathPargo

import (
	"mathPargo/utils/errors"
	"testing"
)

type test struct {
	input    string
	expected error
	obtained error
}

func (m *test) HasError() bool {
	return m.expected.Error() != m.obtained.Error()
}

func TestShoutNotThrowErrorOnCreateCreate(t *testing.T) {
	tests := []test{
		{
			input:    "",
			expected: errors.GenerateNoExpressionEnteredError(),
		},
		{
			input:    "(x+2+2)-(",
			expected: errors.GenerateInvalidSyntaxParenthesis("(x+2+2)-("),
		},
		{
			input:    "(x+2+2))",
			expected: errors.GenerateInvalidSyntaxParenthesis("(x+2+2))"),
		},
	}

	for _, el := range tests {
		_, err := NewExpression(el.input)
		el.obtained = err
		if err != nil && el.HasError() {
			t.Error("\nInput: ", el.input, "\nShould: ", el.expected.Error(), "\nThrow: ", err.Error())
		}
	}
}
