// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// UiSettingsTileMetric represents a UiSettingsTileMetric struct.
type UiSettingsTileMetric struct {
    ApiName              *string                `json:"apiName,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UiSettingsTileMetric,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UiSettingsTileMetric) String() string {
    return fmt.Sprintf(
    	"UiSettingsTileMetric[ApiName=%v, AdditionalProperties=%v]",
    	u.ApiName, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UiSettingsTileMetric.
// It customizes the JSON marshaling process for UiSettingsTileMetric objects.
func (u UiSettingsTileMetric) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "apiName"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UiSettingsTileMetric object to a map representation for JSON marshaling.
func (u UiSettingsTileMetric) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.ApiName != nil {
        structMap["apiName"] = u.ApiName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UiSettingsTileMetric.
// It customizes the JSON unmarshaling process for UiSettingsTileMetric objects.
func (u *UiSettingsTileMetric) UnmarshalJSON(input []byte) error {
    var temp tempUiSettingsTileMetric
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "apiName")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.ApiName = temp.ApiName
    return nil
}

// tempUiSettingsTileMetric is a temporary struct used for validating the fields of UiSettingsTileMetric.
type tempUiSettingsTileMetric  struct {
    ApiName *string `json:"apiName,omitempty"`
}
