package safe

import "go.octolab.org/errors"

// ErrBadCall is returned by a function with flexible contract if it is used incorrectly.
const ErrBadCall errors.Message = "bad function call"
