package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UtilsTuntermBouncePort represents a UtilsTuntermBouncePort struct.
type UtilsTuntermBouncePort struct {
    // in milli seconds, hold time between multiple port bounces
    HoldTime             *int                   `json:"hold_time,omitempty"`
    // list of ports to bounce
    Ports                []string               `json:"ports"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsTuntermBouncePort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsTuntermBouncePort) String() string {
    return fmt.Sprintf(
    	"UtilsTuntermBouncePort[HoldTime=%v, Ports=%v, AdditionalProperties=%v]",
    	u.HoldTime, u.Ports, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsTuntermBouncePort.
// It customizes the JSON marshaling process for UtilsTuntermBouncePort objects.
func (u UtilsTuntermBouncePort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "hold_time", "ports"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsTuntermBouncePort object to a map representation for JSON marshaling.
func (u UtilsTuntermBouncePort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.HoldTime != nil {
        structMap["hold_time"] = u.HoldTime
    }
    structMap["ports"] = u.Ports
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsTuntermBouncePort.
// It customizes the JSON unmarshaling process for UtilsTuntermBouncePort objects.
func (u *UtilsTuntermBouncePort) UnmarshalJSON(input []byte) error {
    var temp tempUtilsTuntermBouncePort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "hold_time", "ports")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.HoldTime = temp.HoldTime
    u.Ports = *temp.Ports
    return nil
}

// tempUtilsTuntermBouncePort is a temporary struct used for validating the fields of UtilsTuntermBouncePort.
type tempUtilsTuntermBouncePort  struct {
    HoldTime *int      `json:"hold_time,omitempty"`
    Ports    *[]string `json:"ports"`
}

func (u *tempUtilsTuntermBouncePort) validate() error {
    var errs []string
    if u.Ports == nil {
        errs = append(errs, "required field `ports` is missing for type `utils_tunterm_bounce_port`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
