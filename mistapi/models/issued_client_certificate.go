package models

import (
    "encoding/json"
)

// IssuedClientCertificate represents a IssuedClientCertificate struct.
type IssuedClientCertificate struct {
    CreatedTime          *float64       `json:"created_time,omitempty"`
    DeviceId             *string        `json:"device_id,omitempty"`
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    SerialNumber         *string        `json:"serial_number,omitempty"`
    SsoNameId            *string        `json:"sso_name_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssuedClientCertificate.
// It customizes the JSON marshaling process for IssuedClientCertificate objects.
func (i IssuedClientCertificate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IssuedClientCertificate object to a map representation for JSON marshaling.
func (i IssuedClientCertificate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
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
    if i.SsoNameId != nil {
        structMap["sso_name_id"] = i.SsoNameId
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "device_id", "modified_time", "serial_number", "sso_name_id")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.CreatedTime = temp.CreatedTime
    i.DeviceId = temp.DeviceId
    i.ModifiedTime = temp.ModifiedTime
    i.SerialNumber = temp.SerialNumber
    i.SsoNameId = temp.SsoNameId
    return nil
}

// tempIssuedClientCertificate is a temporary struct used for validating the fields of IssuedClientCertificate.
type tempIssuedClientCertificate  struct {
    CreatedTime  *float64 `json:"created_time,omitempty"`
    DeviceId     *string  `json:"device_id,omitempty"`
    ModifiedTime *float64 `json:"modified_time,omitempty"`
    SerialNumber *string  `json:"serial_number,omitempty"`
    SsoNameId    *string  `json:"sso_name_id,omitempty"`
}
