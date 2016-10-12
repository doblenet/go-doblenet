package eflags

import (
	"errors"
)

type stringChoiceValue struct {
	options	[]string
	value	*string
}


//////////////////////////////////////////////////////////////////////////////
// implement the Value interface

func (scv *stringChoiceValue) Type() string {
	return "stringChoice"
}

func (scv *stringChoiceValue) Value() string {
	return *scv.value
}

func (scv *stringChoiceValue) String() string {
	return *scv.value
}

func (scv *stringChoiceValue) Set(val string) error {

	if nil == scv.value {
		return errors.New("StringChoiceValue::Set(): nil value")
	}

	for _,c := range scv.options {
		if val == c {
			*scv.value = val
		}
	}	
	return nil
}


//////////////////////////////////////////////////////////////////////////////
func NewStringChoiceValue(valp *string, defval string, choices []string) *stringChoiceValue {
	if nil == valp {
		var v string
		return &stringChoiceValue{value: &v}
	}
	scv := &stringChoiceValue{value: valp, options:choices}
	*scv.value=defval	// !
	return scv
}
