package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DeviceUpgrade represents a DeviceUpgrade struct.
type DeviceUpgrade struct {
    // Reboot device immediately after upgrade is completed (Available on Junos OS devices)
    Reboot               *bool          `json:"reboot,omitempty"`
    // reboot start time in epoch
    RebootAt             *int           `json:"reboot_at,omitempty"`
    // Perform recovery snapshot after device is rebooted (Available on Junos OS devices)
    Snapshot             *bool          `json:"snapshot,omitempty"`
    // firmware download start time in epoch
    StartTime            *float64       `json:"start_time,omitempty"`
    // specific version / `stable`, default is to use the latest
    Version              string         `json:"version"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceUpgrade.
// It customizes the JSON marshaling process for DeviceUpgrade objects.
func (d DeviceUpgrade) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceUpgrade object to a map representation for JSON marshaling.
func (d DeviceUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Reboot != nil {
        structMap["reboot"] = d.Reboot
    }
    if d.RebootAt != nil {
        structMap["reboot_at"] = d.RebootAt
    }
    if d.Snapshot != nil {
        structMap["snapshot"] = d.Snapshot
    }
    if d.StartTime != nil {
        structMap["start_time"] = d.StartTime
    }
    structMap["version"] = d.Version
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceUpgrade.
// It customizes the JSON unmarshaling process for DeviceUpgrade objects.
func (d *DeviceUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempDeviceUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "reboot", "reboot_at", "snapshot", "start_time", "version")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Reboot = temp.Reboot
    d.RebootAt = temp.RebootAt
    d.Snapshot = temp.Snapshot
    d.StartTime = temp.StartTime
    d.Version = *temp.Version
    return nil
}

// tempDeviceUpgrade is a temporary struct used for validating the fields of DeviceUpgrade.
type tempDeviceUpgrade  struct {
    Reboot    *bool    `json:"reboot,omitempty"`
    RebootAt  *int     `json:"reboot_at,omitempty"`
    Snapshot  *bool    `json:"snapshot,omitempty"`
    StartTime *float64 `json:"start_time,omitempty"`
    Version   *string  `json:"version"`
}

func (d *tempDeviceUpgrade) validate() error {
    var errs []string
    if d.Version == nil {
        errs = append(errs, "required field `version` is missing for type `device_upgrade`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
