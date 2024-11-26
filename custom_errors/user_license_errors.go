package custom_errors

import "errors"

var (
	ErrUserHasActiveLicense = errors.New("user has an active license")
)
