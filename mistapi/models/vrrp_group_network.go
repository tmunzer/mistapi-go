package models

import (
    "encoding/json"
)

// VrrpGroupNetwork represents a VrrpGroupNetwork struct.
type VrrpGroupNetwork struct {
    Ip                   *string                `json:"ip,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VrrpGroupNetwork.
// It customizes the JSON marshaling process for VrrpGroupNetwork objects.
func (v VrrpGroupNetwork) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "ip"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VrrpGroupNetwork object to a map representation for JSON marshaling.
func (v VrrpGroupNetwork) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Ip != nil {
        structMap["ip"] = v.Ip
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrrpGroupNetwork.
// It customizes the JSON unmarshaling process for VrrpGroupNetwork objects.
func (v *VrrpGroupNetwork) UnmarshalJSON(input []byte) error {
    var temp tempVrrpGroupNetwork
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Ip = temp.Ip
    return nil
}

// tempVrrpGroupNetwork is a temporary struct used for validating the fields of VrrpGroupNetwork.
type tempVrrpGroupNetwork  struct {
    Ip *string `json:"ip,omitempty"`
}
