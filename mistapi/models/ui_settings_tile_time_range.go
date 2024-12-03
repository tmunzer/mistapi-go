package models

import (
    "encoding/json"
)

// UiSettingsTileTimeRange represents a UiSettingsTileTimeRange struct.
type UiSettingsTileTimeRange struct {
    End                  *float64               `json:"end,omitempty"`
    EndDate              *string                `json:"endDate,omitempty"`
    Interval             *string                `json:"interval,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    ShortName            *string                `json:"shortName,omitempty"`
    Start                *int                   `json:"start,omitempty"`
    UsePreset            *bool                  `json:"usePreset,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UiSettingsTileTimeRange.
// It customizes the JSON marshaling process for UiSettingsTileTimeRange objects.
func (u UiSettingsTileTimeRange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "end", "endDate", "interval", "name", "shortName", "start", "usePreset"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UiSettingsTileTimeRange object to a map representation for JSON marshaling.
func (u UiSettingsTileTimeRange) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for UiSettingsTileTimeRange.
// It customizes the JSON unmarshaling process for UiSettingsTileTimeRange objects.
func (u *UiSettingsTileTimeRange) UnmarshalJSON(input []byte) error {
    var temp tempUiSettingsTileTimeRange
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

// tempUiSettingsTileTimeRange is a temporary struct used for validating the fields of UiSettingsTileTimeRange.
type tempUiSettingsTileTimeRange  struct {
    End       *float64 `json:"end,omitempty"`
    EndDate   *string  `json:"endDate,omitempty"`
    Interval  *string  `json:"interval,omitempty"`
    Name      *string  `json:"name,omitempty"`
    ShortName *string  `json:"shortName,omitempty"`
    Start     *int     `json:"start,omitempty"`
    UsePreset *bool    `json:"usePreset,omitempty"`
}
