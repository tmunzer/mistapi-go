// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// IssuedClientCertificate represents a IssuedClientCertificate struct.
type IssuedClientCertificate struct {
    CommonName           *string                `json:"common_name,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    DeviceId             *uuid.UUID             `json:"device_id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    SerialNumber         *string                `json:"serial_number,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IssuedClientCertificate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssuedClientCertificate) String() string {
    return fmt.Sprintf(
    	"IssuedClientCertificate[CommonName=%v, CreatedTime=%v, DeviceId=%v, ModifiedTime=%v, SerialNumber=%v, AdditionalProperties=%v]",
    	i.CommonName, i.CreatedTime, i.DeviceId, i.ModifiedTime, i.SerialNumber, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IssuedClientCertificate.
// It customizes the JSON marshaling process for IssuedClientCertificate objects.
func (i IssuedClientCertificate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "common_name", "created_time", "device_id", "modified_time", "serial_number"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IssuedClientCertificate object to a map representation for JSON marshaling.
func (i IssuedClientCertificate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.CommonName != nil {
        structMap["common_name"] = i.CommonName
    }
    if i.CreatedTime != nil {
        structMap["created_time"] = i.CreatedTime
    }
    if i.DeviceId != nil {
        structMap["device_id"] = i.DeviceId
    }
    if i.ModifiedTime != nil {
        structMap["modified_time"] = i.ModifiedTime
    }
    if i.SerialNumber != nil {
        structMap["serial_number"] = i.SerialNumber
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssuedClientCertificate.
// It customizes the JSON unmarshaling process for IssuedClientCertificate objects.
func (i *IssuedClientCertificate) UnmarshalJSON(input []byte) error {
    var temp tempIssuedClientCertificate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "common_name", "created_time", "device_id", "modified_time", "serial_number")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.CommonName = temp.CommonName
    i.CreatedTime = temp.CreatedTime
    i.DeviceId = temp.DeviceId
    i.ModifiedTime = temp.ModifiedTime
    i.SerialNumber = temp.SerialNumber
    return nil
}

// tempIssuedClientCertificate is a temporary struct used for validating the fields of IssuedClientCertificate.
type tempIssuedClientCertificate  struct {
    CommonName   *string    `json:"common_name,omitempty"`
    CreatedTime  *float64   `json:"created_time,omitempty"`
    DeviceId     *uuid.UUID `json:"device_id,omitempty"`
    ModifiedTime *float64   `json:"modified_time,omitempty"`
    SerialNumber *string    `json:"serial_number,omitempty"`
}
