package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SponsorLinkValidityDuration represents a SponsorLinkValidityDuration struct.
// Optional if `sponsor_enabled`==`true`. How long to remain valid sponsored guest request approve/deny link received in email, in minutes. Value is between 5 and 60.
type SponsorLinkValidityDuration struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for SponsorLinkValidityDuration,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SponsorLinkValidityDuration) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SponsorLinkValidityDuration.
// It customizes the JSON marshaling process for SponsorLinkValidityDuration objects.
func (s SponsorLinkValidityDuration) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SponsorLinkValidityDurationContainer.From*` functions to initialize the SponsorLinkValidityDuration object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SponsorLinkValidityDuration object to a map representation for JSON marshaling.
func (s *SponsorLinkValidityDuration) toMap() any {
    switch obj := s.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SponsorLinkValidityDuration.
// It customizes the JSON unmarshaling process for SponsorLinkValidityDuration objects.
func (s *SponsorLinkValidityDuration) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(string), false, &s.isString),
    )
    
    s.value = result
    return err
}

func (s *SponsorLinkValidityDuration) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SponsorLinkValidityDuration) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

// internalSponsorLinkValidityDuration represents a sponsorLinkValidityDuration struct.
// Optional if `sponsor_enabled`==`true`. How long to remain valid sponsored guest request approve/deny link received in email, in minutes. Value is between 5 and 60.
type internalSponsorLinkValidityDuration struct {}

var SponsorLinkValidityDurationContainer internalSponsorLinkValidityDuration

// The internalSponsorLinkValidityDuration instance, wrapping the provided int value.
func (s *internalSponsorLinkValidityDuration) FromNumber(val int) SponsorLinkValidityDuration {
    return SponsorLinkValidityDuration{value: &val}
}

// The internalSponsorLinkValidityDuration instance, wrapping the provided string value.
func (s *internalSponsorLinkValidityDuration) FromString(val string) SponsorLinkValidityDuration {
    return SponsorLinkValidityDuration{value: &val}
}
