package api

import (
	"fmt"
	"reflect"
)

type QueryParamRequiredError struct {
	msg string
}

func (v QueryParamRequiredError) Error() string {
	return v.msg
}

type InvalidIDError struct {
	Type     interface{}
	TypeName string
	Id       string
}

func (i InvalidIDError) Error() string {
	if i.TypeName == "" {
		i.TypeName = reflect.TypeOf(i.Type).Name()
	}
	return fmt.Sprintf("%s is invalid ID for %s", i.Id, i.TypeName)
}
