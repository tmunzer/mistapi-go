// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayMgmtAutoSignatureUpdate represents a GatewayMgmtAutoSignatureUpdate struct.
// Automatic security signature update schedule
type GatewayMgmtAutoSignatureUpdate struct {
	// enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
	DayOfWeek *DayOfWeekEnum `json:"day_of_week,omitempty"`
	// Whether automatic security signature updates are enabled
	Enable *bool `json:"enable,omitempty"`
	// Optional, Mist will decide the timing
	TimeOfDay            *string                `json:"time_of_day,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMgmtAutoSignatureUpdate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMgmtAutoSignatureUpdate) String() string {
	return fmt.Sprintf(
		"GatewayMgmtAutoSignatureUpdate[DayOfWeek=%v, Enable=%v, TimeOfDay=%v, AdditionalProperties=%v]",
		g.DayOfWeek, g.Enable, g.TimeOfDay, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMgmtAutoSignatureUpdate.
// It customizes the JSON marshaling process for GatewayMgmtAutoSignatureUpdate objects.
func (g GatewayMgmtAutoSignatureUpdate) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"day_of_week", "enable", "time_of_day"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMgmtAutoSignatureUpdate object to a map representation for JSON marshaling.
func (g GatewayMgmtAutoSignatureUpdate) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.DayOfWeek != nil {
		structMap["day_of_week"] = g.DayOfWeek
	}
	if g.Enable != nil {
		structMap["enable"] = g.Enable
	}
	if g.TimeOfDay != nil {
		structMap["time_of_day"] = g.TimeOfDay
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMgmtAutoSignatureUpdate.
// It customizes the JSON unmarshaling process for GatewayMgmtAutoSignatureUpdate objects.
func (g *GatewayMgmtAutoSignatureUpdate) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMgmtAutoSignatureUpdate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "day_of_week", "enable", "time_of_day")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.DayOfWeek = temp.DayOfWeek
	g.Enable = temp.Enable
	g.TimeOfDay = temp.TimeOfDay
	return nil
}

// tempGatewayMgmtAutoSignatureUpdate is a temporary struct used for validating the fields of GatewayMgmtAutoSignatureUpdate.
type tempGatewayMgmtAutoSignatureUpdate struct {
	DayOfWeek *DayOfWeekEnum `json:"day_of_week,omitempty"`
	Enable    *bool          `json:"enable,omitempty"`
	TimeOfDay *string        `json:"time_of_day,omitempty"`
}
