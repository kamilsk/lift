package safe

import "github.com/pkg/errors"

// Do reliably runs the action and captures panic as its error.
// If an error is not nil, it passes it to the handler.
//
//  go safe.Do(
//  	func() error { ... },
//  	func(err error) {
//  		if recovered, is := errors.Unwrap(err).(errors.Recovered); is {
//  			log.Println(recovered.Cause())
//  		}
//  		log.Println(err)
//  	},
//  )
//
func Do(action func() error, handler func(error)) {
	var err error
	defer func() {
		if err != nil {
			handler(err)
		}
	}()
	defer func() {
		if r := recover(); r != nil {
			err = errors.WithStack(&recovered{r})
		}
	}()
	err = action()
}

type recovered struct {
	cause interface{}
}

// Error returns a string representation of the error.
func (r *recovered) Error() string {
	return "unexpected panic occurred"
}

// Cause returns the original cause of panic.
func (r *recovered) Cause() interface{} {
	return r.cause
}
