package mercurie_assessment

import (
	"fmt"

)

type MultiError struct {
	errors []error
}

func NewMultiError() *MultiError {
	return &MultiError{
		errors: make([]error, 0),
	}
}

func (me *MultiError) Append(e error) {
	me.errors = append(me.errors, e)
}

func (me *MultiError) HasErrors() bool {
	return len(me.errors) > 0
}

func (me *MultiError) Error() string {
	return fmt.Sprintf(
		"multiple errors occured: %v",
		me.errors,
	)
}
