package models

import (
    "encoding/json"
)

// DhcpClientOption represents a DhcpClientOption struct.
type DhcpClientOption struct {
    Code                 *string                `json:"code,omitempty"`
    Data                 *string                `json:"data,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpClientOption.
// It customizes the JSON marshaling process for DhcpClientOption objects.
func (d DhcpClientOption) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "code", "data"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpClientOption object to a map representation for JSON marshaling.
func (d DhcpClientOption) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Code != nil {
        structMap["code"] = d.Code
    }
    if d.Data != nil {
        structMap["data"] = d.Data
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpClientOption.
// It customizes the JSON unmarshaling process for DhcpClientOption objects.
func (d *DhcpClientOption) UnmarshalJSON(input []byte) error {
    var temp tempDhcpClientOption
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "code", "data")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Code = temp.Code
    d.Data = temp.Data
    return nil
}

// tempDhcpClientOption is a temporary struct used for validating the fields of DhcpClientOption.
type tempDhcpClientOption  struct {
    Code *string `json:"code,omitempty"`
    Data *string `json:"data,omitempty"`
}
