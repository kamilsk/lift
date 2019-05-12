package safe

import "github.com/pkg/errors"

// Do runs the action and captures a panic as its error.
//
//  serve := make(chan error, 1)
//
//  go Safe(func() error {
//  	return server.ListenAndServe()
//  }, func(err error) {
//  	serve <- errors.Wrap(err, "tried to listen and serve a connection")
//  	close(serve)
//  })
//
func Do(action func() error, closer func(error)) {
	var err error
	defer func() { closer(err) }()
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("unexpected panic handled: %+v", r)
		}
	}()
	err = action()
}
