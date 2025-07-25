// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// OtherDeviceUpdate represents a OtherDeviceUpdate struct.
type OtherDeviceUpdate struct {
    DeviceMac            *string                `json:"device_mac,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OtherDeviceUpdate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OtherDeviceUpdate) String() string {
    return fmt.Sprintf(
    	"OtherDeviceUpdate[DeviceMac=%v, SiteId=%v, AdditionalProperties=%v]",
    	o.DeviceMac, o.SiteId, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OtherDeviceUpdate.
// It customizes the JSON marshaling process for OtherDeviceUpdate objects.
func (o OtherDeviceUpdate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "device_mac", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OtherDeviceUpdate object to a map representation for JSON marshaling.
func (o OtherDeviceUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
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
    var temp tempOtherDeviceUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_mac", "site_id")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.DeviceMac = temp.DeviceMac
    o.SiteId = temp.SiteId
    return nil
}

// tempOtherDeviceUpdate is a temporary struct used for validating the fields of OtherDeviceUpdate.
type tempOtherDeviceUpdate  struct {
    DeviceMac *string    `json:"device_mac,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
}
