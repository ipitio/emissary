// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/health_checkers/redis/v3/redis.proto

package redisv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Redis with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Redis) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Redis with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RedisMultiError, or nil if none found.
func (m *Redis) ValidateAll() error {
	return m.validate(true)
}

func (m *Redis) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return RedisMultiError(errors)
	}

	return nil
}

// RedisMultiError is an error wrapping multiple validation errors returned by
// Redis.ValidateAll() if the designated constraints aren't met.
type RedisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedisMultiError) AllErrors() []error { return m }

// RedisValidationError is the validation error returned by Redis.Validate if
// the designated constraints aren't met.
type RedisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisValidationError) ErrorName() string { return "RedisValidationError" }

// Error satisfies the builtin error interface
func (e RedisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisValidationError{}
