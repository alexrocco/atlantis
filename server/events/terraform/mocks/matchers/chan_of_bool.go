// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	
)

func AnyChanOfBool() chan bool {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(chan bool))(nil)).Elem()))
	var nullValue chan bool
	return nullValue
}

func EqChanOfBool(value chan bool) chan bool {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue chan bool
	return nullValue
}
