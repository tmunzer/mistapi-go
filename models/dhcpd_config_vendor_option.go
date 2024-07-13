package models

import (
    "encoding/json"
)

// DhcpdConfigVendorOption represents a DhcpdConfigVendorOption struct.
type DhcpdConfigVendorOption struct {
    Type                 *DhcpdConfigVendorOptionTypeEnum `json:"type,omitempty"`
    Value                *string                          `json:"value,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfigVendorOption.
// It customizes the JSON marshaling process for DhcpdConfigVendorOption objects.
func (d DhcpdConfigVendorOption) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfigVendorOption object to a map representation for JSON marshaling.
func (d DhcpdConfigVendorOption) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.Value != nil {
        structMap["value"] = d.Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfigVendorOption.
// It customizes the JSON unmarshaling process for DhcpdConfigVendorOption objects.
func (d *DhcpdConfigVendorOption) UnmarshalJSON(input []byte) error {
    var temp dhcpdConfigVendorOption
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "type", "value")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Type = temp.Type
    d.Value = temp.Value
    return nil
}

// dhcpdConfigVendorOption is a temporary struct used for validating the fields of DhcpdConfigVendorOption.
type dhcpdConfigVendorOption  struct {
    Type  *DhcpdConfigVendorOptionTypeEnum `json:"type,omitempty"`
    Value *string                          `json:"value,omitempty"`
}
