package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsTuntermBouncePort represents a UtilsTuntermBouncePort struct.
type UtilsTuntermBouncePort struct {
    // in milli seconds, hold time between multiple port bounces
    HoldTime             *int           `json:"hold_time,omitempty"`
    // list of ports to bounce
    Ports                []string       `json:"ports"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsTuntermBouncePort.
// It customizes the JSON marshaling process for UtilsTuntermBouncePort objects.
func (u UtilsTuntermBouncePort) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsTuntermBouncePort object to a map representation for JSON marshaling.
func (u UtilsTuntermBouncePort) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.HoldTime != nil {
        structMap["hold_time"] = u.HoldTime
    }
    structMap["ports"] = u.Ports
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsTuntermBouncePort.
// It customizes the JSON unmarshaling process for UtilsTuntermBouncePort objects.
func (u *UtilsTuntermBouncePort) UnmarshalJSON(input []byte) error {
    var temp utilsTuntermBouncePort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "hold_time", "ports")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.HoldTime = temp.HoldTime
    u.Ports = *temp.Ports
    return nil
}

// utilsTuntermBouncePort is a temporary struct used for validating the fields of UtilsTuntermBouncePort.
type utilsTuntermBouncePort  struct {
    HoldTime *int      `json:"hold_time,omitempty"`
    Ports    *[]string `json:"ports"`
}

func (u *utilsTuntermBouncePort) validate() error {
    var errs []string
    if u.Ports == nil {
        errs = append(errs, "required field `ports` is missing for type `Utils_Tunterm_Bounce_Port`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
