// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	terraform "github.com/runatlantis/atlantis/server/events/terraform"
)

func AnySliceOfTerraformTfOutputFile() []terraform.TfOutputFile {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]terraform.TfOutputFile))(nil)).Elem()))
	var nullValue []terraform.TfOutputFile
	return nullValue
}

func EqSliceOfTerraformTfOutputFile(value []terraform.TfOutputFile) []terraform.TfOutputFile {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []terraform.TfOutputFile
	return nullValue
}
