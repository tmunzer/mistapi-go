package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// DeviceIdString represents a DeviceIdString struct.
type DeviceIdString struct {
    DeviceId             uuid.UUID      `json:"device_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceIdString.
// It customizes the JSON marshaling process for DeviceIdString objects.
func (d DeviceIdString) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceIdString object to a map representation for JSON marshaling.
func (d DeviceIdString) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["device_id"] = d.DeviceId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceIdString.
// It customizes the JSON unmarshaling process for DeviceIdString objects.
func (d *DeviceIdString) UnmarshalJSON(input []byte) error {
    var temp deviceIdString
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_id")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.DeviceId = *temp.DeviceId
    return nil
}

// deviceIdString is a temporary struct used for validating the fields of DeviceIdString.
type deviceIdString  struct {
    DeviceId *uuid.UUID `json:"device_id"`
}

func (d *deviceIdString) validate() error {
    var errs []string
    if d.DeviceId == nil {
        errs = append(errs, "required field `device_id` is missing for type `Device_Id_String`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}