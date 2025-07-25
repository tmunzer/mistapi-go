// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// UiSettingsDefaultTimeRange represents a UiSettingsDefaultTimeRange struct.
type UiSettingsDefaultTimeRange struct {
    End                  *int                   `json:"end,omitempty"`
    EndDate              *string                `json:"endDate,omitempty"`
    Interval             *string                `json:"interval,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    ShortName            *string                `json:"shortName,omitempty"`
    Start                *int                   `json:"start,omitempty"`
    UsePreset            *bool                  `json:"usePreset,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UiSettingsDefaultTimeRange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UiSettingsDefaultTimeRange) String() string {
    return fmt.Sprintf(
    	"UiSettingsDefaultTimeRange[End=%v, EndDate=%v, Interval=%v, Name=%v, ShortName=%v, Start=%v, UsePreset=%v, AdditionalProperties=%v]",
    	u.End, u.EndDate, u.Interval, u.Name, u.ShortName, u.Start, u.UsePreset, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UiSettingsDefaultTimeRange.
// It customizes the JSON marshaling process for UiSettingsDefaultTimeRange objects.
func (u UiSettingsDefaultTimeRange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "end", "endDate", "interval", "name", "shortName", "start", "usePreset"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UiSettingsDefaultTimeRange object to a map representation for JSON marshaling.
func (u UiSettingsDefaultTimeRange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.End != nil {
        structMap["end"] = u.End
    }
    if u.EndDate != nil {
        structMap["endDate"] = u.EndDate
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.ShortName != nil {
        structMap["shortName"] = u.ShortName
    }
    if u.Start != nil {
        structMap["start"] = u.Start
    }
    if u.UsePreset != nil {
        structMap["usePreset"] = u.UsePreset
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UiSettingsDefaultTimeRange.
// It customizes the JSON unmarshaling process for UiSettingsDefaultTimeRange objects.
func (u *UiSettingsDefaultTimeRange) UnmarshalJSON(input []byte) error {
    var temp tempUiSettingsDefaultTimeRange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "endDate", "interval", "name", "shortName", "start", "usePreset")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.End = temp.End
    u.EndDate = temp.EndDate
    u.Interval = temp.Interval
    u.Name = temp.Name
    u.ShortName = temp.ShortName
    u.Start = temp.Start
    u.UsePreset = temp.UsePreset
    return nil
}

// tempUiSettingsDefaultTimeRange is a temporary struct used for validating the fields of UiSettingsDefaultTimeRange.
type tempUiSettingsDefaultTimeRange  struct {
    End       *int    `json:"end,omitempty"`
    EndDate   *string `json:"endDate,omitempty"`
    Interval  *string `json:"interval,omitempty"`
    Name      *string `json:"name,omitempty"`
    ShortName *string `json:"shortName,omitempty"`
    Start     *int    `json:"start,omitempty"`
    UsePreset *bool   `json:"usePreset,omitempty"`
}
