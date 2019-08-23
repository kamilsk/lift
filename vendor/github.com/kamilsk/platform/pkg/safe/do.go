package safe

import "github.com/pkg/errors"

// Do reliably runs the action and captures panic as its error.
//
//  serve := make(chan error, 1)
//
//  go safe.Do(func() error {
//  	return server.ListenAndServe()
//  }, func(err error) {
//  	serve <- errors.Wrap(err, "tried to listen and serve a connection")
//  	close(serve)
//  })
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
			err = errors.Errorf("panic unexpected: %#+v", r)
		}
	}()
	err = action()
}
