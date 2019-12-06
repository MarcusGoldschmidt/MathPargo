package lexicon

// AnalyzerError a error from Text
type lexiconError struct {
	err    string
	collum int
}

func (e *lexiconError) Error() string {
	return e.err + " on position " + string(e.collum)
}

// GenerateInvalidSyntaxParenthesis when has no expresion valid or found
func GenerateInvalidSyntaxParenthesis(collum int) error {
	return &lexiconError{
		err:    "Parenthesis declaration failed",
		collum: collum,
	}
}
