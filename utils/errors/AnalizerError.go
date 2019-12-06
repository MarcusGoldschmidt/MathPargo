package errors

// AnalyzerError a error from Text
type AnalyzerError struct {
	Err string
}

func (e *AnalyzerError) Error() string {
	return e.Err
}

// GenerateNoExpressionEnteredError when has no expresion to be calculated
func GenerateNoExpressionEnteredError() error {
	return &AnalyzerError{
		Err: "No expression entered",
	}
}

// GenerateNotValidExpression when has no expresion valid or found
func GenerateNotValidExpression(value string) error {
	return &AnalyzerError{
		Err: "Not a valid expression: " + value,
	}
}

// GenerateInvalidSyntaxParenthesis when has no expresion valid or found
func GenerateInvalidSyntaxParenthesis(value string) error {
	return &AnalyzerError{
		Err: "Parenthesis declaration failed: " + value,
	}
}
