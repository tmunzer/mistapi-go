// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// DeviceSearchRadiusStat represents a DeviceSearchRadiusStat struct.
type DeviceSearchRadiusStat struct {
    // Number of accepted authentication requests
    AuthAccepts          *int                                `json:"auth_accepts,omitempty"`
    // Number of rejected authentication requests
    AuthRejects          *int                                `json:"auth_rejects,omitempty"`
    // Status of the device search radius filter. enum: `up`, `down`, `unreachable`
    AuthServerStatus     *DeviceSearchRadiusFilterStatusEnum `json:"auth_server_status,omitempty"`
    // Number of authentication timeouts
    AuthTimeouts         *int                                `json:"auth_timeouts,omitempty"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceSearchRadiusStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceSearchRadiusStat) String() string {
    return fmt.Sprintf(
    	"DeviceSearchRadiusStat[AuthAccepts=%v, AuthRejects=%v, AuthServerStatus=%v, AuthTimeouts=%v, AdditionalProperties=%v]",
    	d.AuthAccepts, d.AuthRejects, d.AuthServerStatus, d.AuthTimeouts, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceSearchRadiusStat.
// It customizes the JSON marshaling process for DeviceSearchRadiusStat objects.
func (d DeviceSearchRadiusStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "auth_accepts", "auth_rejects", "auth_server_status", "auth_timeouts"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceSearchRadiusStat object to a map representation for JSON marshaling.
func (d DeviceSearchRadiusStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.AuthAccepts != nil {
        structMap["auth_accepts"] = d.AuthAccepts
    }
    if d.AuthRejects != nil {
        structMap["auth_rejects"] = d.AuthRejects
    }
    if d.AuthServerStatus != nil {
        structMap["auth_server_status"] = d.AuthServerStatus
    }
    if d.AuthTimeouts != nil {
        structMap["auth_timeouts"] = d.AuthTimeouts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceSearchRadiusStat.
// It customizes the JSON unmarshaling process for DeviceSearchRadiusStat objects.
func (d *DeviceSearchRadiusStat) UnmarshalJSON(input []byte) error {
    var temp tempDeviceSearchRadiusStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_accepts", "auth_rejects", "auth_server_status", "auth_timeouts")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.AuthAccepts = temp.AuthAccepts
    d.AuthRejects = temp.AuthRejects
    d.AuthServerStatus = temp.AuthServerStatus
    d.AuthTimeouts = temp.AuthTimeouts
    return nil
}

// tempDeviceSearchRadiusStat is a temporary struct used for validating the fields of DeviceSearchRadiusStat.
type tempDeviceSearchRadiusStat  struct {
    AuthAccepts      *int                                `json:"auth_accepts,omitempty"`
    AuthRejects      *int                                `json:"auth_rejects,omitempty"`
    AuthServerStatus *DeviceSearchRadiusFilterStatusEnum `json:"auth_server_status,omitempty"`
    AuthTimeouts     *int                                `json:"auth_timeouts,omitempty"`
}
