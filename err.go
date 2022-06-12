package flower

import (
	"fmt"
)

func errNotEnoughArgs(line int, name string, expected int, curretVal int) error {
	return newErrLineName(line, name, fmt.Sprintf("Not enough arguments %d, need to be at least %d", curretVal, expected))
}

func errWrongType(line int, name string, currentType any, expectedType any) error {
	return newErrLineName(line, name, fmt.Sprintf("Wrong type %T. Expected %T", currentType, expectedType))
}

func newErrLineName(line int, name string, message string) error {
	return fmt.Errorf("Line %d >> %s: %s", line, name, message)
}
