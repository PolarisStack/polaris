// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/metrics/v3/metrics_service.proto

package envoy_service_metrics_v3

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
var _metrics_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on StreamMetricsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamMetricsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamMetricsResponseValidationError is the validation error returned by
// StreamMetricsResponse.Validate if the designated constraints aren't met.
type StreamMetricsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamMetricsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamMetricsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamMetricsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamMetricsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamMetricsResponseValidationError) ErrorName() string {
	return "StreamMetricsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StreamMetricsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamMetricsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamMetricsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamMetricsResponseValidationError{}

// Validate checks the field values on StreamMetricsMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamMetricsMessage) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIdentifier()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamMetricsMessageValidationError{
				field:  "Identifier",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEnvoyMetrics() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamMetricsMessageValidationError{
					field:  fmt.Sprintf("EnvoyMetrics[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamMetricsMessageValidationError is the validation error returned by
// StreamMetricsMessage.Validate if the designated constraints aren't met.
type StreamMetricsMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamMetricsMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamMetricsMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamMetricsMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamMetricsMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamMetricsMessageValidationError) ErrorName() string {
	return "StreamMetricsMessageValidationError"
}

// Error satisfies the builtin error interface
func (e StreamMetricsMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamMetricsMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamMetricsMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamMetricsMessageValidationError{}

// Validate checks the field values on StreamMetricsMessage_Identifier with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamMetricsMessage_Identifier) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNode() == nil {
		return StreamMetricsMessage_IdentifierValidationError{
			field:  "Node",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamMetricsMessage_IdentifierValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// StreamMetricsMessage_IdentifierValidationError is the validation error
// returned by StreamMetricsMessage_Identifier.Validate if the designated
// constraints aren't met.
type StreamMetricsMessage_IdentifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamMetricsMessage_IdentifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamMetricsMessage_IdentifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamMetricsMessage_IdentifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamMetricsMessage_IdentifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamMetricsMessage_IdentifierValidationError) ErrorName() string {
	return "StreamMetricsMessage_IdentifierValidationError"
}

// Error satisfies the builtin error interface
func (e StreamMetricsMessage_IdentifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamMetricsMessage_Identifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamMetricsMessage_IdentifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamMetricsMessage_IdentifierValidationError{}
