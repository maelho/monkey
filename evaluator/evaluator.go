// Package evaluator
package evaluator

import (
	"maelho.github.io/monkey/ast"
	"maelho.github.io/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Stataments
	case *ast.Program:
		return evalStataments(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	// Bollean
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}

	return nil
}

func evalStataments(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
