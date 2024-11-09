package validator

import (
	"fmt"
	"reflect"

	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
)

func ValidatorFormGoods(form collection.FormGoods) error {
	val := reflect.ValueOf(form)
	typ := reflect.TypeOf(form)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		// if fieldName == "PhotoUrl" {
		// 	continue
		// }
		if field.IsZero() {
			return fmt.Errorf("field %s is empty", fieldName)
		}
	}
	return nil
}
func ValidatorFormCategory(form collection.FormCategory) error {
	val := reflect.ValueOf(form)
	typ := reflect.TypeOf(form)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		if field.IsZero() {
			return fmt.Errorf("field %s is empty", fieldName)
		}
	}
	return nil
}
