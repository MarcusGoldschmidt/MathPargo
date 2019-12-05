package tests

import (
	"gorser"
	"gorser/utils/errors"
	"testing"
)

func TestShoutThrowErrorCode55(t *testing.T) {
	express , _ := gorser.NewExpression("")

	_, err := express.Calculate(6)

	if err != nil && err.Error() == errors.GenerateNoExpressionEnteredError().Error() {
		return
	}
	t.Error("Calculate nothing from nothing: ", err) // to indicate test failed
}
