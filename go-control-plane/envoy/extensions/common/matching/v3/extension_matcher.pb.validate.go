// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/matching/v3/extension_matcher.proto

package envoy_extensions_common_matching_v3

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on ExtensionWithMatcher with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExtensionWithMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtensionWithMatcherValidationError{
				field:  "Matcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetXdsMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtensionWithMatcherValidationError{
				field:  "XdsMatcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetExtensionConfig() == nil {
		return ExtensionWithMatcherValidationError{
			field:  "ExtensionConfig",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetExtensionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtensionWithMatcherValidationError{
				field:  "ExtensionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExtensionWithMatcherValidationError is the validation error returned by
// ExtensionWithMatcher.Validate if the designated constraints aren't met.
type ExtensionWithMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtensionWithMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtensionWithMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtensionWithMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtensionWithMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtensionWithMatcherValidationError) ErrorName() string {
	return "ExtensionWithMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e ExtensionWithMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtensionWithMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtensionWithMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtensionWithMatcherValidationError{}
