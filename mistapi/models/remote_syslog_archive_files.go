// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RemoteSyslogArchiveFiles represents a RemoteSyslogArchiveFiles struct.
type RemoteSyslogArchiveFiles struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for RemoteSyslogArchiveFiles,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RemoteSyslogArchiveFiles) String() string {
	return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogArchiveFiles.
// It customizes the JSON marshaling process for RemoteSyslogArchiveFiles objects.
func (r RemoteSyslogArchiveFiles) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.RemoteSyslogArchiveFilesContainer.From*` functions to initialize the RemoteSyslogArchiveFiles object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogArchiveFiles object to a map representation for JSON marshaling.
func (r *RemoteSyslogArchiveFiles) toMap() any {
	switch obj := r.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogArchiveFiles.
// It customizes the JSON unmarshaling process for RemoteSyslogArchiveFiles objects.
func (r *RemoteSyslogArchiveFiles) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(string), false, &r.isString),
		NewTypeHolder(new(int), false, &r.isNumber),
	)

	r.value = result
	return err
}

func (r *RemoteSyslogArchiveFiles) AsString() (
	*string,
	bool) {
	if !r.isString {
		return nil, false
	}
	return r.value.(*string), true
}

func (r *RemoteSyslogArchiveFiles) AsNumber() (
	*int,
	bool) {
	if !r.isNumber {
		return nil, false
	}
	return r.value.(*int), true
}

// internalRemoteSyslogArchiveFiles represents a remoteSyslogArchiveFiles struct.
type internalRemoteSyslogArchiveFiles struct{}

var RemoteSyslogArchiveFilesContainer internalRemoteSyslogArchiveFiles

// The internalRemoteSyslogArchiveFiles instance, wrapping the provided string value.
func (r *internalRemoteSyslogArchiveFiles) FromString(val string) RemoteSyslogArchiveFiles {
	return RemoteSyslogArchiveFiles{value: &val}
}

// The internalRemoteSyslogArchiveFiles instance, wrapping the provided int value.
func (r *internalRemoteSyslogArchiveFiles) FromNumber(val int) RemoteSyslogArchiveFiles {
	return RemoteSyslogArchiveFiles{value: &val}
}
