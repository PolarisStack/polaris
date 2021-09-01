// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/core/health_check.proto

package envoy_api_v2_core

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

	_type "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/type"
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

	_ = _type.CodecClientType(0)
)

// define the regex for a UUID once up-front
var _health_check_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTimeout() == nil {
		return HealthCheckValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
	}

	if d := m.GetTimeout(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "Timeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "Timeout",
				reason: "value must be greater than 0s",
			}
		}

	}

	if m.GetInterval() == nil {
		return HealthCheckValidationError{
			field:  "Interval",
			reason: "value is required",
		}
	}

	if d := m.GetInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "Interval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if v, ok := interface{}(m.GetInitialJitter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "InitialJitter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetIntervalJitter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "IntervalJitter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IntervalJitterPercent

	if m.GetUnhealthyThreshold() == nil {
		return HealthCheckValidationError{
			field:  "UnhealthyThreshold",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUnhealthyThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "UnhealthyThreshold",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetHealthyThreshold() == nil {
		return HealthCheckValidationError{
			field:  "HealthyThreshold",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHealthyThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "HealthyThreshold",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAltPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "AltPort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetReuseConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "ReuseConnection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetNoTrafficInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "NoTrafficInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "NoTrafficInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if d := m.GetUnhealthyInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "UnhealthyInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "UnhealthyInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if d := m.GetUnhealthyEdgeInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "UnhealthyEdgeInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "UnhealthyEdgeInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if d := m.GetHealthyEdgeInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "HealthyEdgeInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "HealthyEdgeInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	// no validation rules for EventLogPath

	if v, ok := interface{}(m.GetEventService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "EventService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AlwaysLogHealthCheckFailures

	if v, ok := interface{}(m.GetTlsOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "TlsOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.HealthChecker.(type) {

	case *HealthCheck_HttpHealthCheck_:

		if v, ok := interface{}(m.GetHttpHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "HttpHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_TcpHealthCheck_:

		if v, ok := interface{}(m.GetTcpHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "TcpHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_GrpcHealthCheck_:

		if v, ok := interface{}(m.GetGrpcHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "GrpcHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_CustomHealthCheck_:

		if v, ok := interface{}(m.GetCustomHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "CustomHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return HealthCheckValidationError{
			field:  "HealthChecker",
			reason: "value is required",
		}

	}

	return nil
}

// HealthCheckValidationError is the validation error returned by
// HealthCheck.Validate if the designated constraints aren't met.
type HealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckValidationError) ErrorName() string { return "HealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e HealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckValidationError{}

// Validate checks the field values on HealthCheck_Payload with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_Payload) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Payload.(type) {

	case *HealthCheck_Payload_Text:

		if len(m.GetText()) < 1 {
			return HealthCheck_PayloadValidationError{
				field:  "Text",
				reason: "value length must be at least 1 bytes",
			}
		}

	case *HealthCheck_Payload_Binary:
		// no validation rules for Binary

	default:
		return HealthCheck_PayloadValidationError{
			field:  "Payload",
			reason: "value is required",
		}

	}

	return nil
}

// HealthCheck_PayloadValidationError is the validation error returned by
// HealthCheck_Payload.Validate if the designated constraints aren't met.
type HealthCheck_PayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_PayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_PayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_PayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_PayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_PayloadValidationError) ErrorName() string {
	return "HealthCheck_PayloadValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_PayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_Payload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_PayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_PayloadValidationError{}

// Validate checks the field values on HealthCheck_HttpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_HttpHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Host

	if len(m.GetPath()) < 1 {
		return HealthCheck_HttpHealthCheckValidationError{
			field:  "Path",
			reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetSend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheck_HttpHealthCheckValidationError{
				field:  "Send",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetReceive()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheck_HttpHealthCheckValidationError{
				field:  "Receive",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ServiceName

	if len(m.GetRequestHeadersToAdd()) > 1000 {
		return HealthCheck_HttpHealthCheckValidationError{
			field:  "RequestHeadersToAdd",
			reason: "value must contain no more than 1000 item(s)",
		}
	}

	for idx, item := range m.GetRequestHeadersToAdd() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_HttpHealthCheckValidationError{
					field:  fmt.Sprintf("RequestHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for UseHttp2

	for idx, item := range m.GetExpectedStatuses() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_HttpHealthCheckValidationError{
					field:  fmt.Sprintf("ExpectedStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if _, ok := _type.CodecClientType_name[int32(m.GetCodecClientType())]; !ok {
		return HealthCheck_HttpHealthCheckValidationError{
			field:  "CodecClientType",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetServiceNameMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheck_HttpHealthCheckValidationError{
				field:  "ServiceNameMatcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HealthCheck_HttpHealthCheckValidationError is the validation error returned
// by HealthCheck_HttpHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_HttpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_HttpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_HttpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_HttpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_HttpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_HttpHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_HttpHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_HttpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_HttpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_HttpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_HttpHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_TcpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_TcpHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheck_TcpHealthCheckValidationError{
				field:  "Send",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetReceive() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_TcpHealthCheckValidationError{
					field:  fmt.Sprintf("Receive[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HealthCheck_TcpHealthCheckValidationError is the validation error returned
// by HealthCheck_TcpHealthCheck.Validate if the designated constraints aren't met.
type HealthCheck_TcpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_TcpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_TcpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_TcpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_TcpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_TcpHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_TcpHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_TcpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_TcpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_TcpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_TcpHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_RedisHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_RedisHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Key

	return nil
}

// HealthCheck_RedisHealthCheckValidationError is the validation error returned
// by HealthCheck_RedisHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_RedisHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_RedisHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_RedisHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_RedisHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_RedisHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_RedisHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_RedisHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_RedisHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_RedisHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_RedisHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_RedisHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_GrpcHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_GrpcHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServiceName

	// no validation rules for Authority

	return nil
}

// HealthCheck_GrpcHealthCheckValidationError is the validation error returned
// by HealthCheck_GrpcHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_GrpcHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_GrpcHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_GrpcHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_GrpcHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_GrpcHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_GrpcHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_GrpcHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_GrpcHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_GrpcHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_GrpcHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_GrpcHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_CustomHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_CustomHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return HealthCheck_CustomHealthCheckValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *HealthCheck_CustomHealthCheck_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_CustomHealthCheckValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_CustomHealthCheck_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_CustomHealthCheckValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HealthCheck_CustomHealthCheckValidationError is the validation error
// returned by HealthCheck_CustomHealthCheck.Validate if the designated
// constraints aren't met.
type HealthCheck_CustomHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_CustomHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_CustomHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_CustomHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_CustomHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_CustomHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_CustomHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_CustomHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_CustomHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_CustomHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_CustomHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_TlsOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_TlsOptions) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// HealthCheck_TlsOptionsValidationError is the validation error returned by
// HealthCheck_TlsOptions.Validate if the designated constraints aren't met.
type HealthCheck_TlsOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_TlsOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_TlsOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_TlsOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_TlsOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_TlsOptionsValidationError) ErrorName() string {
	return "HealthCheck_TlsOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_TlsOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_TlsOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_TlsOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_TlsOptionsValidationError{}
