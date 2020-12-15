package main

import (
	"github.com/pkg/errors"
	"reflect"
)

var UNCLOSED_PARENTHESIS = errors.New("Syntax Error : Unclosed Parenthesis")

var TOO_FEW_ARGUMENTS = errors.New("too few arguments")

func TypeMismatchErrorNumber(actual interface{}) error {
	valType := reflect.TypeOf(actual).String()
	return errors.New("Type Error : Expected number got " + valType + " instead")

}