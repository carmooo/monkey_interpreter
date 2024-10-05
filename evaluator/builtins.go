package evaluator

import (
	"github.com/carmooo/monkey_interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"puts":  object.GetBuiltinByName("puts"),
	"len":   object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"last":  object.GetBuiltinByName("last"),
	"rest":  object.GetBuiltinByName("rest"),
	"push":  object.GetBuiltinByName("push"),
}
