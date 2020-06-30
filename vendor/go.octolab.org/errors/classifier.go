// +build go1.13

package errors

import (
	"errors"
	"net"
	"strings"
)

// Classifier provides functionality to classify errors
// and represents them as a string, e.g. for metrics system.
type Classifier map[string][]error

// Classify classifies the error and returns its class name.
// If it cannot to classify it returns the Unknown.
//
//  func (service *Service) Do(ctx context.Context, payload interface{}) {
//  	resp, err := service.proxy.Call(ctx, Data{Payload: payload})
//  	if err != nil {
//  		service.telemetry.Increment(global.Classifier.Classify(err, errors.Unknown))
//  		...
//  	}
//  	...
//  }
//
func (classifier Classifier) Classify(err error) string {
	if err = Unwrap(err); err == nil {
		return Unknown
	}
	for class, list := range classifier {
		for _, target := range list {
			if errors.Is(target, err) {
				return class
			}
		}
	}
	return Unknown
}

// ClassifyAs unwraps the errors and stores them with the class name.
//
//  classifier := make(errors.Classifier).
//  	ClassifyAs("network", new(errors.NetworkError)).
//  	ClassifyAs("fs", os.ErrExist, os.ErrNotExist)
//
func (classifier Classifier) ClassifyAs(class string, list ...error) Classifier {
	for _, err := range list {
		if err = Unwrap(err); err == nil {
			continue
		}
		var present bool
		for i, matcher := range classifier[class] {
			present = errors.Is(matcher, err)
			if !present && errors.Is(err, matcher) {
				classifier[class][i], present = err, true // subset case
			}
			if present {
				break
			}
		}
		if !present {
			classifier[class] = append(classifier[class], err)
		}
	}
	return classifier
}

// Consistent checks that different groups don't contain similar errors.
func (classifier Classifier) Consistent() bool {
	var total int
	for _, list := range classifier {
		total += len(list)
	}
	if total == 0 {
		return true
	}
	flat := make([]error, 0, total)
	for _, list := range classifier {
		flat = append(flat, list...)
	}
	for i, err := range flat {
		for _, target := range flat[i+1:] {
			if errors.Is(err, target) || errors.Is(target, err) {
				return false
			}
		}
	}
	return true
}

// Merge returns a copy of the current Classifier with data
// from passed classifiers.
func (classifier Classifier) Merge(classifiers ...Classifier) Classifier {
	dst := make(Classifier, len(classifier))
	for k, v := range classifier {
		dst[k] = make([]error, len(v))
		copy(dst[k], v)
	}
	for _, src := range classifiers {
		for class, list := range src {
			dst.ClassifyAs(class, list...)
		}
	}
	return dst
}

// Unknown can be used as the fallback class name of error classification.
const Unknown = "unknown"

// MessageError can check errors by their error message.
type MessageError struct{ Message string }

func (matcher MessageError) Error() string { return matcher.Message }

// Is reports whether the error matches message error class.
func (matcher MessageError) Is(err error) bool {
	return err != nil && strings.Contains(err.Error(), matcher.Message)
}

// NetworkError can check network errors.
type NetworkError struct{}

func (*NetworkError) Error() string   { return "network error" }
func (*NetworkError) Temporary() bool { return true }
func (*NetworkError) Timeout() bool   { return true }

// Is reports whether the error matches network error class.
func (*NetworkError) Is(err error) bool {
	_, is := err.(net.Error)
	return is
}

// RecoveredError can check recovered errors.
type RecoveredError struct{}

func (*RecoveredError) Error() string      { return "recovered after panic" }
func (*RecoveredError) Cause() interface{} { return nil }

// Is reports whether the error matches recovered error class.
func (*RecoveredError) Is(err error) bool {
	_, is := err.(Recovered)
	return is
}

type RetriableError struct{}

func (*RetriableError) Error() string   { return "retriable action error" }
func (*RetriableError) Retriable() bool { return true }

// Is reports whether the error matches retriable error class.
func (*RetriableError) Is(err error) bool {
	casted, is := err.(Retriable)
	return is && casted.Retriable()
}

// TemporaryError can check temporary errors.
type TemporaryError struct{}

func (*TemporaryError) Error() string   { return "temporary error" }
func (*TemporaryError) Temporary() bool { return true }

// Is reports whether the error matches temporary error class.
func (*TemporaryError) Is(err error) bool {
	casted, is := err.(interface{ Temporary() bool })
	return is && casted.Temporary()
}

// TimeoutError can check timeout errors.
type TimeoutError struct{}

func (*TimeoutError) Error() string { return "timeout error" }
func (*TimeoutError) Timeout() bool { return true }

// Is reports whether the error matches timeout error class.
func (*TimeoutError) Is(err error) bool {
	casted, is := err.(interface{ Timeout() bool })
	return is && casted.Timeout()
}
