package errors

// Message represents a textual error.
//
//  const ErrInterrupted errors.Message = "operation interrupted"
//
//  func Action() error {
//  	...
//  }
//
//  switch err := Action(); err {
//  case ErrInterrupted:
//  	http.Error(..., err.Error(), http.StatusRequestTimeout)
//  case ...:
//  	http.Error(..., http.StatusInternalServerError)
//  }
//
type Message string

// Error returns a string representation of the error.
func (err Message) Error() string {
	return string(err)
}

// Recovered represents a recovered after panic error.
type Recovered interface {
	error
	// Cause returns the original cause of panic.
	Cause() interface{}
}

// Retriable represents a retriable action error.
//
// It is compatible with github.com/kamilsk/retry (v4 and later).
type Retriable interface {
	error
	Retriable() bool // Is the error retriable?
}

// Unwrap returns the result of calling the Unwrap or Cause methods
// on the error, otherwise it returns error itself.
//
//  func Caller(req *http.Request) error {
//  	resp, err := http.DefaultClient.Do(req)
//  	if err != nil {
//  		return errors.WithStack(fmt.Errorf("caller: %w", err))
//  	}
//  	...
//  }
//
//  if err, is := Unwrap(Caller(req)).(net.Error); is {
//  	...
//  }
//
// It is compatible with github.com/pkg/errors
// and built-in errors since 1.13.
func Unwrap(err error) error {
	// compatible with github.com/pkg/errors
	type causer interface {
		Cause() error
	}
	// compatible with built-in errors since 1.13
	type wrapper interface {
		Unwrap() error
	}

	for err != nil {
		layer, is := err.(wrapper)
		if is {
			err = layer.Unwrap()
			continue
		}
		cause, is := err.(causer)
		if is {
			err = cause.Cause()
			continue
		}
		break
	}
	return err
}
