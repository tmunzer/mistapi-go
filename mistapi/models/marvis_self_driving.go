// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisSelfDriving represents a MarvisSelfDriving struct.
// Self-driving network automation settings per domain
type MarvisSelfDriving struct {
	Wan                  *MarvisSelfDrivingDomain `json:"wan,omitempty"`
	Wired                *MarvisSelfDrivingDomain `json:"wired,omitempty"`
	Wireless             *MarvisSelfDrivingDomain `json:"wireless,omitempty"`
	AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisSelfDriving,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisSelfDriving) String() string {
	return fmt.Sprintf(
		"MarvisSelfDriving[Wan=%v, Wired=%v, Wireless=%v, AdditionalProperties=%v]",
		m.Wan, m.Wired, m.Wireless, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisSelfDriving.
// It customizes the JSON marshaling process for MarvisSelfDriving objects.
func (m MarvisSelfDriving) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"wan", "wired", "wireless"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisSelfDriving object to a map representation for JSON marshaling.
func (m MarvisSelfDriving) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Wan != nil {
		structMap["wan"] = m.Wan.toMap()
	}
	if m.Wired != nil {
		structMap["wired"] = m.Wired.toMap()
	}
	if m.Wireless != nil {
		structMap["wireless"] = m.Wireless.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisSelfDriving.
// It customizes the JSON unmarshaling process for MarvisSelfDriving objects.
func (m *MarvisSelfDriving) UnmarshalJSON(input []byte) error {
	var temp tempMarvisSelfDriving
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wan", "wired", "wireless")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Wan = temp.Wan
	m.Wired = temp.Wired
	m.Wireless = temp.Wireless
	return nil
}

// tempMarvisSelfDriving is a temporary struct used for validating the fields of MarvisSelfDriving.
type tempMarvisSelfDriving struct {
	Wan      *MarvisSelfDrivingDomain `json:"wan,omitempty"`
	Wired    *MarvisSelfDrivingDomain `json:"wired,omitempty"`
	Wireless *MarvisSelfDrivingDomain `json:"wireless,omitempty"`
}
