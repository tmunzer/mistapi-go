package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactSummaryDeviceTypeItem represents a SleImpactSummaryDeviceTypeItem struct.
type SleImpactSummaryDeviceTypeItem struct {
    Degraded             float64        `json:"degraded"`
    DeviceType           string         `json:"device_type"`
    Duration             float64        `json:"duration"`
    Name                 string         `json:"name"`
    Total                float64        `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryDeviceTypeItem.
// It customizes the JSON marshaling process for SleImpactSummaryDeviceTypeItem objects.
func (s SleImpactSummaryDeviceTypeItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryDeviceTypeItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryDeviceTypeItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["degraded"] = s.Degraded
    structMap["device_type"] = s.DeviceType
    structMap["duration"] = s.Duration
    structMap["name"] = s.Name
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactSummaryDeviceTypeItem.
// It customizes the JSON unmarshaling process for SleImpactSummaryDeviceTypeItem objects.
func (s *SleImpactSummaryDeviceTypeItem) UnmarshalJSON(input []byte) error {
    var temp sleImpactSummaryDeviceTypeItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "degraded", "device_type", "duration", "name", "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Degraded = *temp.Degraded
    s.DeviceType = *temp.DeviceType
    s.Duration = *temp.Duration
    s.Name = *temp.Name
    s.Total = *temp.Total
    return nil
}

// sleImpactSummaryDeviceTypeItem is a temporary struct used for validating the fields of SleImpactSummaryDeviceTypeItem.
type sleImpactSummaryDeviceTypeItem  struct {
    Degraded   *float64 `json:"degraded"`
    DeviceType *string  `json:"device_type"`
    Duration   *float64 `json:"duration"`
    Name       *string  `json:"name"`
    Total      *float64 `json:"total"`
}

func (s *sleImpactSummaryDeviceTypeItem) validate() error {
    var errs []string
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `Sle_Impact_Summary_Device_Type_Item`")
    }
    if s.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `Sle_Impact_Summary_Device_Type_Item`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `Sle_Impact_Summary_Device_Type_Item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sle_Impact_Summary_Device_Type_Item`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Sle_Impact_Summary_Device_Type_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}