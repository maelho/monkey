package evaluator

import (
	"maelho.github.io/monkey/ast"
	"maelho.github.io/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
