package tests

import (
	"gorser"
	"gorser/utils/errors"
	"testing"
)

type test struct {
	input    string
	expected error
	obtained error
}

func (m *test) AssertErrorEquals() bool {
	if m.obtained == nil {
		return false
	}
	return m.expected.Error() == m.expected.Error()
}

func TestShoutNotThrowErrorOnCreateCreate(t *testing.T) {
	tests := []test{
		test{
			input:    "",
			expected: errors.GenerateNoExpressionEnteredError(),
		},
		test{
			input:    "",
			expected: errors.GenerateNotValidExpression(""),
		},
		test{
			input:    "(x+2",
			expected: errors.GenerateInvalidSyntaxParenthesis("(x+2"),
		},
	}

	for _, el := range tests{
		_, err := gorser.NewExpression(el.input)
		el.obtained = err
		if !el.AssertErrorEquals() {
			t.Error("Calculate nothing from nothing: ", err)
		}
	}
}
