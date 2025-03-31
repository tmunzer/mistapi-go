package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SwitchPortUsageMtu represents a SwitchPortUsageMtu struct.
// Only if `mode`!=`dynamic` media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514.
type SwitchPortUsageMtu struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for SwitchPortUsageMtu,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortUsageMtu) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsageMtu.
// It customizes the JSON marshaling process for SwitchPortUsageMtu objects.
func (s SwitchPortUsageMtu) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchPortUsageMtuContainer.From*` functions to initialize the SwitchPortUsageMtu object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsageMtu object to a map representation for JSON marshaling.
func (s *SwitchPortUsageMtu) toMap() any {
    switch obj := s.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsageMtu.
// It customizes the JSON unmarshaling process for SwitchPortUsageMtu objects.
func (s *SwitchPortUsageMtu) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(string), false, &s.isString),
    )
    
    s.value = result
    return err
}

func (s *SwitchPortUsageMtu) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SwitchPortUsageMtu) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

// internalSwitchPortUsageMtu represents a switchPortUsageMtu struct.
// Only if `mode`!=`dynamic` media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514.
type internalSwitchPortUsageMtu struct {}

var SwitchPortUsageMtuContainer internalSwitchPortUsageMtu

// The internalSwitchPortUsageMtu instance, wrapping the provided int value.
func (s *internalSwitchPortUsageMtu) FromNumber(val int) SwitchPortUsageMtu {
    return SwitchPortUsageMtu{value: &val}
}

// The internalSwitchPortUsageMtu instance, wrapping the provided string value.
func (s *internalSwitchPortUsageMtu) FromString(val string) SwitchPortUsageMtu {
    return SwitchPortUsageMtu{value: &val}
}
