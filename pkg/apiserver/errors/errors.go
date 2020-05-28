package errors

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Error struct {
	Type ErrorType
}

type ErrorType string

const (
	BadRequest   ErrorType = "badrequest"
	Unauthorized ErrorType = "unauthorized"
	Forbidden    ErrorType = "forbidden"
	NotFound	 ErrorType = "notfound"
	Internal	 ErrorType = "internal"
)

func (e Error) Error() string {
	return string(e.Type)
}

func New(et ErrorType) Error {
	return Error{
		Type :et,
	}
}

func IsBadRequest(err error) bool {
	return isType(err, BadRequest)
}

func IsUnauthorized(err error) bool {
	return isType(err, Unauthorized)
}

func IsForbidden(err error) bool {
	return isType(err, Forbidden)
}

func IsNotFound(err error) bool {
	return isType(err, NotFound)
}

func IsInternal(err error) bool {
	return isType(err, Internal)
}

func isType(err error, et ErrorType) bool {
	// check if this is a k8s error
	e, ok := err.(*errors.StatusError)

	if ok {
		// this is a k8s error
		return translateK8sType(e) == et
	}

	e2, ok := err.(Error)

	if ok {
		return e2.Type == et
	}

	return false
}

func translateK8sType(err error) ErrorType {
	switch errors.ReasonForError(err) {
	case v1.StatusReasonNotFound:
		return NotFound
	case v1.StatusReasonBadRequest:
		return BadRequest
	case v1.StatusReasonUnauthorized:
		return Unauthorized
	case v1.StatusReasonForbidden:
		return Forbidden
	case v1.StatusReasonInternalError:
		return Internal
	}

	return Internal
}
