// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: video/sfu/models/models.proto

package sfu_models

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

// Validate checks the field values on CallState with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CallState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallState with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CallStateMultiError, or nil
// if none found.
func (m *CallState) ValidateAll() error {
	return m.validate(true)
}

func (m *CallState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetParticipants() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CallStateValidationError{
						field:  fmt.Sprintf("Participants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CallStateValidationError{
						field:  fmt.Sprintf("Participants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CallStateValidationError{
					field:  fmt.Sprintf("Participants[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CallStateMultiError(errors)
	}

	return nil
}

// CallStateMultiError is an error wrapping multiple validation errors returned
// by CallState.ValidateAll() if the designated constraints aren't met.
type CallStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallStateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallStateMultiError) AllErrors() []error { return m }

// CallStateValidationError is the validation error returned by
// CallState.Validate if the designated constraints aren't met.
type CallStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallStateValidationError) ErrorName() string { return "CallStateValidationError" }

// Error satisfies the builtin error interface
func (e CallStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallStateValidationError{}

// Validate checks the field values on Participant with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Participant) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Participant with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ParticipantMultiError, or
// nil if none found.
func (m *Participant) ValidateAll() error {
	return m.validate(true)
}

func (m *Participant) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for SessionId

	if all {
		switch v := interface{}(m.GetJoinedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ParticipantValidationError{
					field:  "JoinedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ParticipantValidationError{
					field:  "JoinedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetJoinedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ParticipantValidationError{
				field:  "JoinedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TrackLookupPrefix

	// no validation rules for ConnectionQuality

	// no validation rules for IsSpeaking

	// no validation rules for IsDominantSpeaker

	// no validation rules for AudioLevel

	// no validation rules for Name

	// no validation rules for Image

	if all {
		switch v := interface{}(m.GetCustom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ParticipantValidationError{
					field:  "Custom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ParticipantValidationError{
					field:  "Custom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCustom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ParticipantValidationError{
				field:  "Custom",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ParticipantMultiError(errors)
	}

	return nil
}

// ParticipantMultiError is an error wrapping multiple validation errors
// returned by Participant.ValidateAll() if the designated constraints aren't met.
type ParticipantMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ParticipantMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ParticipantMultiError) AllErrors() []error { return m }

// ParticipantValidationError is the validation error returned by
// Participant.Validate if the designated constraints aren't met.
type ParticipantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParticipantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParticipantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParticipantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParticipantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParticipantValidationError) ErrorName() string { return "ParticipantValidationError" }

// Error satisfies the builtin error interface
func (e ParticipantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParticipant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParticipantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParticipantValidationError{}

// Validate checks the field values on StreamQuality with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamQuality) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamQuality with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamQualityMultiError, or
// nil if none found.
func (m *StreamQuality) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamQuality) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for VideoQuality

	// no validation rules for UserId

	if len(errors) > 0 {
		return StreamQualityMultiError(errors)
	}

	return nil
}

// StreamQualityMultiError is an error wrapping multiple validation errors
// returned by StreamQuality.ValidateAll() if the designated constraints
// aren't met.
type StreamQualityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamQualityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamQualityMultiError) AllErrors() []error { return m }

// StreamQualityValidationError is the validation error returned by
// StreamQuality.Validate if the designated constraints aren't met.
type StreamQualityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamQualityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamQualityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamQualityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamQualityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamQualityValidationError) ErrorName() string { return "StreamQualityValidationError" }

// Error satisfies the builtin error interface
func (e StreamQualityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamQuality.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamQualityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamQualityValidationError{}

// Validate checks the field values on VideoDimension with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VideoDimension) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VideoDimension with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VideoDimensionMultiError,
// or nil if none found.
func (m *VideoDimension) ValidateAll() error {
	return m.validate(true)
}

func (m *VideoDimension) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Width

	// no validation rules for Height

	if len(errors) > 0 {
		return VideoDimensionMultiError(errors)
	}

	return nil
}

// VideoDimensionMultiError is an error wrapping multiple validation errors
// returned by VideoDimension.ValidateAll() if the designated constraints
// aren't met.
type VideoDimensionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VideoDimensionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VideoDimensionMultiError) AllErrors() []error { return m }

