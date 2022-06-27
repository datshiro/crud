package errors

import "fmt"

type CustomError string

func (e CustomError) Error() string {
	return string(e)
}

func NewParamErr(err string) CustomParamError {
	return CustomParamError(err)
}

type CustomParamError string

func (e CustomParamError) Error() string {
	return fmt.Sprintf("%s: %s", InvalidParamError, string(e))
}
