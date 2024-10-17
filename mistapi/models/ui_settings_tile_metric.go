package models

import (
	"encoding/json"
)

// UiSettingsTileMetric represents a UiSettingsTileMetric struct.
type UiSettingsTileMetric struct {
	ApiName              *string        `json:"apiName,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UiSettingsTileMetric.
// It customizes the JSON marshaling process for UiSettingsTileMetric objects.
func (u UiSettingsTileMetric) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UiSettingsTileMetric object to a map representation for JSON marshaling.
func (u UiSettingsTileMetric) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, u.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "apiName")
	if err != nil {
		return err
	}

	u.AdditionalProperties = additionalProperties
	u.ApiName = temp.ApiName
	return nil
}

// tempUiSettingsTileMetric is a temporary struct used for validating the fields of UiSettingsTileMetric.
type tempUiSettingsTileMetric struct {
	ApiName *string `json:"apiName,omitempty"`
}
