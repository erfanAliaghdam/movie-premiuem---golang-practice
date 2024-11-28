package custom_errors

import "errors"

var AlreadyExists = errors.New("already exists")
var NotExists = errors.New("does not exists")
