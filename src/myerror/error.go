package myerror

import (
	"errors"
	"fmt"
)

func New(text string) error {
	return errors.New(text)
}

func Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
