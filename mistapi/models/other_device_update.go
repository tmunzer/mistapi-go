package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// OtherDeviceUpdate represents a OtherDeviceUpdate struct.
type OtherDeviceUpdate struct {
    DeviceMac            *string        `json:"device_mac,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OtherDeviceUpdate.
// It customizes the JSON marshaling process for OtherDeviceUpdate objects.
func (o OtherDeviceUpdate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OtherDeviceUpdate object to a map representation for JSON marshaling.
func (o OtherDeviceUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.DeviceMac != nil {
        structMap["device_mac"] = o.DeviceMac
    }
    if o.SiteId != nil {
        structMap["site_id"] = o.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OtherDeviceUpdate.
// It customizes the JSON unmarshaling process for OtherDeviceUpdate objects.
func (o *OtherDeviceUpdate) UnmarshalJSON(input []byte) error {
    var temp otherDeviceUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "site_id")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.DeviceMac = temp.DeviceMac
    o.SiteId = temp.SiteId
    return nil
}

// otherDeviceUpdate is a temporary struct used for validating the fields of OtherDeviceUpdate.
type otherDeviceUpdate  struct {
    DeviceMac *string    `json:"device_mac,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
}
