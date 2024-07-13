package models

import (
    "encoding/json"
)

// OspfConfigArea represents a OspfConfigArea struct.
type OspfConfigArea struct {
    // for a stub/nssa area, where to avoid forwarding type-3 LSA to this area
    NoSummary            *bool          `json:"no_summary,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OspfConfigArea.
// It customizes the JSON marshaling process for OspfConfigArea objects.
func (o OspfConfigArea) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OspfConfigArea object to a map representation for JSON marshaling.
func (o OspfConfigArea) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.NoSummary != nil {
        structMap["no_summary"] = o.NoSummary
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OspfConfigArea.
// It customizes the JSON unmarshaling process for OspfConfigArea objects.
func (o *OspfConfigArea) UnmarshalJSON(input []byte) error {
    var temp ospfConfigArea
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "no_summary")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.NoSummary = temp.NoSummary
    return nil
}

// ospfConfigArea is a temporary struct used for validating the fields of OspfConfigArea.
type ospfConfigArea  struct {
    NoSummary *bool `json:"no_summary,omitempty"`
}