// VideoDimensionValidationError is the validation error returned by
// VideoDimension.Validate if the designated constraints aren't met.
type VideoDimensionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoDimensionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoDimensionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoDimensionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoDimensionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoDimensionValidationError) ErrorName() string { return "VideoDimensionValidationError" }

// Error satisfies the builtin error interface
func (e VideoDimensionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideoDimension.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoDimensionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoDimensionValidationError{}

// Validate checks the field values on VideoLayer with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VideoLayer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VideoLayer with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VideoLayerMultiError, or
// nil if none found.
func (m *VideoLayer) ValidateAll() error {
	return m.validate(true)
}

func (m *VideoLayer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Rid

	if all {
		switch v := interface{}(m.GetVideoDimension()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VideoLayerValidationError{
					field:  "VideoDimension",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VideoLayerValidationError{
					field:  "VideoDimension",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVideoDimension()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VideoLayerValidationError{
				field:  "VideoDimension",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Bitrate

	// no validation rules for Fps

	// no validation rules for Quality

	if len(errors) > 0 {
		return VideoLayerMultiError(errors)
	}

	return nil
}

// VideoLayerMultiError is an error wrapping multiple validation errors
// returned by VideoLayer.ValidateAll() if the designated constraints aren't met.
type VideoLayerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VideoLayerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VideoLayerMultiError) AllErrors() []error { return m }

// VideoLayerValidationError is the validation error returned by
// VideoLayer.Validate if the designated constraints aren't met.
type VideoLayerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoLayerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoLayerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoLayerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoLayerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoLayerValidationError) ErrorName() string { return "VideoLayerValidationError" }

// Error satisfies the builtin error interface
func (e VideoLayerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideoLayer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoLayerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoLayerValidationError{}

// Validate checks the field values on Codec with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Codec) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Codec with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CodecMultiError, or nil if none found.
func (m *Codec) ValidateAll() error {
	return m.validate(true)
}

func (m *Codec) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PayloadType

	// no validation rules for Name

	// no validation rules for FmtpLine

	// no validation rules for ClockRate

	// no validation rules for EncodingParameters

	if len(errors) > 0 {
		return CodecMultiError(errors)
	}

	return nil
}

// CodecMultiError is an error wrapping multiple validation errors returned by
// Codec.ValidateAll() if the designated constraints aren't met.
type CodecMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CodecMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CodecMultiError) AllErrors() []error { return m }

// CodecValidationError is the validation error returned by Codec.Validate if
// the designated constraints aren't met.
type CodecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CodecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CodecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CodecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CodecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CodecValidationError) ErrorName() string { return "CodecValidationError" }

// Error satisfies the builtin error interface
func (e CodecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCodec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CodecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CodecValidationError{}

// Validate checks the field values on ICETrickle with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ICETrickle) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ICETrickle with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ICETrickleMultiError, or
// nil if none found.
func (m *ICETrickle) ValidateAll() error {
	return m.validate(true)
}

func (m *ICETrickle) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PeerType

	// no validation rules for IceCandidate

	// no validation rules for SessionId

	if len(errors) > 0 {
		return ICETrickleMultiError(errors)
	}

	return nil
}

// ICETrickleMultiError is an error wrapping multiple validation errors
// returned by ICETrickle.ValidateAll() if the designated constraints aren't met.
type ICETrickleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ICETrickleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ICETrickleMultiError) AllErrors() []error { return m }

// ICETrickleValidationError is the validation error returned by
// ICETrickle.Validate if the designated constraints aren't met.
type ICETrickleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ICETrickleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ICETrickleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ICETrickleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ICETrickleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ICETrickleValidationError) ErrorName() string { return "ICETrickleValidationError" }

