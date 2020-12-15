package main

import (
	"github.com/pkg/errors"
	"reflect"
)

var UNCLOSED_PARENTHESIS = errors.New("Syntax Error : unclosed parenthesis")

var TOO_FEW_ARGUMENTS = errors.New("Parsing Error : too few arguments")

var UNCLOSED_STRING = errors.New("Syntax Error : unclosed string")

func TypeMismatchErrorNumber(actual interface{}) error {
	valType := reflect.TypeOf(actual).String()
	return errors.New("Type Error : Expected number got " + valType + " instead")

}