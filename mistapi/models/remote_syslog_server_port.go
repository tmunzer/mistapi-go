package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// RemoteSyslogServerPort represents a RemoteSyslogServerPort struct.
// Syslog Service Port, value from 1 to 65535
type RemoteSyslogServerPort struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for RemoteSyslogServerPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RemoteSyslogServerPort) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogServerPort.
// It customizes the JSON marshaling process for RemoteSyslogServerPort objects.
func (r RemoteSyslogServerPort) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RemoteSyslogServerPortContainer.From*` functions to initialize the RemoteSyslogServerPort object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogServerPort object to a map representation for JSON marshaling.
func (r *RemoteSyslogServerPort) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogServerPort.
// It customizes the JSON unmarshaling process for RemoteSyslogServerPort objects.
func (r *RemoteSyslogServerPort) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RemoteSyslogServerPort) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *RemoteSyslogServerPort) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRemoteSyslogServerPort represents a remoteSyslogServerPort struct.
// Syslog Service Port, value from 1 to 65535
type internalRemoteSyslogServerPort struct {}

var RemoteSyslogServerPortContainer internalRemoteSyslogServerPort

// The internalRemoteSyslogServerPort instance, wrapping the provided int value.
func (r *internalRemoteSyslogServerPort) FromNumber(val int) RemoteSyslogServerPort {
    return RemoteSyslogServerPort{value: &val}
}

// The internalRemoteSyslogServerPort instance, wrapping the provided string value.
func (r *internalRemoteSyslogServerPort) FromString(val string) RemoteSyslogServerPort {
    return RemoteSyslogServerPort{value: &val}
}
