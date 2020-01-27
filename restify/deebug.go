package restify

import (
	"fmt"
	"reflect"
)

func CheckP(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintP(i interface{}) {
	if i != nil {
		panic(fmt.Sprint("Don't know this type", reflect.TypeOf(i)))
	}
}
