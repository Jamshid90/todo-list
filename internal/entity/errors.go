package entity

import "fmt"

// error not found
type ErrNotFound struct {
	objectName string
	identifier string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s '%s' not found", e.objectName, e.identifier)
}

func NewErrNotFound(objectName, identifier string) *ErrNotFound {
	return &ErrNotFound{objectName, identifier}
}

// error conflict
type ErrConflict struct {
	objectName string
}

func (e *ErrConflict) Error() string {
	return e.objectName + " already exist"
}

func NewErrConflict(text string) *ErrConflict {
	return &ErrConflict{text}
}
