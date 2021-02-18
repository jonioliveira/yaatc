package error

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")

var ErrInvalidEntity = errors.New("Invalid entity")

// ErrCannotBeDeleted bookmark cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")
