package safe

import "io"

// Close gracefully closes the io.Closer and calls the cleaners if an error occurred.
//
//  func handler(rw http.ResponseWriter, req *http.Request) {
//
//  	defer safe.Close(req.Body, func(err error) { log.Println(err) })
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
// Deprecated: use go.octolab.org/safe.Close instead.
func Close(closer io.Closer, cleaners ...func(error)) {
	if err := closer.Close(); err != nil {
		for _, clean := range cleaners {
			clean(err)
		}
	}
}

// The Closer type is an adapter to allow the use of ordinary functions
// as the io.Closer interface. If fn is a function with the appropriate signature,
// Closer(fn) is a Closer that calls fn. It can be used by the Close function.
//
//  ticket, err := semaphore.Acquire(breaker.BreakByTimeout(time.Second))
//  if err != nil {
//  	log.Fatal(err)
//  }
//  defer safe.Close(safe.Closer(ticket), func(err error) { log.Println(err) })
//
// Deprecated: use go.octolab.org/safe.Closer instead.
type Closer func() error

// Close releases resources associated with the Closer.
func (fn Closer) Close() error {
	return fn()
}
