package auth

import "errors"

var ErrNoAuthHeaderIncluded = errors.New("no auth header included in request")
