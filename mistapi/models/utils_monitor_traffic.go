package models

import (
    "encoding/json"
    "fmt"
)

// UtilsMonitorTraffic represents a UtilsMonitorTraffic struct.
type UtilsMonitorTraffic struct {
    // Port name, if no port input is provided then all ports will be monitored
    Port                 *string                `json:"port,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsMonitorTraffic,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsMonitorTraffic) String() string {
    return fmt.Sprintf(
    	"UtilsMonitorTraffic[Port=%v, AdditionalProperties=%v]",
    	u.Port, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsMonitorTraffic.
// It customizes the JSON marshaling process for UtilsMonitorTraffic objects.
func (u UtilsMonitorTraffic) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "port"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsMonitorTraffic object to a map representation for JSON marshaling.
func (u UtilsMonitorTraffic) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Port != nil {
        structMap["port"] = u.Port
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsMonitorTraffic.
// It customizes the JSON unmarshaling process for UtilsMonitorTraffic objects.
func (u *UtilsMonitorTraffic) UnmarshalJSON(input []byte) error {
    var temp tempUtilsMonitorTraffic
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Port = temp.Port
    return nil
}

// tempUtilsMonitorTraffic is a temporary struct used for validating the fields of UtilsMonitorTraffic.
type tempUtilsMonitorTraffic  struct {
    Port *string `json:"port,omitempty"`
}
