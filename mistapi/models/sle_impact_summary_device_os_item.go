package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactSummaryDeviceOsItem represents a SleImpactSummaryDeviceOsItem struct.
type SleImpactSummaryDeviceOsItem struct {
    Degraded             float64        `json:"degraded"`
    DeviceOs             string         `json:"device_os"`
    Duration             float64        `json:"duration"`
    Name                 string         `json:"name"`
    Total                float64        `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryDeviceOsItem.
// It customizes the JSON marshaling process for SleImpactSummaryDeviceOsItem objects.
func (s SleImpactSummaryDeviceOsItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryDeviceOsItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryDeviceOsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["degraded"] = s.Degraded
    structMap["device_os"] = s.DeviceOs
    structMap["duration"] = s.Duration
    structMap["name"] = s.Name
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactSummaryDeviceOsItem.
// It customizes the JSON unmarshaling process for SleImpactSummaryDeviceOsItem objects.
func (s *SleImpactSummaryDeviceOsItem) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactSummaryDeviceOsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "degraded", "device_os", "duration", "name", "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Degraded = *temp.Degraded
    s.DeviceOs = *temp.DeviceOs
    s.Duration = *temp.Duration
    s.Name = *temp.Name
    s.Total = *temp.Total
    return nil
}

// tempSleImpactSummaryDeviceOsItem is a temporary struct used for validating the fields of SleImpactSummaryDeviceOsItem.
type tempSleImpactSummaryDeviceOsItem  struct {
    Degraded *float64 `json:"degraded"`
    DeviceOs *string  `json:"device_os"`
    Duration *float64 `json:"duration"`
    Name     *string  `json:"name"`
    Total    *float64 `json:"total"`
}

func (s *tempSleImpactSummaryDeviceOsItem) validate() error {
    var errs []string
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `sle_impact_summary_device_os_item`")
    }
    if s.DeviceOs == nil {
        errs = append(errs, "required field `device_os` is missing for type `sle_impact_summary_device_os_item`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `sle_impact_summary_device_os_item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sle_impact_summary_device_os_item`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `sle_impact_summary_device_os_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
