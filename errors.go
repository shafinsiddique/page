package main

import (
	"github.com/pkg/errors"
	"reflect"
)

var UNCLOSED_PARENTHESIS = errors.New("Syntax Error : unclosed parenthesis")

var TOO_FEW_ARGUMENTS = errors.New("Parsing Error : too few arguments")

var TOO_MANY_ARGUMENTS = errors.New("Parsing Error : too many arguments")

var UNCLOSED_STRING = errors.New("Syntax Error : unclosed string")

var FIRST_ON_EMPTY = errors.New("Runtime Error : called first on empty list")

var UNEXPECTED_TOKEN = errors.New("Parsing Error : unexpected token")

var FUNCTION_ARGS_ERROR = errors.New("Parsing Error : expected function arguments list.")

var FUNCTION_NAME_REQUIRED = errors.New("Syntax Error : function name required in define.")

var ONLY_SYMBOLS_ALLOWED = errors.New("Parsing Error :  only symbols allowed")

var FUNCTION_NO_BODY = errors.New("Syntax Error : function body not found.")
func TypeMismatchErrorNumber(actual interface{}) error {
	valType := reflect.TypeOf(actual).String()
	return errors.New("Type Error : Expected number got " + valType + " instead")
}

func TypeMismatchError(expected string, actual interface{}) error {
	valType := reflect.TypeOf(actual).String()
	return errors.New("Type Error : Expected " + expected + " got " + valType + " instead")
}

func TypeMismatcErrorList(actual ExpressionType) error {
	valType := reflect.TypeOf(actual).String()
	return errors.New("Type Error : Expected list got " + valType + " instead")

}