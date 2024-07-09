package models

import (
    "encoding/json"
)

// UtilsShowServicePath represents a UtilsShowServicePath struct.
// The exact service name for which to display the service path
type UtilsShowServicePath struct {
    // only for HA
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    ServiceName          *string            `json:"service_name,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowServicePath.
// It customizes the JSON marshaling process for UtilsShowServicePath objects.
func (u UtilsShowServicePath) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowServicePath object to a map representation for JSON marshaling.
func (u UtilsShowServicePath) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.ServiceName != nil {
        structMap["service_name"] = u.ServiceName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowServicePath.
// It customizes the JSON unmarshaling process for UtilsShowServicePath objects.
func (u *UtilsShowServicePath) UnmarshalJSON(input []byte) error {
    var temp utilsShowServicePath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "node", "service_name")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Node = temp.Node
    u.ServiceName = temp.ServiceName
    return nil
}

// utilsShowServicePath is a temporary struct used for validating the fields of UtilsShowServicePath.
type utilsShowServicePath  struct {
    Node        *HaClusterNodeEnum `json:"node,omitempty"`
    ServiceName *string            `json:"service_name,omitempty"`
}
