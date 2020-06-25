package errors

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Error struct {
	Type ErrorType
	Message string
}

type ErrorType string

const (
	StatusBadRequest ErrorType = "badrequest"
	StatusUnauthorized     ErrorType = "unauthorized"
	StatusForbidden        ErrorType = "forbidden"
	StatusNotFound         ErrorType = "notfound"
	StatusInternal         ErrorType = "internal"
	StatusUnknown          ErrorType = "unknown"
	StatusConflict		   ErrorType = "conflict"
)

func (e Error) Error() string {
	return string(e.Message)
}

func new(et ErrorType, msg string) Error {
	return Error{
		Type :et,
	}
}

func BadRequest(msg string) Error {
	return new(StatusBadRequest, msg)
}

func IsBadRequest(err error) bool {
	return isType(err, StatusBadRequest)
}

func Unauthorized(msg string) Error {
	return new(StatusUnauthorized, msg)
}

func IsUnauthorized(err error) bool {
	return isType(err, StatusUnauthorized)
}

func Forbidden(msg string) Error {
	return new(StatusForbidden, msg)
}

func IsForbidden(err error) bool {
	return isType(err, StatusForbidden)
}

func NotFound(msg string) Error {
	return new(StatusNotFound, msg)
}

func IsNotFound(err error) bool {
	return isType(err, StatusNotFound)
}

func Internal(msg string) Error {
	return new(StatusInternal, msg)
}

func IsInternal(err error) bool {
	return isType(err, StatusInternal)
}

func IsConflict(err error) bool {
	return isType(err, StatusConflict)
}

func GetType(err error) ErrorType {
	_, ok := err.(*errors.StatusError)

	if ok {
		return translateK8sType(err)
	}

	e, ok := err.(Error)

	if ok {
		return e.Type
	}

	return StatusUnknown
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
		return StatusNotFound
	case v1.StatusReasonBadRequest:
		return StatusBadRequest
	case v1.StatusReasonUnauthorized:
		return StatusUnauthorized
	case v1.StatusReasonForbidden:
		return StatusForbidden
	case v1.StatusReasonInternalError:
		return StatusInternal
	case v1.StatusReasonAlreadyExists:
		fallthrough
	case v1.StatusReasonConflict:
		return StatusConflict
	}

	return StatusInternal
}
