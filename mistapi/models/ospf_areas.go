package models

import (
    "encoding/json"
)

// OspfAreas represents a OspfAreas struct.
// Junos OSPF areas
type OspfAreas struct {
    IncludeLoopback      *bool                       `json:"include_loopback,omitempty"`
    // Networks to participate in an OSPF area. 
    // Property key is the network name
    Networks             map[string]OspfAreasNetwork `json:"networks,omitempty"`
    // OSPF type. enum: `default`, `nssa`, `stub`
    Type                 *OspfAreasTypeEnum          `json:"type,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OspfAreas.
// It customizes the JSON marshaling process for OspfAreas objects.
func (o OspfAreas) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OspfAreas object to a map representation for JSON marshaling.
func (o OspfAreas) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.IncludeLoopback != nil {
        structMap["include_loopback"] = o.IncludeLoopback
    }
    if o.Networks != nil {
        structMap["networks"] = o.Networks
    }
    if o.Type != nil {
        structMap["type"] = o.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OspfAreas.
// It customizes the JSON unmarshaling process for OspfAreas objects.
func (o *OspfAreas) UnmarshalJSON(input []byte) error {
    var temp ospfAreas
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "include_loopback", "networks", "type")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.IncludeLoopback = temp.IncludeLoopback
    o.Networks = temp.Networks
    o.Type = temp.Type
    return nil
}

// ospfAreas is a temporary struct used for validating the fields of OspfAreas.
type ospfAreas  struct {
    IncludeLoopback *bool                       `json:"include_loopback,omitempty"`
    Networks        map[string]OspfAreasNetwork `json:"networks,omitempty"`
    Type            *OspfAreasTypeEnum          `json:"type,omitempty"`
}
