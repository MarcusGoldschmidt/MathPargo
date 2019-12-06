package utils

import (
	"gorser/utils/errors"
	"strconv"
)

const (
	// VariableChar the initial variable that can be change
	VariableChar = "x"
)

// Node of a binary tree
type Node struct {
	value    string
	left     *Node
	right    *Node
	nodeType OperationType
	Error    *errors.AnalyzerError
}

// Calculate a math expression
func (b *Node) Calculate(x *float64) float64 {
	switch b.nodeType {
	case BinaryExpression:
	case Value:
		result, _ := strconv.ParseFloat(b.value, 64)
		return result
	case Variable:
		return *x
	case Function:
	}
	return 1
}

// NewNode create a binary tree from a string
func NewNode(expression *string) (*Node, error) {
	errorChannel := make(chan error)
	nodeChannel := make(chan *Node)

	// Validate general string
	err := ValidateGeneralExpression(expression)
	if err != nil {
		return nil, err
	}

	go newNodeRoutine(expression, errorChannel, nodeChannel)

	return <-nodeChannel, <-errorChannel
}

// newNodeRoutine nodeChannel is the response
func newNodeRoutine(expression *string, errorChannel chan error, nodeChannel chan *Node) {
	var err error
	var node Node

	// Some channel has detected a error, so stop all
	if errorChannel != nil {
		nodeChannel <- nil
	}

	// Currente node symbols error
	// Stop all routines with errorChannel
	result, err := SeparateSymbols(expression)
	if err != nil {
		errorChannel <- err
		nodeChannel <- nil
		close(errorChannel)
		close(nodeChannel)
		return
	}

	// Generate the binary tree
	switch result.operationType {
	case BinaryExpression:
		nodeChannelLeft := make(chan *Node)
		nodeChannelRight := make(chan *Node)

		go newNodeRoutine(result.leftOperator, errorChannel, nodeChannelLeft)
		go newNodeRoutine(result.rightOperator, errorChannel, nodeChannelRight)

		node.value = *result.operation

		node.left, node.right = <-nodeChannelLeft, <-nodeChannelRight
	case Value:
		node.value = *result.leftOperator
	case Variable:
		// TODO: Make it work for any variable that que user inputs
		node.value = VariableChar
	case Function:
		nodeChannelRight := make(chan *Node)

		go newNodeRoutine(result.leftOperator, errorChannel, nodeChannelRight)

		node.value = *result.operation
		node.right = <-nodeChannelRight
	}

	nodeChannel <- &node
}
