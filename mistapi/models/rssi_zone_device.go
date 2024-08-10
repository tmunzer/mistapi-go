package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// RssiZoneDevice represents a RssiZoneDevice struct.
type RssiZoneDevice struct {
    DeviceId             uuid.UUID      `json:"device_id"`
    // RSSI threshold
    Rssi                 int            `json:"rssi"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RssiZoneDevice.
// It customizes the JSON marshaling process for RssiZoneDevice objects.
func (r RssiZoneDevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RssiZoneDevice object to a map representation for JSON marshaling.
func (r RssiZoneDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["device_id"] = r.DeviceId
    structMap["rssi"] = r.Rssi
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RssiZoneDevice.
// It customizes the JSON unmarshaling process for RssiZoneDevice objects.
func (r *RssiZoneDevice) UnmarshalJSON(input []byte) error {
    var temp tempRssiZoneDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_id", "rssi")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.DeviceId = *temp.DeviceId
    r.Rssi = *temp.Rssi
    return nil
}

// tempRssiZoneDevice is a temporary struct used for validating the fields of RssiZoneDevice.
type tempRssiZoneDevice  struct {
    DeviceId *uuid.UUID `json:"device_id"`
    Rssi     *int       `json:"rssi"`
}

func (r *tempRssiZoneDevice) validate() error {
    var errs []string
    if r.DeviceId == nil {
        errs = append(errs, "required field `device_id` is missing for type `rssi_zone_device`")
    }
    if r.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `rssi_zone_device`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
