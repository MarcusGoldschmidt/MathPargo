package mathPargo

import (
	"mathPargo/utils"
)

// MathExpression represent a math expression to parser
type MathExpression struct {
	Expression string
	binaryTree *utils.Node
}

// Calculate calculate x from a string
func (m *MathExpression) Calculate(x float64) (float64, error) {
	var err error

	if m.binaryTree == nil {
		err = m.Generate()
		return -1, err
	}

	return m.binaryTree.Calculate(&x), nil
}

// Generate the structure to run que math expression
func (m *MathExpression) Generate() error {
	var err error

	m.binaryTree, err = utils.NewNode(&m.Expression)

	return err
}

// NewExpression return a math expression that can be calculated
func NewExpression(expression string) (MathExpression, error) {
	mathExpression := MathExpression{
		Expression: expression,
	}
	err := mathExpression.Generate()
	return mathExpression, err
}
