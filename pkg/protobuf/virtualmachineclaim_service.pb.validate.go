// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hobbyfarm/virtualmachineclaim_service.proto

package protobuf

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _virtualmachineclaim_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on VirtualMachineClaimList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *VirtualMachineClaimList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetVirtualMachineClaims() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return VirtualMachineClaimListValidationError{
					field:  fmt.Sprintf("VirtualMachineClaims[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// VirtualMachineClaimListValidationError is the validation error returned by
// VirtualMachineClaimList.Validate if the designated constraints aren't met.
type VirtualMachineClaimListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VirtualMachineClaimListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VirtualMachineClaimListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VirtualMachineClaimListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VirtualMachineClaimListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VirtualMachineClaimListValidationError) ErrorName() string {
	return "VirtualMachineClaimListValidationError"
}

// Error satisfies the builtin error interface
func (e VirtualMachineClaimListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVirtualMachineClaimList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VirtualMachineClaimListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VirtualMachineClaimListValidationError{}
