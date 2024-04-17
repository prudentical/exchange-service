package service

import (
	"fmt"
	"reflect"
)

type InvalidStateError struct {
	msg string
}

func (i InvalidStateError) Error() string {
	return i.msg
}

type NotFoundError struct {
	Type interface{}
	Id   int
}

func (i NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID[%d] not found!", reflect.TypeOf(i.Type).Name(), i.Id)
}

type ExchangePairMismatchError struct {
}

func (i ExchangePairMismatchError) Error() string {
	return "Pair with given exchange doesn't exists!"
}

type InvalidPageParameterError struct {
}

func (i InvalidPageParameterError) Error() string {
	return "Page and size should be grater than zero"
}

type OutOfBoundPageSizeError struct {
	Max int
}

func (i OutOfBoundPageSizeError) Error() string {
	return fmt.Sprintf("Page size out of bound [%d-%d]", 1, i.Max)
}

type InvalidPageError struct {
}

func (i InvalidPageError) Error() string {
	return "Invalid page!"
}