// Error satisfies the builtin error interface
func (e ICETrickleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sICETrickle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ICETrickleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ICETrickleValidationError{}

// Validate checks the field values on TrackInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TrackInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TrackInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TrackInfoMultiError, or nil
// if none found.
func (m *TrackInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *TrackInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TrackId

	// no validation rules for TrackType

	for idx, item := range m.GetLayers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TrackInfoValidationError{
						field:  fmt.Sprintf("Layers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TrackInfoValidationError{
						field:  fmt.Sprintf("Layers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TrackInfoValidationError{
					field:  fmt.Sprintf("Layers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Mid

	// no validation rules for Dtx

	// no validation rules for Stereo

	// no validation rules for Red

	if len(errors) > 0 {
		return TrackInfoMultiError(errors)
	}

	return nil
}

// TrackInfoMultiError is an error wrapping multiple validation errors returned
// by TrackInfo.ValidateAll() if the designated constraints aren't met.
type TrackInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TrackInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TrackInfoMultiError) AllErrors() []error { return m }

// TrackInfoValidationError is the validation error returned by
// TrackInfo.Validate if the designated constraints aren't met.
type TrackInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrackInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrackInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrackInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrackInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrackInfoValidationError) ErrorName() string { return "TrackInfoValidationError" }

// Error satisfies the builtin error interface
func (e TrackInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrackInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrackInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrackInfoValidationError{}

// Validate checks the field values on Call with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Call) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Call with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CallMultiError, or nil if none found.
func (m *Call) ValidateAll() error {
	return m.validate(true)
}

func (m *Call) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Id

	// no validation rules for CreatedByUserId

	// no validation rules for HostUserId

	if all {
		switch v := interface{}(m.GetCustom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CallValidationError{
					field:  "Custom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CallValidationError{
					field:  "Custom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCustom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CallValidationError{
				field:  "Custom",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CallValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CallValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CallValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CallValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CallValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CallValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CallMultiError(errors)
	}

	return nil
}

// CallMultiError is an error wrapping multiple validation errors returned by
// Call.ValidateAll() if the designated constraints aren't met.
type CallMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallMultiError) AllErrors() []error { return m }

// CallValidationError is the validation error returned by Call.Validate if the
// designated constraints aren't met.
type CallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallValidationError) ErrorName() string { return "CallValidationError" }

// Error satisfies the builtin error interface
func (e CallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallValidationError{}

// Validate checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Error) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ErrorMultiError, or nil if none found.
func (m *Error) ValidateAll() error {
	return m.validate(true)
}

func (m *Error) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for ShouldRetry

	if len(errors) > 0 {
		return ErrorMultiError(errors)
	}

	return nil
}

// ErrorMultiError is an error wrapping multiple validation errors returned by
// Error.ValidateAll() if the designated constraints aren't met.
type ErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ErrorMultiError) AllErrors() []error { return m }

// ErrorValidationError is the validation error returned by Error.Validate if
// the designated constraints aren't met.
type ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorValidationError) ErrorName() string { return "ErrorValidationError" }

// Error satisfies the builtin error interface
func (e ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorValidationError{}

// Validate checks the field values on ClientDetails with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientDetails) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientDetails with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientDetailsMultiError, or
// nil if none found.
func (m *ClientDetails) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientDetails) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	// no validation rules for Os

	// no validation rules for Browser

	if len(errors) > 0 {
		return ClientDetailsMultiError(errors)
	}

	return nil
}

// ClientDetailsMultiError is an error wrapping multiple validation errors
// returned by ClientDetails.ValidateAll() if the designated constraints
// aren't met.
type ClientDetailsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientDetailsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientDetailsMultiError) AllErrors() []error { return m }

// ClientDetailsValidationError is the validation error returned by
// ClientDetails.Validate if the designated constraints aren't met.
type ClientDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientDetailsValidationError) ErrorName() string { return "ClientDetailsValidationError" }

// Error satisfies the builtin error interface
func (e ClientDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientDetailsValidationError{}
