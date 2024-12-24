package models

import (
    "encoding/json"
    "fmt"
)

// UtilsBouncePort represents a UtilsBouncePort struct.
type UtilsBouncePort struct {
    // list of ports to bounce
    Ports                []string               `json:"ports,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsBouncePort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsBouncePort) String() string {
    return fmt.Sprintf(
    	"UtilsBouncePort[Ports=%v, AdditionalProperties=%v]",
    	u.Ports, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsBouncePort.
// It customizes the JSON marshaling process for UtilsBouncePort objects.
func (u UtilsBouncePort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "ports"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsBouncePort object to a map representation for JSON marshaling.
func (u UtilsBouncePort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Ports != nil {
        structMap["ports"] = u.Ports
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsBouncePort.
// It customizes the JSON unmarshaling process for UtilsBouncePort objects.
func (u *UtilsBouncePort) UnmarshalJSON(input []byte) error {
    var temp tempUtilsBouncePort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ports")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Ports = temp.Ports
    return nil
}

// tempUtilsBouncePort is a temporary struct used for validating the fields of UtilsBouncePort.
type tempUtilsBouncePort  struct {
    Ports []string `json:"ports,omitempty"`
}
