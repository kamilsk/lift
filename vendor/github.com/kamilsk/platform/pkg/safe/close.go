package safe

import "io"

// Close gracefully closes the Closer and calls the cleaners if an error occurred.
//
//  func handler(rw http.ResponseWriter, req *http.Request) {
//
//  	defer Close(req.Body, func(err error) { log.Println(err) })
//
//  	var data map[string]interface{}
//  	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
//  		rw.WriteHeader(http.StatusBadRequest)
//  		return
//  	}
//
//  	...
//  }
//
func Close(closer io.Closer, cleaners ...func(error)) {
	if err := closer.Close(); err != nil {
		for _, clean := range cleaners {
			clean(err)
		}
	}
}

// The Releaser type is an adapter to allow the use of ordinary functions
// as the io.Closer interface. If fn is a function with the appropriate signature,
// Releaser(fn) is a Closer that calls fn. It can be used by the Close function.
//
//  ticket, err := semaphore.Acquire(breaker.BreakByTimeout(time.Second))
//  if err != nil {
//  	log.Fatal(err)
//  }
//  defer Close(Releaser(ticket), func(err error) { log.Println(err) })
//
type Releaser func() error

// Close releases resources associated with the Releaser.
func (fn Releaser) Close() error {
	return fn()
}
