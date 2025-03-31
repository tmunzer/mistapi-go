package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SwitchPortUsageMacLimit represents a SwitchPortUsageMacLimit struct.
// Only if `mode`!=`dynamic` max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)
type SwitchPortUsageMacLimit struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for SwitchPortUsageMacLimit,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortUsageMacLimit) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsageMacLimit.
// It customizes the JSON marshaling process for SwitchPortUsageMacLimit objects.
func (s SwitchPortUsageMacLimit) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchPortUsageMacLimitContainer.From*` functions to initialize the SwitchPortUsageMacLimit object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsageMacLimit object to a map representation for JSON marshaling.
func (s *SwitchPortUsageMacLimit) toMap() any {
    switch obj := s.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsageMacLimit.
// It customizes the JSON unmarshaling process for SwitchPortUsageMacLimit objects.
func (s *SwitchPortUsageMacLimit) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(string), false, &s.isString),
    )
    
    s.value = result
    return err
}

func (s *SwitchPortUsageMacLimit) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SwitchPortUsageMacLimit) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

// internalSwitchPortUsageMacLimit represents a switchPortUsageMacLimit struct.
// Only if `mode`!=`dynamic` max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)
type internalSwitchPortUsageMacLimit struct {}

var SwitchPortUsageMacLimitContainer internalSwitchPortUsageMacLimit

// The internalSwitchPortUsageMacLimit instance, wrapping the provided int value.
func (s *internalSwitchPortUsageMacLimit) FromNumber(val int) SwitchPortUsageMacLimit {
    return SwitchPortUsageMacLimit{value: &val}
}

// The internalSwitchPortUsageMacLimit instance, wrapping the provided string value.
func (s *internalSwitchPortUsageMacLimit) FromString(val string) SwitchPortUsageMacLimit {
    return SwitchPortUsageMacLimit{value: &val}
}
