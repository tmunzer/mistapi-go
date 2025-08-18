// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgUiSettingsTilePosition represents a OrgUiSettingsTilePosition struct.
type OrgUiSettingsTilePosition struct {
	Col                  *int                   `json:"col,omitempty"`
	ColSpan              *int                   `json:"colSpan,omitempty"`
	Row                  *int                   `json:"row,omitempty"`
	RowSpan              *int                   `json:"rowSpan,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgUiSettingsTilePosition,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgUiSettingsTilePosition) String() string {
	return fmt.Sprintf(
		"OrgUiSettingsTilePosition[Col=%v, ColSpan=%v, Row=%v, RowSpan=%v, AdditionalProperties=%v]",
		o.Col, o.ColSpan, o.Row, o.RowSpan, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgUiSettingsTilePosition.
// It customizes the JSON marshaling process for OrgUiSettingsTilePosition objects.
func (o OrgUiSettingsTilePosition) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"col", "colSpan", "row", "rowSpan"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgUiSettingsTilePosition object to a map representation for JSON marshaling.
func (o OrgUiSettingsTilePosition) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Col != nil {
		structMap["col"] = o.Col
	}
	if o.ColSpan != nil {
		structMap["colSpan"] = o.ColSpan
	}
	if o.Row != nil {
		structMap["row"] = o.Row
	}
	if o.RowSpan != nil {
		structMap["rowSpan"] = o.RowSpan
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgUiSettingsTilePosition.
// It customizes the JSON unmarshaling process for OrgUiSettingsTilePosition objects.
func (o *OrgUiSettingsTilePosition) UnmarshalJSON(input []byte) error {
	var temp tempOrgUiSettingsTilePosition
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "col", "colSpan", "row", "rowSpan")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Col = temp.Col
	o.ColSpan = temp.ColSpan
	o.Row = temp.Row
	o.RowSpan = temp.RowSpan
	return nil
}

// tempOrgUiSettingsTilePosition is a temporary struct used for validating the fields of OrgUiSettingsTilePosition.
type tempOrgUiSettingsTilePosition struct {
	Col     *int `json:"col,omitempty"`
	ColSpan *int `json:"colSpan,omitempty"`
	Row     *int `json:"row,omitempty"`
	RowSpan *int `json:"rowSpan,omitempty"`
}
