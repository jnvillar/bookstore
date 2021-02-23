package validator

import (
	"errors"
	"reflect"
)

type Validator func(interface{}) bool

type SimpleValidation struct {
	Parameter interface{}
	Validator Validator
	ErrorMsg  string
}

func (v SimpleValidation) Err() error {
	return errors.New(v.ErrorMsg)
}

type Func struct {
	Fn func(interface{}) bool
}

func (f *Func) Satisfy(that interface{}) bool {
	return f.Fn(that)
}

type InSet struct {
	ThingInSet []EqualityComparison
}

func (i *InSet) InSet(that interface{}) bool {
	for _, e := range i.ThingInSet {
		if e.Equal(that) {
			return true
		}
	}
	return false
}

type EqualityComparison struct {
	CompareTo interface{}
}

func (ec *EqualityComparison) NonEqual(that interface{}) bool {
	return !ec.Equal(that)
}

func (ec *EqualityComparison) DeepNonEqual(that interface{}) bool {
	return !ec.DeepEqual(that)
}

func (ec *EqualityComparison) Equal(that interface{}) bool {
	return that == ec.CompareTo
}

func (ec *EqualityComparison) DeepEqual(that interface{}) bool {
	return reflect.DeepEqual(that, ec.CompareTo)
}

func NonNil(param interface{}) bool {
	if param == nil {
		return false
	}

	switch reflect.TypeOf(param).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return !reflect.ValueOf(param).IsNil()
	}

	return true
}

func StringPresent(param interface{}) bool {
	return len(param.(string)) > 0
}

func PositiveInt(param interface{}) bool {
	return param.(int64) > 0
}

func ValidateParameter(param interface{}, validator func(interface{}) bool) bool {
	return validator(param)
}

func FirstNonValid(validations ...SimpleValidation) error {
	for _, v := range validations {
		if !ValidateParameter(v.Parameter, v.Validator) {
			return v.Err()
		}
	}
	return nil
}

func FirstNonValidSameValidation(singleValidator Validator, validations ...SimpleValidation) error {
	for _, v := range validations {
		if !ValidateParameter(v.Parameter, singleValidator) {
			return v.Err()
		}
	}
	return nil
}
